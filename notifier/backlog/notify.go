package backlog

import (
	"github.com/vvatanabe/tfnotify/terraform"
)

// NotifyService handles communication with the notification related
// methods of Backlog API
type NotifyService service

// Notify posts comment optimized for notifications
func (g *NotifyService) Notify(body string) (exit int, err error) {
	cfg := g.client.Config
	parser := g.client.Config.Parser
	template := g.client.Config.Template

	result := parser.Parse(body)
	if result.Error != nil {
		return result.ExitCode, result.Error
	}
	if result.Result == "" {
		return result.ExitCode, result.Error
	}

	template.SetValue(terraform.CommonTemplate{
		Title:   cfg.PR.Title,
		Message: cfg.PR.Message,
		Result:  result.Result,
		Body:    body,
		Link:    cfg.CI,
	})
	body, err = template.Execute()
	if err != nil {
		return result.ExitCode, err
	}

	return result.ExitCode, g.client.Comment.Post(body, PostOptions{
		Number: cfg.PR.Number,
	})
}
