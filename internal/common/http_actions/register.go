package http_actions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shortener-auth/internal/common/service"
)

type RegisterAction struct {
	registerService *service.RegisterService
}

type RegisterRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a RegisterAction) HandleRegister(context *gin.Context) {
	var registerRequest RegisterRequest

	if err := context.ShouldBind(&registerRequest); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	token, err := a.registerService.Register(registerRequest.Login, registerRequest.Password)

	if err != nil {
		fmt.Println(err)
		context.JSON(400, gin.H{"error": "error while saving shorten url"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
