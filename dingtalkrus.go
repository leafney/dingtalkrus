package dingtalkrus

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var(
	dtHost="https://oapi.dingtalk.com/robot/send"
)

type DingTalkHook struct {
	Token string
	Secret string
	AcceptLevels []logrus.Level
}

type DingTalkResponse struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

func sign(t int64, secret string) string {
	strToHash := fmt.Sprintf("%d\n%s", t, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}


func NewHook(token string,secret string,levels []logrus.Level) *DingTalkHook  {
	return &DingTalkHook{
		Token: token,
		Secret: secret,
		AcceptLevels: levels,
	}
}

func (dh *DingTalkHook) Levels()[]logrus.Level  {
	if dh.AcceptLevels==nil{
		return AllLevels
	}
	return dh.AcceptLevels
}

func (dh *DingTalkHook) Fire(entry *logrus.Entry) error {
	b, err := json.Marshal(entry.Data)
	if err != nil {
		return errors.New("Marshal Fields to JSON error: " + err.Error())
	}
	body := ioutil.NopCloser(bytes.NewBuffer(b))

	value:=url.Values{}
	value.Set("access_token",dh.Token)
	if dh.Secret!=""{
		t := time.Now().UnixNano() / 1e6
		value.Set("timestamp",fmt.Sprintf("%d",t))
		value.Set("sign",sign(t,dh.Secret))
	}

	request,err:=http.NewRequest(http.MethodPost,dtHost,body)
	if err!=nil{
		return errors.New("Error request: " + err.Error())
	}
	request.URL.RawQuery=value.Encode()
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := http.Client{}
	response,err:=client.Do(request)
	if err != nil {
		return errors.New("Send to DingTalk error: " + err.Error())
	}

	defer response.Body.Close()

	rb, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New("Read DingTalk response error: " + err.Error())
	}

	dr := &DingTalkResponse{}
	err = json.Unmarshal(rb, dr)
	if err != nil {
		return errors.New("Unmarshal DingTalk response to JSON error: " + err.Error())
	}

	if dr.ErrCode != 0 {
		return errors.New("DingTalk return error message: " + dr.ErrMsg)
	}

	return nil
}


func SendTextMsg(content string, atMobiles []string, isAtAll bool) logrus.Fields  {
	return map[string]interface{}{
		"msgtype":"text",
		"text":map[string]string{
			"content":content,
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"isAtAll":   isAtAll,
		},
	}
}

func SendMarkdownMsg(title string, text string, atMobiles []string, isAtAll bool)logrus.Fields {
	return map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": title,
			"text":  text,
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"isAtAll":   isAtAll,
		},
	}
}

//
func SendLinkMsg(title string, text string, messageUrl string, picUrl string) logrus.Fields {
	return map[string]interface{}{
		"msgtype": "link",
		"link": map[string]string{
			"title":      title,
			"text":       text,
			"messageUrl": messageUrl,
			"picUrl":     picUrl,
		},
	}
}
