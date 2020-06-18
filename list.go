package gomailman

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrorListGet is an error returned by GetList when response status != 2XX
	ErrorListGet = errors.New("go-mailman: Error getting list")
	// ErrorListAdd is an error returned by AddList when response status != 2XX
	ErrorListAdd = errors.New("go-mailman: Error adding list")
	// ErrorListDelete is an error returned by DeleteList when response status != 2XX
	ErrorListDelete = errors.New("go-mailman: Error deleting list")
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

func (c *Client) GetList(listID string) (*List, error) {
	res, err := c.conn.do(http.MethodGet, buildURL("lists", listID), http.NoBody)
	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		re := &RequestError{}
		if err := json.NewDecoder(res.Body).Decode(re); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("%s: %s", ErrorListGet, re)
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

	if res.StatusCode/100 != 2 {
		re := &RequestError{}
		if err := json.NewDecoder(res.Body).Decode(re); err != nil {
			return err
		}
		return fmt.Errorf("%s %s", ErrorListAdd, re)
	}

	return res.Body.Close()
}

func (c *Client) DeleteList(listID string) error {
	res, err := c.conn.do(http.MethodDelete, buildURL("lists", listID), http.NoBody)
	if err != nil {
		return err
	}

	if res.StatusCode/100 != 2 {
		re := &RequestError{}
		if err := json.NewDecoder(res.Body).Decode(re); err != nil {
			return err
		}

		return fmt.Errorf("%s %s", ErrorListDelete, re)
	}

	return res.Body.Close()
}
