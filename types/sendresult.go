package types

type SendResult interface {
	ReplyResponse | APIResponse | CallbackQueryAnswerResponse | GetChatRepsone
}
