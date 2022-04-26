package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.New()
	engine.GET("/ping", Ping)
	engine.POST("/test", Test)
	engine.Run(":3000")
}

type Body struct {
	Name string `json:"name"`
}

func Test(ctx *gin.Context) {
	var body Body
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(http.StatusOK, body)
}

func Ping(ctx *gin.Context) {
	ctx.String(200, "pong")
}