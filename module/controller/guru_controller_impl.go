package controller

import (
	model "backend/module/model/guru"
	"backend/module/model/response"
	service "backend/module/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GuruControllerImpl struct {
	GuruSrv service.GuruService
}

func NewGuruControllerImpl(gurusrv service.GuruService) GuruController {
	return &GuruControllerImpl{
		GuruSrv: gurusrv,
	}
}

func (g_ctrl *GuruControllerImpl) CreateGuruCtr(ctx *gin.Context) {
	reqIn := model.Guru{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	guru, err := g_ctrl.GuruSrv.CreateGuru(ctx, reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Status: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success Create Guru",
		Data:   guru,
	})
}
func (g_ctrl *GuruControllerImpl) UpdateGuruCtr(ctx *gin.Context) {
	id, err := g_ctrl.getIdFromParams(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	reqIn := model.Guru{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	reqIn.Id = id
	_, err = g_ctrl.GuruSrv.UpdateGuru(ctx, reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success Update Guru",
	})
}
func (g_ctrl *GuruControllerImpl) FindByIdGuruCtr(ctx *gin.Context) {
	id, err := g_ctrl.getIdFromParams(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	guru, err := g_ctrl.GuruSrv.FindByIdGuru(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ResponseWeb{
			Code:   http.StatusNotFound,
			Status: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ResponseWeb{
		Code:   http.StatusOK,
		Status: "success find guru",
		Data:   guru,
	})
}
func (g_ctrl *GuruControllerImpl) FindAllGuruCtr(ctx *gin.Context) {
	guru, err := g_ctrl.GuruSrv.FindAllGuru(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ResponseWeb{
			Code:   http.StatusNotFound,
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWeb{
		Code:   http.StatusOK,
		Status: "success find AllGuru",
		Data:   guru,
	})
}
func (g_ctrl *GuruControllerImpl) DeleteGuruCtr(ctx *gin.Context) {
	id, err := g_ctrl.getIdFromParams(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	err = g_ctrl.GuruSrv.DeleleByIdGuru(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWeb{
		Code:   http.StatusOK,
		Status: "success delele Guru",
	})
}

func (g_ctrl *GuruControllerImpl) getIdFromParams(ctx *gin.Context) (idUint uint64, err error) {
	id := ctx.Query("id")
	if id == "" {
		err = errors.New("failed id")
		ctx.JSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	// transform id string to uint64
	idUint, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		err = errors.New("failed parse id")
		ctx.JSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	return idUint, err

}
