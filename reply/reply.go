package reply

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mabta/tgbot/types"
)

const (
	ReplyAPINameSendMessage = "sendMessage"
)

func getReplyAPIUrl(token, apiName string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, apiName)
}
func Send[R types.SendResult](token, apiName string, msg interface{}) (*R, error) {
	url := getReplyAPIUrl(token, apiName)
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	buf, err := httpClient(url, data)
	if err != nil {
		return nil, err
	}
	log.Printf("reply.send 返回数据： %s\n", buf)
	resp := new(R)
	if err := json.Unmarshal(buf, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func Message(token string, msg *types.ReplyMessage) (*types.ReplyResponse, error) {
	return Send[types.ReplyResponse](token, ReplyAPINameSendMessage, msg)
}
