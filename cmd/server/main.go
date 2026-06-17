package main

import (
	"net/http"

	_ "github.com/EldenNetizen/test/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func pong(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func someFunc() error {
	return errors.New("something went wrong")
}

var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	// engine := gin.Default()
	// engine.GET("/ping", pong)
	// engine.Run()

	// err := someFunc()
	// if err != nil {
	// 	err = errors.WithStack(err)

	// 	err = errors.Wrap(err, "failed to execute someFunc")
	// 	fmt.Printf("%+v\n", err)

	// }
	InitLogger()
	defer logger.Sync()
}
