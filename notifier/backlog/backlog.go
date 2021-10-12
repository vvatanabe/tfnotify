package backlog

import (
	"context"

	backlogShared "github.com/vvatanabe/go-backlog/backlog/shared"
	backlog "github.com/vvatanabe/go-backlog/backlog/v2"
)

// API is Backlog API interface
type API interface {
	AddPullRequestComment(ctx context.Context, number int, opt *backlog.AddPullRequestCommentOptions) (*backlog.PullRequestComment, *backlogShared.Response, error)
}

// Backlog represents the attribute information necessary for requesting Backlog API
type Backlog struct {
	*backlog.Client
	project, repository string
}

// AddPullRequestComment is a wrapper of https://developer.nulab.com/docs/backlog/api/2/add-pull-request-comment/
func (b *Backlog) AddPullRequestComment(ctx context.Context, number int, opt *backlog.AddPullRequestCommentOptions) (*backlog.PullRequestComment, *backlogShared.Response, error) {
	return b.Client.PullRequests.AddPullRequestComment(ctx, b.project, b.repository, number, opt)
}
