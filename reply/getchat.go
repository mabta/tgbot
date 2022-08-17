package reply

import "github.com/mabta/tgbot/types"

const (
	ReplyAPINameGetChat = "getChat"
)

func GetChat(token string, chatID types.ID) (*types.GetChatRepsone, error) {
	return Send[types.GetChatRepsone](token, ReplyAPINameGetChat, &types.ReplyGetChat{ChatID: chatID})
}
