package http_actions

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"shortener-auth/internal/common"
	"shortener-auth/internal/common/repository"
	"shortener-auth/internal/common/service"
)

type RegisterAction struct {
	registerService *service.RegisterService
	loginService    *service.LoginService
}

func NewRegisterAction(db *sql.DB, ctx *common.ApplicationContext) *RegisterAction {
	repo := repository.NewUserRepository(db)
	return &RegisterAction{
		registerService: service.NewRegisterService(repo),
		loginService: service.NewLoginService(
			repo,
			service.NewJWTService(ctx.SecretKey),
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
