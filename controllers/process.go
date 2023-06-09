package controllers

import (
	"go-Alerting/services"
	"go-Alerting/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Process(c *gin.Context) {
	w := services.WeChat{
		CorpId:     utils.Config.GetString("wechat.corpid"),
		CorpSecret: utils.Config.GetString("wechat.corpsecret"),
		ToUser:     utils.Config.GetString("wechat.touser"),
		ToParty:    utils.Config.GetString("wechat.toparty"),
		AgentId:    utils.Config.GetInt("wechat.agentid"),
	}

	hook := &services.Hook{}
	if err := c.ShouldBindJSON(hook); err != nil {
		c.ShouldBindJSON(hook)
		utils.Log.Errorf("获取传入数据错误; err = %v \n", err)
		_, _ = c.Writer.WriteString("Error on JSON format")
		return
	}

	for id := range hook.EvalMatches {
		hook.ID = id
		err := w.GetToken()
		if err != nil {
			return
		}
		msg := w.MsgTextCard(hook)
		w.SendMsg(msg)
		time.Sleep(1 * time.Second)
	}

}
