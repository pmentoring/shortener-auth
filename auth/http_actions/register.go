package http_actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shortener-auth/auth/repository"
	service2 "shortener-auth/auth/service"
	"shortener-auth/internal/common"
)

type RegisterAction struct {
	registerService *service2.RegisterService
	loginService    *service2.LoginService
}

func NewRegisterAction(repo repository.UserRepository, ctx *common.ApplicationContext) *RegisterAction {
	return &RegisterAction{
		registerService: service2.NewRegisterService(repo),
		loginService: service2.NewLoginService(
			repo,
			service2.NewJWTService(ctx.SecretKey),
		),
	}
}

type RegisterRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a RegisterAction) HandleRegister(context *gin.Context) {
	var registerRequest RegisterRequest

	err := context.ShouldBind(&registerRequest)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err = a.registerService.Register(registerRequest.Login, registerRequest.Password)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := a.loginService.Login(registerRequest.Login, registerRequest.Password)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
