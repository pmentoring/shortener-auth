package routing

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"shortener-auth/internal/common"
	appactions "shortener-auth/internal/common/http_actions"
)

func Register(engine *gin.Engine, db *sql.DB, ctx *common.ApplicationContext) {
	registerAction := appactions.NewRegisterAction(db, ctx)

	engine.GET("/healthcheck", appactions.HandleHealth)
	engine.POST("/register", registerAction.HandleRegister)
}
