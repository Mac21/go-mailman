package gomailman

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type List struct {
	ID                 string `json:"list_id"`
	Name               string `json:"list_name"`
	MailHost           string `json:"mail_host"`
	Description        string `json:"description"`
	AllowListPosts     bool   `json:"allow_list_posts"`
	Advertised         bool   `json:"advertised"`
	AnonymousList      bool   `json:"anonymous_list"`
	Administrivia      bool   `json:"administrivia"`
	SendGoodbyeMessage bool   `json:"send_goodbye_message"`
	SendWelcomeMessage bool   `json:"send_welcome_message"`
}

// GetList takes a list_id and returns a single list.
func (c *Client) GetList(listID string) (*List, error) {
	res, err := c.conn.do(http.MethodGet, buildURL("lists", listID), http.NoBody)
	if err != nil {
		return nil, err
	}

	if err := parseResponseError(res); err != nil {
		return nil, err
	}

	list := new(List)
	if err = json.NewDecoder(res.Body).Decode(list); err != nil {
		return nil, err
	}

	return list, res.Body.Close()
}

func (c *Client) AddList(listID string) error {
	fakeList := map[string]string{
		"fqdn_listname": listID,
	}

	b, err := json.Marshal(fakeList)
	if err != nil {
		return err
	}

	res, err := c.conn.do(http.MethodPost, buildURL("lists"), bytes.NewReader(b))
	if err != nil {
		return err
	}

	if err := parseResponseError(res); err != nil {
		return err
	}

	return res.Body.Close()
}

func (c *Client) ModifyList(listID string, list *List) error {
	b, err := json.Marshal(list)
	if err != nil {
		return err
	}

	res, err := c.conn.do(http.MethodPost, buildURL("lists"), bytes.NewReader(b))
	if err != nil {
		return err
	}

	if err := parseResponseError(res); err != nil {
		return err
	}

	return res.Body.Close()
}

func (c *Client) DeleteList(listID string) error {
	res, err := c.conn.do(http.MethodDelete, buildURL("lists", listID), http.NoBody)
	if err != nil {
		return err
	}

	if err := parseResponseError(res); err != nil {
		return err
	}

	return res.Body.Close()
}
