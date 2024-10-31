package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"goapp-skeleton/database"
	"goapp-skeleton/internal/common"
	"goapp-skeleton/internal/routing"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortenUrl(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		return
	}

	router := gin.Default()

	routing.Register(router, db, &common.ApplicationContext{
		InstanceId: "01",
		AppBaseUrl: "http://localhost:8000/",
	})

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/heathcheck", nil)

	router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}
