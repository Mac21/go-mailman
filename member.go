package gomailman

type Member struct {
	ID               string           `json:"member_id"`
	ListID           string           `json:"list_id"`
	DisplayName      string           `json:"display_name"`
	DeliveryMode     DeliveryMode     `json:"delivery_mode"`
	SubscriptionMode SubscriptionMode `json:"subscription_mode"`
	ModerationAction Action           `json:"moderation_action"`
	Role             MemberRole       `json:"role"`
}

func (c *Client) AddListMember(member *Member) error {
	return nil
}
