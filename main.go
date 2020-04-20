package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shenlihartono/used-go-module/math"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"net/http"
	"os"
	"strconv"
)

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% [%lvl%] - %msg%",
		},
	}

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		log.Infoln("ping received\n")
		c.String(http.StatusOK, "pong")
	})

	// Test return result from other module
	r.GET("/addition", func(c *gin.Context) {
		first := c.Query("first")
		second := c.Query("second")

		f, _ := strconv.Atoi(first)
		s, _ := strconv.Atoi(second)
		res := math.Add(f, s)
		c.String(http.StatusOK, "%d", res)
	})

	return r
}
