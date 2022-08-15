package types

type ReplyMessageParseMode string

const (
	ReplyMessageParseModeText     ReplyMessageParseMode = ""
	ReplyMessageParseModeMarkdown ReplyMessageParseMode = "MarkdownV2"
	ReplyMessageParseModeHTML     ReplyMessageParseMode = "HTML"
)

type ReplyMessage struct {
	ChatID                   ID                    `json:"chat_id"`
	Text                     string                `json:"text"`
	ParseMode                ReplyMessageParseMode `json:"parse_mode,omitempty"`
	ReplyMarkup              *ReplyMessageMarkup   `json:"reply_markup,omitempty"`
	ReplyToMessageID         *int64                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool                 `json:"allow_sending_without_reply,omitempty"`
}

func NewReplyMessage(chatID ID, text string, mode ReplyMessageParseMode) *ReplyMessage {
	return &ReplyMessage{
		ChatID:                   chatID,
		Text:                     text,
		ParseMode:                mode,
		ReplyMarkup:              nil,
		ReplyToMessageID:         nil,
		AllowSendingWithoutReply: nil,
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
		ChatID:                   chatID,
		Text:                     text,
		ReplyMarkup:              keyboard,
		ReplyToMessageID:         nil,
		AllowSendingWithoutReply: nil,
	}
}
