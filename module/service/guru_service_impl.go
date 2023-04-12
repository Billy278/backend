package service

import (
	model "backend/module/model/guru"
	repository "backend/module/repository"
	"context"
	"log"
	"time"
)

type GuruServiceImpl struct {
	GuruRepo repository.GuruRepo
}

func NewGuruServiceImpl(gururepo repository.GuruRepo) GuruService {
	return &GuruServiceImpl{
		GuruRepo: gururepo,
	}
}

func (g_serv *GuruServiceImpl) CreateGuru(ctx context.Context, GuruIn model.Guru) (Guru model.Guru, err error) {
	tNow := time.Now()
	GuruIn.Create_at = &tNow
	Guru, err = g_serv.GuruRepo.CreateGuru(ctx, GuruIn)
	if err != nil {
		log.Printf("[ERROR] error Create Guru :%v\n", err)
	}
	return
}
func (g_serv *GuruServiceImpl) UpdateGuru(ctx context.Context, GuruIn model.Guru) (Guru model.Guru, err error) {
	Guru, err = g_serv.FindByIdGuru(ctx, GuruIn.Id)
	if err != nil {
		return
	}
	Guru.Nip = GuruIn.Nip
	Guru.Name = GuruIn.Name
	Guru.Alamat = GuruIn.Alamat
	Guru.No_hp = GuruIn.No_hp
	Guru.Golongan = GuruIn.Golongan
	tNow := time.Now()
	Guru.Update_at = &tNow

	Guru, err = g_serv.GuruRepo.UpdateGuru(ctx, Guru)
	if err != nil {
		log.Printf("[ERROR] error Update Guru :%v\n", err)
	}
	return

}
func (g_serv *GuruServiceImpl) FindByIdGuru(ctx context.Context, GuruId uint64) (Guru model.Guru, err error) {
	Guru, err = g_serv.GuruRepo.FindByIdGuru(ctx, GuruId)
	if err != nil {
		log.Printf("[ERROR] error find Guru :%v\n", err)
		return
	}
	return
}
func (g_serv *GuruServiceImpl) FindAllGuru(ctx context.Context) (Guru []model.Guru, err error) {
	Guru, err = g_serv.GuruRepo.FindAllGuru(ctx)
	if err != nil {
		log.Printf("[ERROR] error find Guru :%v\n", err)
		return
	}
	return
}
func (g_serv *GuruServiceImpl) DeleleByIdGuru(ctx context.Context, GuruId uint64) (err error) {
	tNow := time.Now()
	Update_at := &tNow
	err = g_serv.GuruRepo.DeleleByIdGuru(ctx, GuruId, Update_at)
	if err != nil {
		log.Printf("[ERROR] error delete Guru :%v\n", err)
		return
	}
	return
}
