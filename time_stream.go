package main

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("time/", timeStream)
	router.Run(":8080")
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
	for x := range time.Tick(1 * time.Second) {
		timeChan <- x.String()
	}

}
