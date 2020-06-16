package gomailman

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	// ErrorListGet is an error returned by GetList when response status != 2XX
	ErrorListGet = errors.New("go-mailman: Error getting list")
	// ErrorListAdd is an error returned by AddList when response status != 2XX
	ErrorListAdd = errors.New("go-mailman: Error adding list")
	// ErrorListDelete is an error returned by DeleteList when response status != 2XX
	ErrorListAdd = errors.New("go-mailman: Error deleting list")
)

type List struct {
	Name string `json:"list_name"`
	ID   string `json:"list_id"`
}

func (c *Client) GetList(listID string) (*List, error) {
	res, err := c.conn.do(http.MethodGet, c.buildURL("lists", listID), http.NoBody)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		re := &RequestError{}
		if err := json.Unmarshal(b, re); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("%s: %s", ErrorListGet, re)
	}

	list := new(List)
	err = json.Unmarshal(b, list)
	if err != nil {
		return nil, err
	}

	return list, res.Body.Close()
}

func (c *Client) AddList(list *List) error {
	if list == nil {
		return errors.New("go-mailman: Error adding nil list")
	}

	b, err := json.Marshal(list)
	if err != nil {
		return err
	}

	res, err := c.conn.do(http.MethodPost, c.buildURL("lists"), bytes.NewReader(b))
	if err != nil {
		return err
	}

	if res.StatusCode/100 != 2 {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		re := &RequestError{}
		if err := json.Unmarshal(b, re); err != nil {
			return err
		}
		return fmt.Errorf("%s %s", ErrorListAdd, re)
	}

	return res.Body.Close()
}

func (c *Client) DeleteList(listID string) error {
	res, err := c.conn.do(http.MethodDelete, c.buildURL("lists", listID), http.NoBody)
	if err != nil {
		return err
	}

	if res.StatusCode/100 != 2 {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		re := &RequestError{}
		if err := json.Unmarshal(b, re); err != nil {
			return err
		}

		return fmt.Errorf("%s %s", ErrorListDelete, re)
	}

	return res.Body.Close()
}
