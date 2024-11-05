package routing

import (
	ctx "context"
	"github.com/gin-gonic/gin"
	shortenerv1 "github.com/pmentoring/shortener-protoc/gen/go/shortener"
	"golang.org/x/net/context"
	authactions "shortener-auth/auth/http_actions"
	appactions "shortener-auth/internal/common/http_actions"
)

func Register(
	engine *gin.Engine,
	registerAction *authactions.RegisterAction,
	grpcClient shortenerv1.ShortenerClient,
) {
	engine.GET("/healthcheck", appactions.HandleHealth)
	engine.POST("/register", registerAction.HandleRegister)

	ctx2 := ctx.Background()
	engine.POST("/shorten", shorten(grpcClient, ctx2))
	engine.GET("/:urlCode", unshorten(grpcClient, ctx2))
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
		context.JSON(200, resp)
	}
}
