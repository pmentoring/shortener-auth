package routing

import (
	ctx "context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	shortenerv1 "github.com/pmentoring/shortener-protoc/gen/go/shortener"
	"golang.org/x/net/context"
	authactions "shortener-auth/auth/http_actions"
	"shortener-auth/internal/common"
	appactions "shortener-auth/internal/common/http_actions"
)

func Register(
	engine *gin.Engine,
	registerAction *authactions.RegisterAction,
	grpcClient shortenerv1.ShortenerClient,
	context *common.ApplicationContext,
) {
	authorized := engine.Group("/")
	authorized.Use(AuthRequired(context))
	{
		ctx2 := ctx.Background()
		authorized.POST("/shorten", shorten(grpcClient, ctx2))
		authorized.GET("/:urlCode", unshorten(grpcClient, ctx2))
	}
	engine.GET("/healthcheck", appactions.HandleHealth)
	engine.POST("/register", registerAction.HandleRegister)
}

func AuthRequired(
	appContext *common.ApplicationContext,
) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			publicKey := appContext.SecretKey
			return []byte(publicKey), nil
		})
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			return
		}
		if parsedToken.Valid {
			return
		}
		context.JSON(401, gin.H{"error": "invalid token"})
		return
	}
}

func shorten(
	grpcClient shortenerv1.ShortenerClient,
	ctx context.Context,
) gin.HandlerFunc {
	return func(context *gin.Context) {
		var urlShortenRequest shortenerv1.UrlShortenRequest
		err := context.ShouldBindJSON(&urlShortenRequest)
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		resp, err := grpcClient.Shorten(ctx, &urlShortenRequest)
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		context.JSON(200, resp)
	}
}

func unshorten(
	grpcClient shortenerv1.ShortenerClient,
	ctx context.Context,
) gin.HandlerFunc {
	return func(context *gin.Context) {
		urlCode := context.Param("urlCode")
		resp, err := grpcClient.Unshorten(ctx, &shortenerv1.UrlUnshortenRequest{Url: urlCode})
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		context.Redirect(302, resp.Url)
	}
}
