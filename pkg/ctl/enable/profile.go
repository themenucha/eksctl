package enable

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/kris-nova/logger"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
	"github.com/weaveworks/eksctl/pkg/git"
	"github.com/weaveworks/eksctl/pkg/gitops"
	"github.com/weaveworks/eksctl/pkg/gitops/fileprocessor"
	"github.com/weaveworks/eksctl/pkg/utils/file"
)

type options struct {
	gitOptions           git.Options
	quickstartNameArg    string
	gitPrivateSSHKeyPath string
}

func (opts options) validate() error {
	if opts.quickstartNameArg == "" {
		return errors.New("please supply a valid gitops Quick Start URL or name")
	}
	if err := opts.gitOptions.ValidateURL(); err != nil {
		return errors.Wrap(err, "please supply a valid --git-url argument")
	}
	if opts.gitPrivateSSHKeyPath != "" && !file.Exists(opts.gitPrivateSSHKeyPath) {
		return errors.New("please supply a valid --git-private-ssh-key-path argument")
	}
	return nil
}

func enableProfileCmd(cmd *cmdutils.Cmd) {
	cfg := api.NewClusterConfig()
	cmd.ClusterConfig = cfg

	cmd.SetDescription("profile", "Set up Flux and deploy the components from the selected Quick Start profile.", "")

	var opts options

	cmd.SetRunFuncWithNameArg(func() error {
		return doEnableProfile(cmd, opts)
	})

	cmd.FlagSetGroup.InFlagSet("General", func(fs *pflag.FlagSet) {
		fs.StringVar(&cfg.Metadata.Name, "cluster", "", "name of the EKS cluster to use in manifests")
		fs.StringVarP(&opts.quickstartNameArg, "name", "", "", "name or URL of the Quick Start profile. For example, app-dev.")
		fs.StringVarP(&opts.gitOptions.URL, "git-url", "", "", "SSH URL of the Git repository that will contain the cluster components, e.g. git@github.com:<github_org>/<repo_name>")
		fs.StringVarP(&opts.gitOptions.Branch, "git-branch", "", "master", "Git branch")
		fs.StringVar(&opts.gitOptions.User, "git-user", "Flux", "Username to use as Git committer")
		fs.StringVar(&opts.gitOptions.Email, "git-email", "", "Email to use as Git committer")
		fs.StringVar(&opts.gitPrivateSSHKeyPath, "git-private-ssh-key-path", "",
			"Optional path to the private SSH key to use with Git, e.g. ~/.ssh/id_rsa")

		_ = cobra.MarkFlagRequired(fs, "git-url")
		_ = cobra.MarkFlagRequired(fs, "git-email")

		cmdutils.AddRegionFlag(fs, cmd.ProviderConfig)
		cmdutils.AddConfigFileFlag(fs, &cmd.ClusterConfigFile)
		cmdutils.AddTimeoutFlagWithValue(fs, &cmd.ProviderConfig.WaitTimeout, 20*time.Second)
	})

	cmdutils.AddCommonFlagsForAWS(cmd.FlagSetGroup, cmd.ProviderConfig, false)
}

func doEnableProfile(cmd *cmdutils.Cmd, opts options) error {
	if cmd.NameArg != "" && opts.quickstartNameArg != "" {
		return cmdutils.ErrNameFlagAndArg(cmd.NameArg, opts.quickstartNameArg)
	}
	if cmd.NameArg != "" {
		opts.quickstartNameArg = cmd.NameArg
	}
	if err := opts.validate(); err != nil {
		return err
	}

	// TODO move the load of the region outside of the creation of the EKS client
	// currently that is done inside cmd.NewCtl() but we don't need EKS here
	cmd.ClusterConfig.Metadata.Region = cmd.ProviderConfig.Region

	quickstartRepoURL, err := repoURLForQuickstart(opts.quickstartNameArg)
	if err != nil {
		return errors.Wrap(err, "please supply a valid Quick Start name or URL")
	}

	if err := cmdutils.NewGitopsConfigLoader(cmd).Load(); err != nil {
		return err
	}

	// Clone user's repo to apply Quick Start profile
	usersRepoName, err := git.RepoName(opts.gitOptions.URL)
	if err != nil {
		return err
	}
	usersRepoDir, err := ioutil.TempDir("", usersRepoName+"-")
	logger.Debug("Directory %s will be used to clone the configuration repository and install the profile", usersRepoDir)
	profileOutputPath := filepath.Join(usersRepoDir, "base")

	gitClient := git.NewGitClient(git.ClientParams{
		PrivateSSHKeyPath: opts.gitPrivateSSHKeyPath,
	})
	cloneOptions := git.CloneOptions{
		URL:       opts.gitOptions.URL,
		Branch:    opts.gitOptions.Branch,
		Bootstrap: true,
	}
	err = gitClient.CloneRepoInPath(usersRepoDir, cloneOptions)
	if err != nil {
		return err
	}

	processor := &fileprocessor.GoTemplateProcessor{
		Params: fileprocessor.NewTemplateParameters(cmd.ClusterConfig),
	}

	profile := &gitops.Profile{
		Processor: processor,
		Path:      profileOutputPath,
		GitOpts: git.Options{
			URL:    quickstartRepoURL,
			Branch: "master",
		},
		GitCloner: git.NewGitClient(git.ClientParams{}),
		FS:        afero.NewOsFs(),
		IO:        afero.Afero{Fs: afero.NewOsFs()},
	}

	err = profile.Generate(context.Background())
	if err != nil {
		return errors.Wrap(err, "error generating profile")
	}

	// Git add, commit and push component files in the user's repo
	if err = gitClient.Add("."); err != nil {
		return err
	}

	commitMsg := fmt.Sprintf("Add %s quickstart components", opts.quickstartNameArg)
	if err = gitClient.Commit(commitMsg, opts.gitOptions.User, opts.gitOptions.Email); err != nil {
		return err
	}

	if err = gitClient.Push(); err != nil {
		return err
	}

	profile.DeleteClonedDirectory()
	return nil
}

func repoURLForQuickstart(quickstartArgument string) (string, error) {
	if git.IsGitURL(quickstartArgument) {
		return quickstartArgument, nil
	}
	if quickstartArgument == "app-dev" {
		return "git@github.com:weaveworks/eks-quickstart-app-dev.git", nil
	}
	return "", fmt.Errorf("invalid URL or unknown Quick Start %s ", quickstartArgument)
}
