package backlog

import (
	"context"
	"fmt"

	backlog "github.com/vvatanabe/go-backlog/backlog/v2"
)

// CommentService handles communication with the comment related
// methods of Backlog API
type CommentService service

// PostOptions specifies the optional parameters to post comments to a pull request
type PostOptions struct {
	Number int
}

// Post posts comment
func (g *CommentService) Post(body string, opt PostOptions) error {
	if opt.Number != 0 {
		_, _, err := g.client.API.AddPullRequestComment(
			context.Background(),
			opt.Number,
			&backlog.AddPullRequestCommentOptions{
				Content: body,
			},
		)
		return err
	}
	return fmt.Errorf("backlog.comment.post: Number is required")
}
