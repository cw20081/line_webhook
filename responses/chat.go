package responses

import (
	"time"
)

type ChatIndex struct {
	Source  string    `json:"source"`
	Type    string    `json:"type"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

type ChatIndexRes struct {
	Data []ChatIndex `json:"data"`
}
