package reply

import (
	"github.com/mabta/tgbot/types"
)

const (
	ReplyAPINameSetWebhook = "setWebhook"
)

func SetWebhook(token, webhookUrl string) (*types.APIResponse, error) {
	msg := map[string]string{
		"url": webhookUrl,
	}
	return Send[types.APIResponse](token, ReplyAPINameSetWebhook, msg)
}
