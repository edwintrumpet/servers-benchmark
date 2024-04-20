package ginserver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(port string) {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "gin pong",
		})
	})

	log.Fatal(r.Run(port))
}
