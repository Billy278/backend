package service

import (
	model "backend/module/model/siswa"
	"context"
)

type SiswaService interface {
	CreateSiswa(ctx context.Context, siswaIn model.Siswa) (siswa model.Siswa, err error)
	UpdateSiswa(ctx context.Context, siswaIn model.Siswa) (siswa model.Siswa, err error)
	FindByIdSiswa(ctx context.Context, SiswaId uint64) (siswa model.Siswa, err error)
	FindAllSiswa(ctx context.Context) (siswa []model.Siswa, err error)
	DeleleByIdSiswa(ctx context.Context, siswaId uint64) (err error)
}
