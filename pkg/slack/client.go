package slack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	webhookURL string
	msg        *SlackMessage
}

func New(webhookURL string) *Client {
	return &Client{
		webhookURL: webhookURL,
	}
}

func (c *Client) Notify() error {
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

func (c *Client) SetMessage(msg *SlackMessage) *Client {
	c.msg = msg
	return c
}
