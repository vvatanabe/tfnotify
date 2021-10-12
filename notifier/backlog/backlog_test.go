package backlog

import (
	"context"

	backlogShared "github.com/vvatanabe/go-backlog/backlog/shared"
	backlog "github.com/vvatanabe/go-backlog/backlog/v2"
	"github.com/vvatanabe/tfnotify/terraform"
)

type fakeAPI struct {
	API
	FakeAddPullRequestComment func(ctx context.Context, number int, opt *backlog.AddPullRequestCommentOptions) (*backlog.PullRequestComment, *backlogShared.Response, error)
}

func (g *fakeAPI) AddPullRequestComment(ctx context.Context, number int, opt *backlog.AddPullRequestCommentOptions) (*backlog.PullRequestComment, *backlogShared.Response, error) {
	return g.FakeAddPullRequestComment(ctx, number, opt)
}

func newFakeAPI() fakeAPI {
	return fakeAPI{
		FakeAddPullRequestComment: func(ctx context.Context, number int, opt *backlog.AddPullRequestCommentOptions) (*backlog.PullRequestComment, *backlogShared.Response, error) {
			return &backlog.PullRequestComment{
				ID:      371748792,
				Content: "comment 1",
			}, nil, nil
		},
	}
}

func newFakeConfig() Config {
	return Config{
		APIKey:  "token",
		BaseURL: "foo.backlog.com",
		Project: "BAR",
		Repo:    "baz",
		PR: PullRequest{
			Number:  1,
			Message: "message",
		},
		Parser:   terraform.NewPlanParser(),
		Template: terraform.NewPlanTemplate(terraform.DefaultPlanTemplate),
	}
}
