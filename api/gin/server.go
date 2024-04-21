package ginserver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(port string) {
	r := gin.New()

	r.Use(
		gin.Recovery(),
	)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "gin pong",
		})
	})

	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
