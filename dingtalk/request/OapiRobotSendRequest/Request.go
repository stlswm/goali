package OapiRobotSendRequest

import (
	"encoding/json"
)

type Request struct {
	msgType string
	text    Text
	link    Link
}

// 设置文本消息
func (o *Request) SetText(text Text) *Request {
	o.msgType = "text"
	o.text = text
	return o
}

// 设置链接消息
func (o *Request) SetLink(link Link) *Request {
	o.msgType = "link"
	o.link = link
	return o
}

// 导出json格式的请求参数
func (o *Request) ExportJsonParams() []byte {
	switch o.msgType {
	case "text":
		jsonByte, _ := json.Marshal(struct {
			MsgType string `json:"msgtype"`
			Text    struct {
				Content string `json:"content"`
			} `json:"text"`
			At struct {
				AtMobiles []string `json:"atMobiles"`
				IsAtAll   bool     `json:"isAtAll"`
			} `json:"at"`
		}{
			"text",
			struct {
				Content string `json:"content"`
			}{
				o.text.Content,
			},
			struct {
				AtMobiles []string `json:"atMobiles"`
				IsAtAll   bool     `json:"isAtAll"`
			}{
				o.text.AtMobiles,
				o.text.IsAtAll,
			},
		})
		return jsonByte
	case "link":
		jsonByte, _ := json.Marshal(struct {
			MsgType string `json:"msgtype"`
			Link    struct {
				Title      string `json:"title"`
				Text       string `json:"text"`
				MessageUrl string `json:"messageUrl"`
				PicUrl     string `json:"picUrl"`
			} `json:"link"`
		}{
			"link",
			struct {
				Title      string `json:"title"`
				Text       string `json:"text"`
				MessageUrl string `json:"messageUrl"`
				PicUrl     string `json:"picUrl"`
			}{
				o.link.Title,
				o.link.Text,
				o.link.MessageUrl,
				o.link.PicUrl,
			},
		})
		return jsonByte
	default:
		panic("not support msg:" + o.msgType)
	}
}
