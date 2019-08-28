package dingtalk

import (
	"github.com/stlswm/goali/dingtalk/client"
	"github.com/stlswm/goali/dingtalk/request/OapiRobotSendRequest"
	"testing"
)

func TestSend(t *testing.T) {
	cli := client.NewDefaultDingTalkClient("https://oapi.dingtalk.com/robot/send?access_token=03288262d6872a6b4d1a881a84e8005a1dec9c9a9242f7d0353be1d75dfd8b84")
	req := &OapiRobotSendRequest.Request{}
	// 文本消息
	text := OapiRobotSendRequest.Text{}
	text.SetContent("测试文本消息")
	req.SetText(text)
	// 连接消息
	link := OapiRobotSendRequest.Link{}
	link.SetTitle("时代的火车向前开")
	link.SetText("这个即将发布的新版本，创始人陈航（花名“无招”）称它为“红树林”。\n" +
		"而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林")
	link.SetMessageUrl("https://image.baidu.com/")
	link.SetPicUrl("http://e.hiphotos.baidu.com/image/h%3D300/sign=a9e671b9a551f3dedcb2bf64a4eff0ec/4610b912c8fcc3cef70d70409845d688d53f20f7.jpg")
	req.SetLink(link)
	err, res := cli.ExecuteJsonReq(req)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}
