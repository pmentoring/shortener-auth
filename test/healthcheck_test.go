package test

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	authactions "shortener-auth/auth/http_actions"
	"shortener-auth/auth/repository"
	"shortener-auth/database"
	"shortener-auth/internal/app/grpc"
	"shortener-auth/internal/common"
	"shortener-auth/internal/routing"
	"testing"
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
	grpcClient := grpc.NewGrpc()
	routing.Register(router, registerAction, grpcClient)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
