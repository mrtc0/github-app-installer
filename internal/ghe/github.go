package ghe

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func NewGHEClient(baseURL, uploadURL, token string) *GHEClient {
	return &GHEClient{baseURL: baseURL, uploadURL: uploadURL, token: token}
}

type GHEClient struct {
	baseURL   string
	uploadURL string
	token string
	Client    *github.Client
}

func (c *GHEClient) Login(ctx context.Context) error {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client, err := github.NewEnterpriseClient(c.baseURL, c.uploadURL, tc)
	if err != nil {
		return err
	}
	c.Client = client

	return nil
}

func (c *GHEClient) SearchMergePullRequest(ctx context.Context, owner, repo, title, branch string) (*github.PullRequest, error) {
	opt := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	var pull *github.PullRequest
	for {
		pulls, resp, err := c.Client.PullRequests.List(ctx, owner, repo, opt)
		if err != nil {
			return pull, err
		}

		for _, p := range pulls {
			headLabel := fmt.Sprintf("%s:%s", owner, branch)
			if *p.Title == title && *p.Head.Label == headLabel && *p.User.Type == "Bot" {
				pull = p
				return pull, nil
			}
		}
		if resp.NextPage == 0 {
			return pull, errors.New("Not found PullRequest")
		}
		opt.Page = resp.NextPage
	}
}
