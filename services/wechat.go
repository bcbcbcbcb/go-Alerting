package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-Alerting/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	wechatGetTokenUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	wechatSendUrl     = "https://qyapi.weixin.qq.com/cgi-bin/message/send"
)

type WeChat struct {
	CorpId      string
	CorpSecret  string
	ToUser      string
	ToParty     string
	AgentId     int
	AccessToken string `json:"access_token"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	ExpiresIn   int    `json:"expires_in"`
}

// 获取 Token
func (w *WeChat) GetToken() error {
	// 配置 wechat 地址
	params := url.Values{}
	params.Set("corpid", w.CorpId)
	params.Set("corpsecret", w.CorpSecret)
	Url, _ := url.Parse(wechatGetTokenUrl)
	Url.RawQuery = params.Encode()

	// 获取wechat 请求的 Token
	hclient := utils.GetHttpCilent()
	if resp, err := hclient.Get(Url.String()); err != nil {
		utils.Log.Errorf("发起获取wechat Token 错误; err = %v \n", err)
		return err
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(body, w)
		if err != nil {
			utils.Log.Errorf("转换json 保存token 错误; err = %v \n", err)
			return err
		}
		if w.ErrCode != 0 {
			utils.Log.Errorf("获取 Token 失败; err = %v \n", w.ErrMsg)
			return fmt.Errorf("1")
		}
		utils.Log.Infof("获取 Token 值为: %v \n", w.AccessToken)

	}
	return nil

}

// 发送微信告警消息
func (w *WeChat) SendMsg(msg string) {
	// 消息
	utils.Log.Infoln("发送的消息是：", msg)
	jsonStr := []byte(msg)

	// 地址 + 微信 Token 拼接
	params := url.Values{}
	params.Set("access_token", w.AccessToken)
	Url, _ := url.Parse(wechatSendUrl)
	Url.RawQuery = params.Encode()

	// 发送http POST请求
	req, _ := http.NewRequest("POST", Url.String(), bytes.NewBuffer(jsonStr))
	hclient := utils.GetHttpCilent()
	resp, err := hclient.Do(req)
	if err != nil {
		utils.Log.Errorf("Error sending to WeChat Work API; err = %v \n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	utils.Log.Infoln("返回结果:", string(body))
}

// 发送消息类型 TextCard
func (w *WeChat) MsgTextCard(h *Hook) string {
	return fmt.Sprintf(`
		{
			"touser" : "%s",
			"toparty" : "%s",
			"msgtype" : "textcard",
			"agentid" : %d,
			"textcard": {
					"title": "%s fenglc测试告警,看到请忽略，谢谢~",
					"description": "<div class=\"time\">发送时间: %s</div><div class=\"host\">故障主机: %s</div><div class=\"state\">故障状态: %s</div><div class=\"rulename\">告警规则名称: %s</div><div class=\"value\">当前指标值: %d</div><div class=\"message\">告警信息: %s</div>",
					"url": "URL",
					"picurl": ""
			},
			"enable_id_trans": 0,
			"enable_duplicate_check": 0,
			"duplicate_check_interval": 1800
		  }
		`, w.ToUser, w.ToParty, w.AgentId, h.Title, time.Now().Format("2006/01/02 15:04:05"), "Null", h.State, h.RuleName, h.EvalMatches[h.ID].Value, h.Message)
}
