package client

import (
	"bytes"
	"encoding/json"
	"github.com/stlswm/goali/dingtalk/request"
	"github.com/stlswm/goali/dingtalk/response"
	"io/ioutil"
	"net/http"
)

type DefaultDingTalkClient struct {
	url string
}

// 执行请求
func (c *DefaultDingTalkClient) ExecuteJsonReq(request request.Request) (error, *response.OApiRobotSendResponse) {
	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(request.ExportJsonParams()))
	if err != nil {
		return err, nil
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, nil
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	res := &response.OApiRobotSendResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return err, nil
	}
	return nil, res
}

// 获取实例
func NewDefaultDingTalkClient(url string) DefaultDingTalkClient {
	return DefaultDingTalkClient{
		url: url,
	}
}
