package types

type ReplyMessageParseMode string

const (
	ReplyMessageParseModeText     ReplyMessageParseMode = ""
	ReplyMessageParseModeMarkdown ReplyMessageParseMode = "MarkdownV2"
	ReplyMessageParseModeHTML     ReplyMessageParseMode = "HTML"
)

type ReplyMessage struct {
	ChatID      ID                    `json:"chat_id"`
	Text        string                `json:"text"`
	ParseMode   ReplyMessageParseMode `json:"parse_mode,omitempty"`
	ReplyMarkup *ReplyMessageMarkup   `json:"reply_markup,omitempty"`
}

func NewReplyMessage(chatID ID, text string, mode ReplyMessageParseMode) *ReplyMessage {
	return &ReplyMessage{
		ChatID:      chatID,
		Text:        text,
		ParseMode:   mode,
		ReplyMarkup: nil,
	}
}
func NewReplyTextMessage(chatID ID, text string) *ReplyMessage {
	return NewReplyMessage(chatID, text, ReplyMessageParseModeText)
}
func NewReplyMarkdownMessage(chatID ID, text string) *ReplyMessage {
	return NewReplyMessage(chatID, text, ReplyMessageParseModeMarkdown)
}
func NewReplyKeyboardMessage(chatID ID, text string, keyboard *Keyboard) *ReplyMessage {
	return &ReplyMessage{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: keyboard,
	}
}
