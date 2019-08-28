package sms

import (
	"encoding/json"
	"regexp"
)

// 短信发送
// 文档地址：https://help.aliyun.com/document_detail/55284.html?spm=a2c4g.11186623.6.557.qnD8Uf
type Request struct {
	ClientHttp
	action          string
	version         string
	regionId        string
	phoneNumbers    []string
	signName        string
	templateCode    string
	templateParam   map[string]string
	smsUpExtendCode string
	outId           string
}

// 增加发送手机号
func (q *Request) AddPhone(phone string) *Request {
	q.phoneNumbers = append(q.phoneNumbers, phone)
	return q
}

// 设置发送手机号
func (q *Request) SetPhones(phones []string) *Request {
	q.phoneNumbers = phones
	return q
}

// 获取所的接收手机号
func (q *Request) GetPhones() []string {
	return q.phoneNumbers
}

// 设置签名
func (q *Request) SetSignName(signName string) *Request {
	q.signName = signName
	return q
}

// 获取签名
func (q *Request) GetSignName(signName string) string {
	return q.signName
}

// 设置模板变量
func (q *Request) SetTemplateParam(templateParam map[string]string) *Request {
	q.templateParam = templateParam
	return q
}

// 获取模板变量
func (q *Request) GetTemplateParam() map[string]string {
	return q.templateParam
}

// 设置短信模板
func (q *Request) SetTemplateCode(templateCode string) *Request {
	q.templateCode = templateCode
	return q
}

// 获取短信模板
func (q *Request) GetTemplateCode() string {
	return q.templateCode
}

// 设置上行短信扩展码
func (q *Request) SetSmsUpExtendCode(smsUpExtendCode string) *Request {
	q.smsUpExtendCode = smsUpExtendCode
	return q
}

// 获取上行短信扩展码
func (q *Request) GetSmsUpExtendCode() string {
	return q.smsUpExtendCode
}

// 设置外部流水
func (q *Request) SetOutId(outId string) *Request {
	q.outId = outId
	return q
}

// 获取外部流水
func (q *Request) GetOutId() string {
	return q.outId
}

// 验证手机号
func (q *Request) checkPhone(phone string) bool {
	reg := regexp.MustCompile(`^\d{4,13}$`)
	return reg.MatchString(phone)
}

// 发送短信
func (q *Request) Send() (bool, string, *Response) {
	if len(q.phoneNumbers) == 0 {
		return false, "请设置接收短信的手机号码", nil
	}
	phones := ""
	for _, phone := range q.phoneNumbers {
		if !q.checkPhone(phone) {
			return false, "手机号码：" + phone + "格式错误", nil
		}
		phones += phone + ","
	}
	phones = phones[0 : len(phones)-1]
	params := make(map[string]string)
	params["Action"] = q.action
	params["Version"] = q.version
	params["RegionId"] = q.regionId
	params["PhoneNumbers"] = phones
	params["SignName"] = q.signName
	params["TemplateCode"] = q.templateCode
	if len(q.templateParam) > 0 {
		templateParam, err := json.Marshal(q.templateParam)
		if err != nil {
			return false, err.Error(), nil
		}
		params["TemplateParam"] = string(templateParam)
	}
	if q.smsUpExtendCode != "" {
		params["SmsUpExtendCode"] = q.smsUpExtendCode
	}
	if q.outId != "" {
		params["OutId"] = q.outId
	}
	ok, responseStr := q.http("POST", params)
	if !ok {
		return false, responseStr, nil
	}
	response := &Response{}
	err := json.Unmarshal([]byte(responseStr), response)
	if err != nil {
		return false, err.Error(), nil
	}
	if response.Code == "OK" {
		return true, "ok", response
	}
	return false, response.Message, nil
}

// 初始化短信发送对象
func NewRequest(accessKeyId string, accessKeySecret string) *Request {
	r := &Request{}
	// 父类参数
	r.accessKeyId = accessKeyId
	r.accessKeySecret = accessKeySecret
	r.format = "JSON"
	r.protocol = "https"
	r.domain = "dysmsapi.aliyuncs.com"

	// 本类参数
	r.templateParam = make(map[string]string)
	r.action = "SendSms"
	r.version = "2017-05-25"
	r.regionId = "cn-hangzhou"
	return r
}
