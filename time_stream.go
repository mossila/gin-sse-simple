package main

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grafov/bcast"
	"github.com/mossila/gin-sse-simple/provider"
)

type appContext struct {
	timeGroup *bcast.Group
}

// ChatMsg : message that user talk back to server
type ChatMsg struct {
	User    string `json:"user"`
	Message string `json:"msg"`
}

func main() {
	group := bcast.NewGroup()
	go group.Broadcast(0)
	go provider.TimeProvider(group, 500*time.Millisecond)
	app := appContext{group}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Delims("gin{{", "}}")
	router.LoadHTMLGlob("templates/*")
	router.GET("time/:name", app.timeStream)
	router.POST("msg/", app.message)
	router.GET("/", client)
	router.Run(":8888")
}

func client(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func (app *appContext) message(c *gin.Context) {
	var json ChatMsg
	if err := c.ShouldBindJSON(&json); err == nil {
		app.timeGroup.Send(map[string]interface{}{"type": "chat", "msg": json.Message, "name": json.User})
		c.JSON(http.StatusOK, gin.H{"status": 0})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (app *appContext) timeStream(c *gin.Context) {
	name := c.Param("name")
	recv := app.timeGroup.Join()
	defer recv.Close()
	c.Stream(func(w io.Writer) bool {
		event := recv.Recv().(map[string]interface{})
		if event["type"] == "chat" {
			if event["name"] != name {
				return true
			}
		}
		c.SSEvent("message", event)

		return true
	})
}
