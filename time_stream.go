package main

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grafov/bcast"
	"github.com/mossila/gin-sse-simple/provider"
)

type appContext struct {
	timeGroup *bcast.Group
}

func main() {
	group := bcast.NewGroup()
	go group.Broadcast(0)
	go provider.TimeProvider(group, 10*time.Millisecond)
	app := appContext{group}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("time/", app.timeStream)
	router.GET("/", client)
	router.Run(":8888")
}

func client(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func (app *appContext) timeStream(c *gin.Context) {
	recv := app.timeGroup.Join()
	defer recv.Close()
	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", recv.Recv().(string))
		return true
	})
}
