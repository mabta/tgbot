package reply

import "github.com/mabta/tgbot/types"

const (
	ReplyAPINameSendPhone = "sendPhoto"
)

func Photo(token string, msg *types.ReplyImage) (*types.ReplyResponse, error) {
	return Send[types.ReplyResponse](token, ReplyAPINameSendPhone, msg)
}
