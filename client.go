package gomailman

type Client struct {
	conn Connection
}

func NewClient(host, username, password string) *Client {
	return &Client{
		conn: NewConnection(host, username, password),
	}
}
