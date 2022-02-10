package sendemail

import (
	"bytes"
	"net/http"
	"time"
)

/**
* Send an dingtalk text message
* DingTalk 机器人使用白名单方式
**/
func SendDing(webhook, message, keywords string) (err error) {
	content := `{
		"msgtype": "text",
		"text": {"content": ` + keywords + message + `}
	}` // 构造消息，keywords 是配置机器人时候设置的关键字

	client := &http.Client{Timeout: 5 * time.Second}
	_, err = client.Post(webhook, "application/json; charset=utf-8", bytes.NewBuffer([]byte(content)))
	if err != nil {
		return err
	}
	return nil
}
