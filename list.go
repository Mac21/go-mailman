package gomailman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ErrorListGet = "Error getting list"
)

type List struct {
	Name string `json:"list_name"`
	ID   string `json:"list_id"`
}

func (l List) String() string {
	return fmt.Sprintf("%#v", l)
}

func (c *Client) GetList(listID string) (*List, error) {
	res, err := c.conn.do(http.MethodGet, c.buildURL("lists", listID), http.NoBody)
	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("%s: %s", ErrorListGet, res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	list := new(List)
	err = json.Unmarshal(b, list)
	if err != nil {
		return nil, err
	}

	return list, res.Body.Close()
}
