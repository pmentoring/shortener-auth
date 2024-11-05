package test

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	pb "github.com/pmentoring/shortener-protoc/gen/go/shortener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"net/http/httptest"
	authactions "shortener-auth/auth/http_actions"
	"shortener-auth/auth/repository"
	"shortener-auth/database"
	"shortener-auth/internal/common"
	"shortener-auth/internal/routing"
	"testing"
)

var (
	addr = flag.String("addr", "go-app:8000", "the address to connect to")
)

func TestHealthCheck(t *testing.T) {
	router := gin.Default()

	conn, err := database.GetConnection()
	if err != nil {
		return
	}

	ctx := common.NewApplicationContext("1", "", "secret")
	repo := repository.NewUserRepository(conn)
	registerAction := authactions.NewRegisterAction(repo, ctx)
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	grpcConn, err := grpc.NewClient(*addr, opts...)
	if err != nil {
		panic(err)
	}

	grpcClient := pb.NewShortenerClient(grpcConn)

	routing.Register(router, registerAction, grpcClient)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
