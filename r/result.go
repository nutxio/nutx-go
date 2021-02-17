package r

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一成功返回
func Ok(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    0,
		"msg":     "success",
		"success": true,
	})
}

func OkWithMsg(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    0,
		"msg":     msg,
		"success": true,
	})
}

func OkWithData(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    0,
		"msg":     "success",
		"success": true,
		"data":    data,
	})
}

// 统一失败返回
func Fail(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    -1,
		"msg":     msg,
		"success": false,
	})
}

func FailWithCodeWithMsg(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     msg,
		"success": false,
	})
}

