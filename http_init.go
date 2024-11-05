package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
	authactions "shortener-auth/auth/http_actions"
	"shortener-auth/auth/repository"
	"shortener-auth/database"
	"shortener-auth/internal/app/grpc"
	"shortener-auth/internal/common"
	"shortener-auth/internal/routing"
)

func main() {
	logger := createLogger()

	logger.Debug("Application booting....")

	r := gin.Default()

	conn, err := database.GetConnection()

	if err != nil {
		panic(err)
	}
	repo := repository.NewUserRepository(conn)

	registerAction := authactions.NewRegisterAction(repo, getAppContext())
	grpcClient := grpc.NewGrpc()
	routing.Register(r, registerAction, grpcClient)

	err = r.Run("0.0.0.0:8000")

	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func getAppContext() *common.ApplicationContext {
	return common.NewApplicationContext(os.Getenv("INSTANCE_ID"), os.Getenv("APP_BASE_URL"), os.Getenv("APP_SECRET"))
}

func createLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}
