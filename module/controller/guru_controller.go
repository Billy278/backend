package controller

import "github.com/gin-gonic/gin"

type GuruController interface {
	CreateGuruCtr(ctx *gin.Context)
	UpdateGuruCtr(ctx *gin.Context)
	FindByIdGuruCtr(ctx *gin.Context)
	FindAllGuruCtr(ctx *gin.Context)
	DeleteGuruCtr(ctx *gin.Context)
}
