package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

var Version = "dev"

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", handler)
	r.Run()
}

func handler(c *gin.Context) {
	hostname := os.Getenv("HOSTNAME")
	c.String(http.StatusOK, fmt.Sprintf("another try from %s  %s", hostname, Version))
}
