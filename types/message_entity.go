package types

type MessageEntityType string

const (
	// MessageEntityTypeMention `@username`
	MessageEntityTypeMention MessageEntityType = "mention"
	// MessageEntityTypeHashtag `#tag`
	MessageEntityTypeHashtag = "hashtag"
	// MessageEntityTypeCashtag `$USD`
	MessageEntityTypeCashtag = "cashtag"
	// MessageEntityTypeBotCommand `/start@jobs_bot`
	MessageEntityTypeBotCommand = "bot_command"
	// MessageEntityTypeUrl `https://telegram.org`
	MessageEntityTypeUrl = "url"
	// MessageEntityTypeEmail `not-reply@telegram.org`
	MessageEntityTypeEmail = "email"
	// MessageEntityTypePhoneNumber `+1-234-567-8900`
	MessageEntityTypePhoneNumber = "phone_number"
	// MessageEntityTypeBold `bold text`
	MessageEntityTypeBold = "bold"
	// MessageEntityTypeItalic `italic text`
	MessageEntityTypeItalic = "italic"
)

type MessageEntity struct {
	Length int               `json:"length"`
	Offset int               `json:"offset"`
	Type   MessageEntityType `json:"type"`
	Url    string            `json:"url"`
}
