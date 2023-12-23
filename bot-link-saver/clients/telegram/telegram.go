package telegram

import (
	"bot-saver-v1/lib/customErrors"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

type Client struct {
	host, basePath string
	client         http.Client
}

func New(host, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(limit, offset int) ([]Update, error) {
	q := url.Values{}

	q.Add("limit", string(limit))
	q.Add("offset", string(offset))

	data, err := c.doRequest(getUpdatesMethod, q)

	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil

}

func (c *Client) SendMessage(chatID int, text string) error {

	q := url.Values{}

	q.Add("chat_id", string(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)

	if err != nil {
		return customErrors.Wrap("can't send message", err)
	}

	return nil

}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {

	defer func() { err = customErrors.WrapIfErr("can't do request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}
