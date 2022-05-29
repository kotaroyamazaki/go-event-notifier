package slack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SlackClient interface {
	Notify() error
	SetMessage(msg *SlackMessage) SlackClient
}

type slackClient struct {
	webhookURL string
	msg        *SlackMessage
}

func New(webhookURL string) SlackClient {
	return &slackClient{
		webhookURL: webhookURL,
	}
}

func (c *slackClient) Notify() error {
	payload, err := json.Marshal(c.msg)
	if err != nil {
		return err
	}
	u, err := url.ParseRequestURI(c.webhookURL)
	if err != nil {
		return err
	}

	resp, err := http.PostForm(
		u.String(),
		url.Values{"payload": {string(payload)}},
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func (c *slackClient) SetMessage(msg *SlackMessage) SlackClient {
	c.msg = msg
	return c
}
