package types

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	Url          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

type InlineKeyboard = ReplyMessageMarkup

func NewInlineKeyboard(btns [][]*InlineKeyboardButton) *InlineKeyboard {
	return &InlineKeyboard{
		InlineKeyboard: btns,
		Keyboard:       nil,
	}
}

func NewInlineKeyboardBtn(text, url string) *InlineKeyboardButton {
	return &InlineKeyboardButton{
		Text:         text,
		Url:          url,
		CallbackData: "",
	}
}
func NewInlineKeyboardCallbackBtn(text, callbackData string) *InlineKeyboardButton {
	return &InlineKeyboardButton{
		Text:         text,
		Url:          "",
		CallbackData: callbackData,
	}
}
