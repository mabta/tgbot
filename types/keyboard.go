package types

type KeyboardButton struct {
	Text string `json:"text"`
}

type Keyboard = ReplyMessageMarkup

func NewKeyboard(btns [][]string) *Keyboard {
	kbb := make([][]*KeyboardButton, 0, len(btns))
	for _, row := range btns {
		kb := make([]*KeyboardButton, 0, len(row))
		for _, t := range row {
			kb = append(kb, &KeyboardButton{Text: t})
		}
		kbb = append(kbb, kb)
	}
	return &Keyboard{
		Keyboard:       kbb,
		InlineKeyboard: nil,
	}
}
