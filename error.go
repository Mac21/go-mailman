package gomailman

import (
	"fmt"
)

type RequestError struct {
	error
	Title       string `json:"title"`
	Description string `json:"description"`
}

var _ error = RequestError{}

func (re RequestError) Error() string {
	return fmt.Sprintf("%#v", re)
}

func (re RequestError) String() string {
	return fmt.Sprintf("%#v", re)
}
