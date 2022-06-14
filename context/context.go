package context

import (
	"github.com/gin-gonic/gin"
	"github.com/mabta/tgbot/reply"
	"github.com/mabta/tgbot/types"
)

type Context struct {
	ginContext *gin.Context
	token      string
	Update     *types.Update
}

func New(token string, ginContext *gin.Context, update *types.Update) *Context {
	return &Context{
		ginContext: ginContext,
		Update:     update,
		token:      token,
	}
}

func (c *Context) replyMessage(msg string, mode types.ReplyMessageParseMode) (*types.ReplyResponse, error) {
	replyMsg := types.NewReplyMessage(c.ChatID(), msg, mode)
	return c.ReplyMessage(replyMsg)
}
func (c *Context) ReplyMessage(msg *types.ReplyMessage) (*types.ReplyResponse, error) {
	return reply.Message(c.token, msg)
}
func (c *Context) ReplyTextMessage(msg string) (*types.ReplyResponse, error) {
	return c.replyMessage(msg, types.ReplyMessageParseModeText)
}
func (c *Context) ReplyMarkdownMessage(msg string) (*types.ReplyResponse, error) {
	return c.replyMessage(msg, types.ReplyMessageParseModeMarkdown)
}
func (c *Context) ReplyHTMLMessage(msg string) (*types.ReplyResponse, error) {
	return c.replyMessage(msg, types.ReplyMessageParseModeHTML)
}
func (c *Context) ReplyKeyboardMessage(msg string, btns [][]string) (*types.ReplyResponse, error) {
	kb := types.NewKeyboard(btns)
	replyMsg := types.NewReplyKeyboardMessage(c.ChatID(), msg, kb)
	return reply.Message(c.token, replyMsg)
}
func (c *Context) AnswerCallbackQuery(answer *types.CallbackQueryAnswer) (*types.CallbackQueryAnswerResponse, error) {
	answer.CallbackQueryId = c.Update.CallbackQuery.ID
	return reply.AnswerCallbackQuery(c.token, answer)
}
func (c *Context) Answer() (*types.CallbackQueryAnswerResponse, error) {
	answer := &types.CallbackQueryAnswer{}
	return c.AnswerCallbackQuery(answer)
}
func (c *Context) AnswerWithMessage(msg *types.ReplyMessage) (*types.ReplyResponse, *types.CallbackQueryAnswerResponse, error) {
	if msg.ChatID == 0 {
		msg.ChatID = c.ChatID()
	}
	resMsg, err := c.ReplyMessage(msg)
	if err != nil {
		return nil, nil, err
	}
	resAnswer, err := c.Answer()
	if err != nil {
		return nil, nil, err
	}
	return resMsg, resAnswer, nil
}
func (c *Context) ChatID() types.ID {
	if c.Update.CallbackQuery != nil {
		return c.Update.CallbackQuery.Message.Chat.ID
	}
	return c.Update.Message.Chat.ID
}
func (c *Context) UserID() types.ID {
	if c.Update.CallbackQuery != nil {
		return types.ID(c.Update.CallbackQuery.From.ID)
	}
	return types.ID(c.Update.Message.From.ID)
}

func (c *Context) ReplyPhoto(msg *types.ReplyImage) (*types.ReplyResponse, error) {
	return reply.Photo(c.token, msg)
}
