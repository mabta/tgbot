package types

type APIResponse struct {
	OK          bool   `json:"ok"`
	Result      bool   `json:"result,omitempty"`
	Description string `json:"description,omitempty"`
}
