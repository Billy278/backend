package router

import (
	"backend/module/controller"
	middleware "backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, guruCtrl controller.GuruController, siswaCtrl controller.SiswaController) {
	//register all router
	groupSiswa := r.Group("/siswa")
	groupSiswa.OPTIONS("", middleware.Alow())
	groupSiswa.GET("/all", siswaCtrl.FindAllSiswaCtr)
	groupSiswa.GET("/", siswaCtrl.FindByIdSiswaCtr)
	groupSiswa.POST("", siswaCtrl.CreateSiswaCtr)
	groupSiswa.PUT("/", siswaCtrl.UpdateSiswaCtr)
	groupSiswa.DELETE("/", siswaCtrl.DeleteSiswaCtr)

	groupGuru := r.Group("/guru")
	groupGuru.OPTIONS("", middleware.Alow())
	groupGuru.GET("/all", guruCtrl.FindAllGuruCtr)
	groupGuru.GET("/", guruCtrl.FindByIdGuruCtr)
	groupGuru.POST("", guruCtrl.CreateGuruCtr)
	groupGuru.PUT("/", guruCtrl.UpdateGuruCtr)
	groupGuru.DELETE("/", guruCtrl.DeleteGuruCtr)

}
