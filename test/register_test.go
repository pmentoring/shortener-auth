package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"shortener-auth/database"
	"shortener-auth/internal/common"
	"shortener-auth/internal/common/http_actions"
	"shortener-auth/internal/routing"
	"strings"
	"testing"
)

func TestRegisterOk(t *testing.T) {
	// arrange
	router := gin.Default()
	db, err := database.GetConnection()
	if err != nil {
		return
	}

	ctx := common.NewApplicationContext("1", "", "secret")
	registerAction := http_actions.NewRegisterAction(db, ctx)
	routing.Register(router, registerAction)

	w := httptest.NewRecorder()
	request := http_actions.RegisterRequest{
		Login:    "enisey",
		Password: "secret",
	}
	jsonBody, _ := json.Marshal(request)

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(string(jsonBody)))
	req.Header.Add("Content-Type", "application/json")

	// act
	router.ServeHTTP(w, req)

	// assert
	var response Response
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	require.NotEmpty(t, response.Token)
	assert.Equal(t, http.StatusOK, w.Code)

	token, err := jwt.Parse(response.Token, func(token *jwt.Token) (interface{}, error) {
		publicKey := ctx.SecretKey
		return []byte(publicKey), nil
	})

	require.NoError(t, err, "Failed to parse JWT token")
	require.True(t, token.Valid, "Token should be valid")

	claims, ok := token.Claims.(jwt.MapClaims)
	require.True(t, ok, "Should be able to get claims")

	assert.Greater(t, claims["sub"], float64(0), "Subject should match")
}

func TestRegisterBadRequest(t *testing.T) {
	// arrange
	router := gin.Default()
	db, err := database.GetConnection()
	if err != nil {
		return
	}

	ctx := common.NewApplicationContext("1", "", "secret")
	registerAction := http_actions.NewRegisterAction(db, ctx)
	routing.Register(router, registerAction)

	w := httptest.NewRecorder()
	request := http_actions.RegisterRequest{
		Login: "enisey",
	}
	jsonBody, _ := json.Marshal(request)

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(string(jsonBody)))
	req.Header.Add("Content-Type", "application/json")

	// act
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

type Response struct {
	Token string `json:"token"`
}
