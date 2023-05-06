package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Alow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// semua origin mendapat ijin akses
		ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")

		// semua method diperbolehkan masuk
		ctx.Writer.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")

		// semua header diperbolehkan untuk disisipkan
		ctx.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "3600")
		if ctx.Request.Method == http.MethodOptions {
			ctx.ShouldBindJSON(http.StatusOK)
			return
		}
		//fmt.Println("tesssssss")
		ctx.Next()

	}

}
