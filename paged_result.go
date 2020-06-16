package gomailman

import (
	"encoding/json"
)

type PagedResult struct {
	Start     int             `json:"start"`
	TotalSize int             `json:"total_size"`
	Entries   json.RawMessage `json:"entries"`
}
