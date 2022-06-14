package reply

import "github.com/mabta/tgbot/types"

const (
	ReplyAPINamnAnswerCallbackQuery = "answerCallbackQuery"
)

func AnswerCallbackQuery(token string, msg *types.CallbackQueryAnswer) (*types.CallbackQueryAnswerResponse, error) {
	return Send[types.CallbackQueryAnswerResponse](token, ReplyAPINamnAnswerCallbackQuery, msg)
}
