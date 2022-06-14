package types

type ReplyMessageMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard,omitempty"`
	Keyboard       [][]*KeyboardButton       `json:"keyboard,omitempty"`
}
