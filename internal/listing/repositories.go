package listing

import (
	"fmt"
	"context"
	"time"

	"github.com/google/go-github/v32/github"
	"github.com/mrtc0/github-app-installer/internal/ghe"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func ListRepository(c *cli.Context) error {
	ctx := context.Background()

	baseURL := c.String("ghe-base-url")
	uploadURL := c.String("ghe-upload-url")
	token := c.String("ghe-access-token")

	organization := c.Args().First()
	if organization == "" {
		return errors.New(fmt.Sprintf("organization is required"))
	}

	gheClient := ghe.NewGHEClient(baseURL, uploadURL, token)
  err := gheClient.Login(ctx)
  if err != nil {
      return err
  }

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	var allRepos []*github.Repository
	for {
		repos, resp, err := gheClient.Client.Repositories.ListByOrg(ctx, organization, opt)
		if err != nil {
			return err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	for _, repo := range allRepos {
		updatedAt := *repo.UpdatedAt
		d := time.Date(updatedAt.Year(), updatedAt.Month(), updatedAt.Day(), 0, 0, 0, 0, time.Local)
		d2, err := time.Parse("2006-01-02", c.String("after"))
		if err != nil {
			return err
		}

		if d.After(d2) {
			fmt.Printf("%s/%s\n", organization, *repo.Name)
		}
	}

	return nil
}
