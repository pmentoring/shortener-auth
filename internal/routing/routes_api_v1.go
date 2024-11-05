package routing

import (
	"github.com/gin-gonic/gin"
	authactions "shortener-auth/auth/http_actions"
	appactions "shortener-auth/internal/common/http_actions"
)

func Register(engine *gin.Engine, registerAction *authactions.RegisterAction) {
	engine.GET("/healthcheck", appactions.HandleHealth)
	engine.POST("/register", registerAction.HandleRegister)
}
