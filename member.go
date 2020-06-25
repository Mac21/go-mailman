package gomailman

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Member struct {
	ID               string           `json:"member_id"`
	ListID           string           `json:"list_id"`
	DisplayName      string           `json:"display_name"`
	Email            string           `json:"email"`
	DeliveryMode     DeliveryMode     `json:"delivery_mode"`
	SubscriptionMode SubscriptionMode `json:"subscription_mode"`
	ModerationAction Action           `json:"moderation_action"`
	Role             MemberRole       `json:"role"`
}

type ListSubscriber struct {
	ListID       string       `json:"list_id"`
	Subscriber   string       `json:"subscriber"`
	DisplayName  string       `json:"display_name"`
	DeliveryMode DeliveryMode `json:"delivery_mode,omitempty"`
	Role         MemberRole   `json:"role,omitempty"`
	PreVerified  bool         `json:"pre_verified,string"`
	PreConfirmed bool         `json:"pre_confirmed,string"`
	PreApproved  bool         `json:"pre_approved,string"`
}

func (c *Client) AddListMember(sub *ListSubscriber) error {
	b, err := json.Marshal(sub)
	if err != nil {
		return err
	}

	res, err := c.conn.do(http.MethodPost, buildURL("members"), bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	if err := parseResponseError(res); err != nil {
		return err
	}

	return res.Body.Close()
}

func (c *Client) GetListMembers(listID string) ([]*Member, error) {
	params := map[string]string{
		"list_id": listID,
	}

	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := c.conn.do(http.MethodGet, buildURL("members", "find"), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	if err := parseResponseError(res); err != nil {
		return nil, err
	}

	pr := &PagedResult{}
	if err = json.NewDecoder(res.Body).Decode(pr); err != nil {
		return nil, err
	}

	members := make([]*Member, 0)
	if err = json.Unmarshal(pr.Entries, &members); err != nil {
		return nil, err
	}

	return members, res.Body.Close()
}

type RemovedMembers map[string]bool

func (c *Client) DeleteListMembers(listID string, emails []string) (RemovedMembers, error) {
	params := map[string][]string{
		"emails": emails,
	}

	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := c.conn.do(http.MethodDelete, buildURL("lists", listID, "roster", "member"), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	if err := parseResponseError(res); err != nil {
		return nil, err
	}

	removedMembers := make(RemovedMembers)
	if err = json.NewDecoder(res.Body).Decode(&removedMembers); err != nil {
		return nil, err
	}

	return removedMembers, nil
}
