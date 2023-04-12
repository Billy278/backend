package service

import (
	model "backend/module/model/guru"
	"context"
)

type GuruService interface {
	CreateGuru(ctx context.Context, GuruIn model.Guru) (Guru model.Guru, err error)
	UpdateGuru(ctx context.Context, GuruIn model.Guru) (Guru model.Guru, err error)
	FindByIdGuru(ctx context.Context, GuruId uint64) (Guru model.Guru, err error)
	FindAllGuru(ctx context.Context) (Guru []model.Guru, err error)
	DeleleByIdGuru(ctx context.Context, GuruId uint64) (err error)
}
