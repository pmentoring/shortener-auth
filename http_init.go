package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	pb "github.com/pmentoring/shortener-protoc/gen/go/shortener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"os"
	authactions "shortener-auth/auth/http_actions"
	"shortener-auth/auth/repository"
	"shortener-auth/database"
	"shortener-auth/internal/common"
	"shortener-auth/internal/routing"
)

var (
	addr = flag.String("addr", "go-app:8000", "the address to connect to")
)

func main() {
	logger := createLogger()

	logger.Debug("Application booting...")

	r := gin.Default()

	conn, err := database.GetConnection()

	if err != nil {
		panic(err)
	}
	repo := repository.NewUserRepository(conn)

	context := getAppContext()
	registerAction := authactions.NewRegisterAction(repo, context)
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	grpcConn, err := grpc.NewClient(*addr, opts...)
	if err != nil {
		panic(err)
	}

	grpcClient := pb.NewShortenerClient(grpcConn)

	routing.Register(r, registerAction, grpcClient, context)

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
