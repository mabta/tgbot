package types

type ReplyResponse struct {
	OK          bool     `json:"ok"`
	Result      *Message `json:"result,omitempty"`
	Description string   `json:"description,omitempty"`
}
