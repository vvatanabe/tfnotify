package backlog

import (
	"errors"
	"net/http"
	"os"
	"strings"

	backlog "github.com/vvatanabe/go-backlog/backlog/v2"
	"github.com/vvatanabe/tfnotify/terraform"
)

// EnvApiKey is Backlog API Key
const EnvApiKey = "BACKLOG_API_KEY"

// EnvBaseURL is Backlog base URL. This can be set to a domain endpoint to use with Backlog.
const EnvBaseURL = "BACKLOG_BASE_URL"

// Client ...
type Client struct {
	*backlog.Client
	Debug  bool
	Config Config

	common service

	Comment *CommentService
	Notify  *NotifyService

	API API
}

// Config is a configuration for Backlog client
type Config struct {
	APIKey   string
	BaseURL  string
	Project  string
	Repo     string
	PR       PullRequest
	CI       string
	Parser   terraform.Parser
	Template terraform.Template
}

// PullRequest represents Backlog Pull Request metadata
type PullRequest struct {
	Title   string
	Message string
	Number  int
}

type service struct {
	client *Client
}

// NewClient returns Client initialized with Config
func NewClient(cfg Config) (*Client, error) {
	apiKey := cfg.APIKey
	apiKey = strings.TrimPrefix(apiKey, "$")
	if apiKey == EnvApiKey {
		apiKey = os.Getenv(EnvApiKey)
	}
	if apiKey == "" {
		return &Client{}, errors.New("backlog api key is missing")
	}
	baseURL := cfg.BaseURL
	baseURL = strings.TrimPrefix(baseURL, "$")
	if baseURL == EnvBaseURL {
		baseURL = os.Getenv(EnvBaseURL)
	}

	if baseURL == EnvBaseURL {
		baseURL = os.Getenv(EnvBaseURL)
	}
	if baseURL == "" {
		baseURL = os.Getenv(EnvBaseURL)
	}

	client := backlog.NewClient(baseURL, http.DefaultClient)
	client.SetAPIKey(apiKey)

	c := &Client{
		Config: cfg,
		Client: client,
	}
	c.common.client = c
	c.Comment = (*CommentService)(&c.common)
	c.Notify = (*NotifyService)(&c.common)

	c.API = &Backlog{
		Client:     client,
		project:    cfg.Project,
		repository: cfg.Repo,
	}

	return c, nil
}
