package routing

import (
	"github.com/gin-gonic/gin"
	appactions "shortener-auth/internal/common/http_actions"
)

func Register(engine *gin.Engine, registerAction *appactions.RegisterAction) {
	engine.GET("/healthcheck", appactions.HandleHealth)
	engine.POST("/register", registerAction.HandleRegister)
}
