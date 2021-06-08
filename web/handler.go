package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.Redirect(http.StatusFound, "http://localhost:8080")
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 0,
	//	"message": "ok",
	//})
}
