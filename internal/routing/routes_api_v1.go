package routing

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"shortener-auth/internal/common"
	appactions "shortener-auth/internal/common/http_actions"
)

func Register(engine *gin.Engine, db *sql.DB, ctx *common.ApplicationContext) {
	engine.GET("/healthcheck", appactions.HandleHealth)
}
