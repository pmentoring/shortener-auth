package http_actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
