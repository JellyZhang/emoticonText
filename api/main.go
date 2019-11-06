package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// trans api
func Text2EmoticonText(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": process(c.PostForm("sourceText")),
	})
}

// ping
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// process Text to emoticon text
func process(source string) string {
	fmt.Println(source)
	for i := 0; i < len(source); i++ {
		fmt.Println(source[i])
	}
	source += "\xF0\x9F\x8D\x91"
	return source
}
