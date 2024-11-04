package routing

import (
	"github.com/gin-gonic/gin"
	appactions "shortener-auth/internal/common/http_actions"
)

func Register(engine *gin.Engine) {
	engine.GET("/healthcheck", appactions.HandleHealth)
	engine.GET("/register", appactions.HandleRegister)

}
