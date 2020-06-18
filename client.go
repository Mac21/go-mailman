package gomailman

type Client struct {
	conn *Connection
}

func NewClient(host, username, password string) (*Client, error) {
	conn, err := NewConnection(host, username, password)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func buildURL(parts ...string) string {
	var url string
	for indx, p := range parts {
		if indx == 0 {
			url += p
		} else {
			url += "/" + p
		}
	}

	return url
}
