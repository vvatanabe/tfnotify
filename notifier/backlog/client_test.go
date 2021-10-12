package backlog

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	backlogAPIKey := os.Getenv(EnvApiKey)
	defer func() {
		os.Setenv(EnvApiKey, backlogAPIKey)
	}()
	os.Setenv(EnvApiKey, "")

	testCases := []struct {
		config   Config
		envToken string
		expect   string
	}{
		{
			// specify directly
			config:   Config{APIKey: "abcdefg"},
			envToken: "",
			expect:   "",
		},
		{
			// specify via env but not to be set env (part 1)
			config:   Config{APIKey: "BACKLOG_API_KEY"},
			envToken: "",
			expect:   "backlog api key is missing",
		},
		{
			// specify via env (part 1)
			config:   Config{APIKey: "BACKLOG_API_KEY"},
			envToken: "abcdefg",
			expect:   "",
		},
		{
			// specify via env but not to be set env (part 2)
			config:   Config{APIKey: "BACKLOG_API_KEY"},
			envToken: "",
			expect:   "backlog api key is missing",
		},
		{
			// specify via env (part 2)
			config:   Config{APIKey: "BACKLOG_API_KEY"},
			envToken: "abcdefg",
			expect:   "",
		},
		{
			// no specification (part 1)
			config:   Config{},
			envToken: "",
			expect:   "backlog api key is missing",
		},
		{
			// no specification (part 2)
			config:   Config{},
			envToken: "abcdefg",
			expect:   "backlog api key is missing",
		},
	}
	for _, testCase := range testCases {
		os.Setenv(EnvApiKey, testCase.envToken)
		_, err := NewClient(testCase.config)
		if err == nil {
			continue
		}
		if err.Error() != testCase.expect {
			t.Errorf("got %q but want %q", err.Error(), testCase.expect)
		}
	}
}

func TestNewClientWithBaseURL(t *testing.T) {
	backlogBaseURL := os.Getenv(EnvBaseURL)
	defer func() {
		os.Setenv(EnvBaseURL, backlogBaseURL)
	}()
	os.Setenv(EnvBaseURL, "")

	testCases := []struct {
		config     Config
		envBaseURL string
		expect     string
	}{
		{
			// specify directly
			config: Config{
				APIKey:  "abcdefg",
				BaseURL: "https://foo.backlog.com/",
			},
			envBaseURL: "",
			expect:     "https://foo.backlog.com/api/v2/",
		},
		{
			// specify via env (part 1)
			config: Config{
				APIKey:  "abcdefg",
				BaseURL: "BACKLOG_BASE_URL",
			},
			envBaseURL: "https://foo.backlog.com/",
			expect:     "https://foo.backlog.com/api/v2/",
		},
		{
			// specify via env (part 2)
			config: Config{
				APIKey:  "abcdefg",
				BaseURL: "BACKLOG_BASE_URL",
			},
			envBaseURL: "https://foo.backlog.com/",
			expect:     "https://foo.backlog.com/api/v2/",
		},
		{
			// no specification (part 2)
			config:     Config{APIKey: "abcdefg"},
			envBaseURL: "https://foo.backlog.com/",
			expect:     "https://foo.backlog.com/api/v2/",
		},
	}
	for _, testCase := range testCases {
		os.Setenv(EnvBaseURL, testCase.envBaseURL)
		c, err := NewClient(testCase.config)
		if err != nil {
			continue
		}
		url := c.Client.BaseURL().String()
		if url != testCase.expect {
			t.Errorf("got %q but want %q", url, testCase.expect)
		}
	}
}
