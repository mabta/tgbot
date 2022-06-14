package reply

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func httpClient(url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json;charset=utf-8")
	req.Header.Set("user-agent", "mabta-telegram-bot-client")
	cli := http.Client{
		Timeout: 3 * time.Second,
	}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
