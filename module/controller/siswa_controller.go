package controller

import "github.com/gin-gonic/gin"

type SiswaController interface {
	CreateSiswaCtr(ctx *gin.Context)
	UpdateSiswaCtr(ctx *gin.Context)
	FindByIdSiswaCtr(ctx *gin.Context)
	FindAllSiswaCtr(ctx *gin.Context)
	DeleteSiswaCtr(ctx *gin.Context)
}
