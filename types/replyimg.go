package types

type ReplyImage struct {
	ChatID      ID                    `json:"chat_id"`
	Photo       string                `json:"photo"`
	ParseMode   ReplyMessageParseMode `json:"parse_mode,omitempty"`
	ReplyMarkup *ReplyMessageMarkup   `json:"reply_markup,omitempty"`
	Caption     string                `json:"caption"`
}
