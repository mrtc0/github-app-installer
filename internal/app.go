package internal

import (
	"time"

	"github.com/urfave/cli/v2"
	"github.com/mrtc0/github-app-installer/internal/installer"
	"github.com/mrtc0/github-app-installer/internal/listing"
	"github.com/mrtc0/github-app-installer/internal/automerge"
)

var (
	currentTime = time.Now()
	currentDate = time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		0, 0, 0, 0, time.Local,
	)
	latestCommitDay = currentDate.AddDate(0, -6, 0).Format("2006-01-02")

	githubAppInstallationIdFlag = cli.Int64Flag{
		Name: "installation-id",
		EnvVars: []string{"INSTALLATION_ID"},
	}

	githubBaseUrlFlag = cli.StringFlag{
		Name: "ghe-base-url",
		EnvVars: []string{"GHE_BASE_URL"},
	}

	githubUploadUrlFlag = cli.StringFlag{
		Name: "ghe-upload-url",
		EnvVars: []string{"GHE_UPLOAD_URL"},
	}

	githubAccessTokenFlag = cli.StringFlag{
		Name: "ghe-access-token",
		EnvVars: []string{"GHE_ACCESS_TOKEN"},
	}

	latestCommitDayFlag = cli.StringFlag{
		Name: "after",
		Value: latestCommitDay,
	}

	mergePullRequestTitleFlag = cli.StringFlag{
		Name: "title",
	}

	mergePullRequestBranchNameFlag = cli.StringFlag{
		Name: "branch-name",
	}

	mergePullRequestFetchLimitFlag = cli.IntFlag{
		Name: "retry-count",
		Value: 10,
	}

	listFlags = []cli.Flag{
		&githubBaseUrlFlag,
		&githubUploadUrlFlag,
		&githubAccessTokenFlag,
		&latestCommitDayFlag,
	}

	installFlags = []cli.Flag{
		&githubAppInstallationIdFlag,
		&githubBaseUrlFlag,
		&githubUploadUrlFlag,
		&githubAccessTokenFlag,
	}

	mergeFlags = []cli.Flag{
		&githubBaseUrlFlag,
		&githubUploadUrlFlag,
		&githubAccessTokenFlag,
		&mergePullRequestBranchNameFlag,
		&mergePullRequestTitleFlag,
		&mergePullRequestFetchLimitFlag,
	}
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "github-app-installer"
	app.ArgsUsage = "org"
	app.EnableBashCompletion = true

	app.Commands = []*cli.Command{
		NewInstallCommand(),
		RepositoryListCommand(),
		AutoMergeCommand(),
	}
	app.Action = installer.RunInstall
	return app
}

func NewInstallCommand() *cli.Command {
	return &cli.Command{
		Name: "install",
		Aliases: []string{"i"},
		ArgsUsage: "repository",
		Usage: "install",
		Action: installer.RunInstall,
		Flags: installFlags,
	}
}

func RepositoryListCommand() *cli.Command {
	return &cli.Command{
		Name: "list",
		Aliases: []string{"l"},
		ArgsUsage: "organization",
		Usage: "list",
		Action:	listing.ListRepository,
		Flags: listFlags,
	}
}

func AutoMergeCommand() *cli.Command {
	return &cli.Command{
		Name: "auto-merge",
		ArgsUsage: "repository",
		Usage: "auto-merge",
		Action: automerge.RunMerge,
		Flags: mergeFlags,
	}
}
