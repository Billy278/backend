package server

import (
	"backend/config"

	controller "backend/module/controller"
	repository "backend/module/repository"
	service "backend/module/service"
)

type Ctrs struct {
	GuruCtrl  controller.GuruController
	SiswaCtrl controller.SiswaController
}

func NewSetup() Ctrs {
	datastore := config.NewDBPostges()
	//siswa
	repoSiswa := repository.NewSiswaRepoImpl(datastore)
	servSiswa := service.NewSiswaServiceImpl(repoSiswa)
	ctrlSiswa := controller.NewSiswaContollerimpl(servSiswa)

	//guru
	repoGuru := repository.NewGuruRepoImpl(datastore)
	servGuru := service.NewGuruServiceImpl(repoGuru)
	ctrlGuru := controller.NewGuruControllerImpl(servGuru)

	return Ctrs{
		GuruCtrl:  ctrlGuru,
		SiswaCtrl: ctrlSiswa,
	}
}
