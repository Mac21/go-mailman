package gomailman

// RequestError is a type that represents a request error returned by Mailman
type RequestError struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var _ error = RequestError{}

func (re RequestError) Error() string {
	if re.Title != re.Description {
		return re.Title + " " + re.Description
	}
	return re.Title
}

func (re RequestError) String() string {
	if re.Title != re.Description {
		return re.Title + " " + re.Description
	}
	return re.Title
}
