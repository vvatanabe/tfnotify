package backlog

import (
	"testing"
)

func TestCommentPost(t *testing.T) {
	testCases := []struct {
		config Config
		body   string
		opt    PostOptions
		ok     bool
	}{
		{
			config: newFakeConfig(),
			body:   "",
			opt: PostOptions{
				Number: 1,
			},
			ok: true,
		},
		{
			config: newFakeConfig(),
			body:   "",
			opt: PostOptions{
				Number: 0,
			},
			ok: false,
		},
		{
			config: newFakeConfig(),
			body:   "",
			opt: PostOptions{
				Number: 2,
			},
			ok: true,
		},
		{
			config: newFakeConfig(),
			body:   "",
			opt: PostOptions{
				Number: 0,
			},
			ok: false,
		},
	}

	for _, testCase := range testCases {
		client, err := NewClient(testCase.config)
		if err != nil {
			t.Fatal(err)
		}
		api := newFakeAPI()
		client.API = &api
		err = client.Comment.Post(testCase.body, testCase.opt)
		if (err == nil) != testCase.ok {
			t.Errorf("got error %q", err)
		}
	}
}
