package app

import (
	"github.com/gin-gonic/gin"
)

func LanguagesListHandler(languageList map[string]string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.IndentedJSON(200, languageList)
	})
}
