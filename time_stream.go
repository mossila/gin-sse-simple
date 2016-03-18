package main

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
    gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
    router.LoadHTMLGlob("templates/*")
	router.GET("time/", timeStream)
    router.GET("client/", client)
	router.Run(":8080")
}
func client(c *gin.Context) {
    c.HTML(200, "index.tmpl", 
    gin.H{})
}

func timeStream(c *gin.Context) {
	timeChan := make(chan string)
	go timeProvider(timeChan)
	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", <-timeChan)
		return true
	})
}
func timeProvider(timeChan chan string) {
	defer close(timeChan)
	for x := range time.Tick(2 * time.Millisecond) {
		timeChan <- x.String()
	}

}
