package types

type GetChatRepsone struct {
	OK          bool   `json:"ok"`
	Result      *Chat  `json:"result,omitempty"`
	Description string `json:"description,omitempty"`
}
