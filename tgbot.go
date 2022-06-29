package tgbot

import (
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/mabta/tgbot/context"
	"github.com/mabta/tgbot/option"
	"github.com/mabta/tgbot/reply"
	"github.com/mabta/tgbot/types"
)

type Context = context.Context
type BotHandlerFun func(*Context)
type BotAnswerFun func(*Context)
type BotHandler struct {
	cmd        string
	handlerFun BotHandlerFun
	regex      *regexp.Regexp
}
type BotAnswerHandler struct {
	data       string
	handlerFun BotAnswerFun
	regex      *regexp.Regexp
}

var (
	BotDefaultHandlerFunc = func(c *Context) {
		log.Println("BotDefaultHandlerFunc")
	}
	BotDefaultAnswerFunc = func(c *Context) {
		c.Answer()
		log.Println("BotDefaultAnswerFunc")
	}
)

type Bot struct {
	ginEngine          *gin.Engine
	handers            []*BotHandler
	answers            []*BotAnswerHandler
	opt                *option.Option
	defaultHandlerFunc BotHandlerFun
	defaultAnswerFunc  BotAnswerFun
}

func New(ginEngine *gin.Engine, opt *option.Option) *Bot {
	return &Bot{
		ginEngine:          ginEngine,
		handers:            make([]*BotHandler, 0),
		answers:            make([]*BotAnswerHandler, 0),
		opt:                opt,
		defaultHandlerFunc: nil,
		defaultAnswerFunc:  BotDefaultAnswerFunc,
	}
}

func (b *Bot) SetDefaultHandlerFunc(f BotHandlerFun) {
	b.defaultHandlerFunc = f
}

func (b *Bot) SetDefaultAnswerFunc(f BotAnswerFun) {
	b.defaultAnswerFunc = f
}

func (b *Bot) AddHandler(cmd string, hander BotHandlerFun) {
	b.handers = append(b.handers, &BotHandler{
		cmd:        cmd,
		handlerFun: hander,
		regex:      nil,
	})
}

func (b *Bot) SetStartHandler(hander BotHandlerFun) {
	b.AddHandler("/start", hander)
}
func (b *Bot) AddRegexpHandler(expr string, hander BotHandlerFun) error {
	regex, err := regexp.Compile(expr)
	if err != nil {
		return err
	}
	b.handers = append(b.handers, &BotHandler{
		handlerFun: hander,
		regex:      regex,
	})
	return nil
}
func (b *Bot) AddAnwser(data string, hander BotAnswerFun) {
	b.answers = append(b.answers, &BotAnswerHandler{
		data:       data,
		handlerFun: hander,
	})
}
func (b *Bot) AddRegexpAnwser(expr string, hander BotAnswerFun) error {
	regex, err := regexp.Compile(expr)
	if err != nil {
		return err
	}
	b.answers = append(b.answers, &BotAnswerHandler{
		handlerFun: hander,
		regex:      regex,
	})
	return nil
}

func (b *Bot) Serve() error {
	b.Setup()
	return b.ginEngine.Run()
}
func (b *Bot) Setup() {
	b.setupHandler()
}

func (b *Bot) SetWebhook() (*types.APIResponse, error) {
	webhookUrl := b.opt.WebhookUrl()
	return reply.SetWebhook(b.opt.Token, webhookUrl)
}

func (b *Bot) setupHandler() {
	b.ginEngine.POST(b.opt.UpdatePath, func(c *gin.Context) {
		//rawBody, _ := ioutil.ReadAll(c.Request.Body)
		//log.Printf("原始请求数据：%s\n", rawBody)
		update := new(types.Update)
		if err := c.ShouldBindJSON(update); err != nil {
			log.Println("err:", err)
			return
		}
		ctx := context.New(b.opt.Token, c, update)
		// 分发
		if update.Message != nil {
			for _, handler := range b.handers {
				if handler.regex != nil {
					if handler.regex.MatchString(update.Message.Text) {
						handler.handlerFun(ctx)
						return
					}
				} else {
					if handler.cmd == update.Message.Text {
						handler.handlerFun(ctx)
						return
					}
				}
			}
			if b.defaultHandlerFunc != nil {
				b.defaultHandlerFunc(ctx)
				return
			}
		}
		if update.CallbackQuery != nil {
			for _, handler := range b.answers {
				if handler.regex != nil {
					if handler.regex.MatchString(update.CallbackQuery.Data) {
						handler.handlerFun(ctx)
						return
					}
				} else {
					if handler.data == update.CallbackQuery.Data {
						handler.handlerFun(ctx)
						return
					}
				}
			}
			if b.defaultAnswerFunc != nil {
				b.defaultAnswerFunc(ctx)
				return
			}
		}
	})
}
