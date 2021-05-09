package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetAuthorizationToken(ctx *gin.Context) string {
	args := strings.Split(ctx.Request.Header.Get("Authorization"), " ")
	if len(args) == 2 {
		return strings.TrimSpace(strings.Split(ctx.Request.Header.Get("Authorization"), " ")[1])
	}
	return ""
}
