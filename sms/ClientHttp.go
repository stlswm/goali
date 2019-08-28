package sms

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// 文档地址：https://help.aliyun.com/document_detail/56189.html?spm=a2c4g.11186623.6.581.l4bMyL
type ClientHttp struct {
	// 系统参数
	accessKeyId string
	format      string
	// 业务参数
	// 其他参数
	accessKeySecret string
	protocol        string
	domain          string
	apiRequestUrl   string
	apiRequestBody  string
	apiResponse     string
}

// 获取请求地址
func (c *ClientHttp) GetApiRequestUrl() string {
	return c.apiRequestUrl
}

// 获取请求参数
func (c *ClientHttp) GetApiRequestBody() string {
	return c.apiRequestBody
}

// 获取请求返回数据
func (c *ClientHttp) GetApiResponse() string {
	return c.apiResponse
}

// 生成随机码
func (c *ClientHttp) generateNonce() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err)
	}
	s := base64.URLEncoding.EncodeToString(b)
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 特殊URL转码
func (c *ClientHttp) specialUrlEncode(dst string) string {
	dst = url.QueryEscape(dst)
	dst = strings.Replace(dst, "+", "%20", -1)
	dst = strings.Replace(dst, "*", "%2A", -1)
	dst = strings.Replace(dst, "%7E", "~", -1)
	return dst
}

// 签名
func (c *ClientHttp) generateSign(httpMethod string, data map[string]string) string {
	delete(data, "Signature")
	keys := make([]string, 0)
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	sortedQueryString := ""
	for _, key := range keys {
		sortedQueryString += c.specialUrlEncode(key) + "=" + c.specialUrlEncode(data[key]) + "&"
	}
	sortedQueryString = sortedQueryString[0 : len(sortedQueryString)-1]
	stringToSign := strings.ToUpper(httpMethod) + "&" + c.specialUrlEncode("/") + "&" + c.specialUrlEncode(sortedQueryString)
	// HMac use sha1
	key := []byte(c.accessKeySecret + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// GET请求
func (c *ClientHttp) httpGet(data map[string]string) (bool, string) {
	c.apiRequestUrl = c.protocol + "://" + c.domain + "/?"
	urlValue := url.Values{}
	for key, value := range data {
		urlValue.Add(key, value)
	}
	c.apiRequestUrl += urlValue.Encode()
	client := &http.Client{}
	request, err := http.NewRequest("GET", c.apiRequestUrl, nil)
	if err != nil {
		return false, err.Error()
	}
	// 处理返回结果
	response, err := client.Do(request)
	if err != nil {
		return false, err.Error()
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err.Error()
	}
	c.apiResponse = string(body)
	return true, c.apiResponse
}

// POST请求
func (c *ClientHttp) httpPost(data map[string]string) (bool, string) {
	c.apiRequestUrl = c.protocol + "://" + c.domain + "/"
	urlValues := url.Values{}
	for key, value := range data {
		urlValues.Add(key, value)
	}
	c.apiRequestBody = urlValues.Encode()
	response, err := http.PostForm(c.apiRequestUrl, urlValues)
	if err != nil {
		return false, err.Error()
	}
	defer response.Body.Close()
	// 处理返回结果
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err.Error()
	}
	c.apiResponse = string(body)
	return true, c.apiResponse
}

// 网络请求
func (c *ClientHttp) http(method string, data map[string]string) (bool, string) {
	if c.protocol == "" {
		return false, "请设置通讯协议"
	}
	if c.domain == "" {
		return false, "请设置接口远程主机域名"
	}
	loc, _ := time.LoadLocation("GMT")
	data["AccessKeyId"] = c.accessKeyId
	data["Timestamp"] = time.Now().In(loc).Format("2006-01-02T15:04:05Z")
	data["Format"] = c.format
	data["SignatureMethod"] = "HMAC-SHA1"
	data["SignatureVersion"] = "1.0"
	data["SignatureNonce"] = c.generateNonce()
	data["Signature"] = c.generateSign(method, data)
	switch strings.ToLower(method) {
	case "get":
		return c.httpGet(data)
	default:
		return c.httpPost(data)
	}
}
