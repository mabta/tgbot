package types

type Message struct {
	Chat     *Chat            `json:"chat"`
	Date     int64            `json:"date"`
	Entities []*MessageEntity `json:"entities"`
	From     *User            `json:"from"`
	ID       int64            `json:"message_id"` // message_id
	Text     string           `json:"text"`
}
