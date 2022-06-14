package types

type CallbackQueryAnswer struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	Url             string `json:"url,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
}
