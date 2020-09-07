package installer

import (
	"fmt"
	"context"
	"strings"

	"github.com/mrtc0/github-app-installer/internal/ghe"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func RunInstall(c *cli.Context) error {
	ctx := context.Background()

	baseURL := c.String("ghe-base-url")
	uploadURL := c.String("ghe-upload-url")
	token := c.String("ghe-access-token")
	installationID := c.Int64("installation-id")

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

	if installationID == 0 {
		return errors.New(fmt.Sprintf("installation-id is required. installation_id is the number at the end of the URL when you select the github app you want to install from %s/organizations/%s/settings/installations", baseURL, organization))
	}

	gheClient := ghe.NewGHEClient(baseURL, uploadURL, token)
	err := gheClient.Login(ctx)
	if err != nil {
		return err
	}

	repo, _, err := gheClient.Client.Repositories.Get(ctx, organization, repository_name)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", *repo.Name)

	repo, _, err = gheClient.Client.Apps.AddRepository(ctx, installationID, *repo.ID)
	if err != nil {
		return err
	}

	fmt.Println("install completed")
	return nil
}

