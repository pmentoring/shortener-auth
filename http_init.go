package main

import (
	"github.com/gin-gonic/gin"
	"goapp-skeleton/database"
	"goapp-skeleton/internal/common"
	"goapp-skeleton/internal/routing"
	"log/slog"
	"os"
)

func main() {
	logger := createLogger()

	logger.Debug("Application booting...")

	r := gin.Default()

	conn, err := database.GetConnection()

	if err != nil {
		panic(err)
	}

	routing.Register(r, conn, getAppContext())

	err = r.Run("0.0.0.0:8000")

	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func getAppContext() *common.ApplicationContext {
	return common.NewApplicationContext(os.Getenv("INSTANCE_ID"), os.Getenv("APP_BASE_URL"))
}

func createLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}
