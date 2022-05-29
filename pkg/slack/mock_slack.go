package slack

import "fmt"

type mockSlackClient struct {
	msg *SlackMessage
}

func NewMockSlackClient() SlackClient {
	return &mockSlackClient{}
}

func (c *mockSlackClient) Notify() error {
	fmt.Println(c.msg)
	return nil
}

func (c *mockSlackClient) SetMessage(msg *SlackMessage) SlackClient {
	c.msg = msg
	return c
}
