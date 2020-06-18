package gomailman

import "fmt"

// RequestError is a type that represents a request error returned by Mailman
type RequestError struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var _ error = RequestError{}

func (re RequestError) Error() string {
	return fmt.Sprintf("go-mailman: Error %s %s", re.Title, re.Description)
}
