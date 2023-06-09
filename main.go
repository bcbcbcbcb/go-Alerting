package main

import (
	"go-Alerting/controllers"
	"go-Alerting/utils"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin 配置日志输出
	gin.DefaultWriter = io.MultiWriter(utils.Logger)
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.POST("/alert", controllers.Process)
	r.Run(utils.Config.GetString("app.address") + ":" + utils.Config.GetString("app.port"))
}
