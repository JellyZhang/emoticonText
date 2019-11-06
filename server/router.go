package server

import (
	"emoticonText/api"
	"emoticonText/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// middleware
	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1")
	{
		//ping
		v1.GET("ping", api.Ping)

		// Text2EmoticonText
		v1.POST("trans", api.Text2EmoticonText)

	}

	return r
}
