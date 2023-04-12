package controller

import (
	"backend/module/model/response"
	model "backend/module/model/siswa"
	service "backend/module/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SiswaControllerimpl struct {
	SiswaSrv service.SiswaService
}

func NewSiswaContollerimpl(siswasrv service.SiswaService) SiswaController {
	return &SiswaControllerimpl{
		SiswaSrv: siswasrv,
	}
}

func (s_ctrl *SiswaControllerimpl) CreateSiswaCtr(ctx *gin.Context) {
	reqIn := model.Siswa{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	siswa, err := s_ctrl.SiswaSrv.CreateSiswa(ctx, reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Status: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success Create Siswa",
		Data:   siswa,
	})

}
func (s_ctrl *SiswaControllerimpl) UpdateSiswaCtr(ctx *gin.Context) {
	id, err := s_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	reqIn := model.Siswa{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	reqIn.Id = id
	_, err = s_ctrl.SiswaSrv.UpdateSiswa(ctx, reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Success Update Siswa",
	})
}

func (s_ctrl *SiswaControllerimpl) FindByIdSiswaCtr(ctx *gin.Context) {
	id, err := s_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	siswa, err := s_ctrl.SiswaSrv.FindByIdSiswa(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ResponseWeb{
			Code:   http.StatusNotFound,
			Status: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ResponseWeb{
		Code:   http.StatusOK,
		Status: "success find Siswa",
		Data:   siswa,
	})
}
func (s_ctrl *SiswaControllerimpl) FindAllSiswaCtr(ctx *gin.Context) {
	siswa, err := s_ctrl.SiswaSrv.FindAllSiswa(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ResponseWeb{
			Code:   http.StatusNotFound,
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWeb{
		Code:   http.StatusOK,
		Status: "success find AllSiswa",
		Data:   siswa,
	})

}
func (s_ctrl *SiswaControllerimpl) DeleteSiswaCtr(ctx *gin.Context) {
	id, err := s_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	err = s_ctrl.SiswaSrv.DeleleByIdSiswa(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResponseWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWeb{
		Code:   http.StatusOK,
		Status: "success delele Siswa",
	})
}

func (s_ctrl *SiswaControllerimpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
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
