package automerge

import (
	"fmt"
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/mrtc0/github-app-installer/internal/ghe"
)

func RunMerge(c *cli.Context) error {
	ctx := context.Background()

	baseURL := c.String("ghe-base-url")
	uploadURL := c.String("ghe-upload-url")
	token := c.String("ghe-access-token")

	gheClient := ghe.NewGHEClient(baseURL, uploadURL, token)
	err := gheClient.Login(ctx)
	if err != nil {
		return err
	}

	repository := c.Args().First()
	if repository == "" {
		return errors.New(fmt.Sprintf("repository is required."))
	}

	s := strings.Split(repository, "/")
	if len(s) != 2 {
		return errors.New(fmt.Sprintf("invalid repository format"))
	}

	organization := s[0]
	repository_name := s[1]

	title := c.String("title")
	branch := c.String("branch-name")

	limit := c.Int("retry-count")
	for i := 0; i < limit; i++ {
		pull, err := gheClient.SearchMergePullRequest(ctx, organization, repository_name, title, branch)
		if err != nil {
			if err.Error() == "Not found PullRequest" {
				fmt.Println("Not found pull request. sleep 10 sec...")
				time.Sleep(10 * time.Second)
				continue
			}
			return err
		}

		result, _, err := gheClient.Client.PullRequests.Merge(ctx, organization, repository_name, *pull.Number, "auto merge by github-app-installer", nil)
		if err != nil {
			return err
		}

		if *result.Merged {
			fmt.Printf("The pull request with %d was merged in %s.\n", *pull.Number, repository)
			break
		} else {
			fmt.Printf("The pull request with %d was merge failed in %s.\n", *pull.Number, repository)
			break
		}
	}

	return nil
}
