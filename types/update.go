package types

type Update struct {
	Message       *Message       `json:"message"`
	ID            int64          `json:"update_id"`
	CallbackQuery *CallbackQuery `json:"callback_query"`
}
