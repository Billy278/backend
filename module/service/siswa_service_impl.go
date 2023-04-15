package service

import (
	model "backend/module/model/siswa"
	repository "backend/module/repository"
	"context"
	"log"
	"time"
)

type SiswaServiceImpl struct {
	SiswaRepo repository.SiswaRepo
}

func NewSiswaServiceImpl(siswarepo repository.SiswaRepo) SiswaService {
	return &SiswaServiceImpl{
		SiswaRepo: siswarepo,
	}
}

func (s_serv *SiswaServiceImpl) CreateSiswa(ctx context.Context, siswaIn model.Siswa) (siswa model.Siswa, err error) {

	tNow := time.Now()
	siswaIn.Create_at = &tNow
	siswa, err = s_serv.SiswaRepo.CreateSiswa(ctx, siswaIn)
	if err != nil {
		log.Printf("[ERROR] error Create Siswa :%v\n", err)
	}
	return
}
func (s_serv *SiswaServiceImpl) UpdateSiswa(ctx context.Context, siswaIn model.Siswa) (siswa model.Siswa, err error) {
	siswa, err = s_serv.SiswaRepo.FindByIdSiswa(ctx, siswaIn.Id)
	if err != nil {
		return siswa, err
	}
	siswa.Nim = siswaIn.Nim
	siswa.Name = siswaIn.Name
	siswa.Alamat = siswaIn.Alamat
	siswa.Jenkel = siswaIn.Jenkel
	siswa.Kelas = siswaIn.Kelas
	tNow := time.Now()
	siswa.Update_at = &tNow

	siswa, err = s_serv.SiswaRepo.UpdateSiswa(ctx, siswa)
	if err != nil {
		log.Printf("[ERROR] error Update Siswa :%v\n", err)
		return
	}
	return siswa, err

}
func (s_serv *SiswaServiceImpl) FindByIdSiswa(ctx context.Context, SiswaId uint64) (siswa model.Siswa, err error) {
	siswa, err = s_serv.SiswaRepo.FindByIdSiswa(ctx, SiswaId)
	if err != nil {
		log.Printf("[ERROR] error find siswa :%v\n", err)
		return
	}
	return
}
func (s_serv *SiswaServiceImpl) FindAllSiswa(ctx context.Context) (siswa []model.Siswa, err error) {
	siswa, err = s_serv.SiswaRepo.FindAllSiswa(ctx)
	if err != nil {
		log.Printf("[ERROR] error find siswa :%v\n", err)
		return
	}
	return
}
func (s_serv *SiswaServiceImpl) DeleleByIdSiswa(ctx context.Context, siswaId uint64) (err error) {
	tNow := time.Now()
	Update_at := &tNow
	err = s_serv.SiswaRepo.DeleleByIdSiswa(ctx, siswaId, Update_at)
	if err != nil {
		log.Printf("[ERROR] error delete Siswa :%v\n", err)
		return
	}
	return
}
