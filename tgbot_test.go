package tgbot

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mabta/tgbot/option"
	"github.com/mabta/tgbot/types"
)

func testNewGinEngine() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	app.GET("/", func(c *gin.Context) {
		c.String(200, "index")
	})
	app.GET("/foobar", func(c *gin.Context) {
		c.String(200, "foobar")
	})
	app.POST("/foobar", func(c *gin.Context) {
		c.String(200, "foobar - post")
	})
	return app
}

func TestServe(t *testing.T) {
	app := testNewGinEngine()
	opt := option.New("TOKEN", "DOMAIN")
	bot := New(app, opt)
	usage := func(c *Context) {
		//reps, err := c.ReplyTextMessage("帮助信息 ")
		//reps, err := c.ReplyMarkdownMessage("__帮助信息__")
		//reps, err := c.ReplyHTMLMessage("<b>帮助信息</b>")
		reps, err := c.ReplyKeyboardMessage(fmt.Sprintf("请选择你要执行的操作%v", c.UserID()), [][]string{
			[]string{"🍎 苹果", "🍌 香蕉", "🍊 桔子"},
			[]string{"🍚 吃好饭", "😪 睡好觉"},
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Log(reps)
	}
	//resp, err := bot.SetWebhook()
	//t.Log(resp, err)
	//return
	bot.SetStartHandler(usage)
	bot.AddHandler("/help", usage)
	bot.AddAnwser("rank", func(c *Context) {
		msg := types.NewReplyTextMessage(c.ChatID(), fmt.Sprintf("苹果%v", c.UserID()))
		msgReps, answerReps, err := c.AnswerWithMessage(msg)
		t.Log(msgReps, answerReps, err)
	})
	bot.AddAnwser("example", func(c *Context) {
		t.Log(c.AnswerWithMessage(types.NewReplyMarkdownMessage(c.ChatID(), "示例")))
	})
	bot.AddHandler("🍎 苹果", func(c *Context) {
		btns := [][]*types.InlineKeyboardButton{
			{
				types.NewInlineKeyboardBtn("在线购买", "https://google.com"),
				types.NewInlineKeyboardCallbackBtn("教程", "example"),
				types.NewInlineKeyboardCallbackBtn("排行榜", "rank"),
			},
		}
		caption := strings.Join([]string{
			"【苹果🍎】\n",
			"💕 食用方法",
			"你可以直接咬，也可以清洗或削皮再吃",
			"\n",
			"🚰 清洗",
			"把水龙头拧开，涂上82年的洗洁精1\\.3克",
			"\n",
			"🍉 削皮",
			"准备好水果刀，别把自己的猪手削了",
			"\n",
			"`AAGpiz9t_t9L9_d9KH81`",
			"👆点击复制",
		}, "\n")
		msg := &types.ReplyImage{
			ChatID:      c.ChatID(),
			Photo:       "https://www.inventicons.com/uploads/iconset/1537/wm/512/Qr-code-52.png",
			Caption:     caption,
			ParseMode:   types.ReplyMessageParseModeMarkdown,
			ReplyMarkup: types.NewInlineKeyboard(btns),
		}
		resp, err := c.ReplyPhoto(msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
	})
	bot.AddRegexpHandler(`^\d+$`, func(c *Context) {
		res, err := c.ReplyTextMessage(fmt.Sprintf("你输入的内容是：%q", c.Update.Message.Text))
		t.Log(res, err)
	})
	bot.Serve()
}
