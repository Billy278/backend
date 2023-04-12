package module

import (
	model "backend/module/model/guru"
	"context"
	"database/sql"
	"errors"
	"time"
)

type GuruRepoImpl struct {
	DB *sql.DB
}

func NewGuruRepoImpl(db *sql.DB) GuruRepo {
	return &GuruRepoImpl{
		DB: db,
	}
}

func (g_repo *GuruRepoImpl) CreateGuru(ctx context.Context, GuruIn model.Guru) (Guru model.Guru, err error) {
	sql := "INSERT INTO data_guru(nip,name,alamat,no_hp,golongan,create_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING id"
	rows, err := g_repo.DB.QueryContext(ctx, sql, GuruIn.Nip, GuruIn.Name, GuruIn.Alamat, GuruIn.No_hp, GuruIn.Golongan, GuruIn.Create_at)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&Guru.Id)
		if err != nil {
			return
		}
	}
	return
}
func (g_repo *GuruRepoImpl) UpdateGuru(ctx context.Context, GuruIn model.Guru) (Guru model.Guru, err error) {
	sql := "UPDATE data_guru SET nip=$1,name=$2,alamat=$3,no_hp=$4,golongan=$5,update_at=$6 WHERE id=$7 RETURNING id"
	rows, err := g_repo.DB.QueryContext(ctx, sql, GuruIn.Nip, GuruIn.Name, GuruIn.Alamat, GuruIn.No_hp, GuruIn.Golongan, GuruIn.Update_at, GuruIn.Id)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&Guru.Id)
		if err != nil {
			return
		}
	}
	return
}
func (g_repo *GuruRepoImpl) FindByIdGuru(ctx context.Context, GuruId uint64) (Guru model.Guru, err error) {
	sql := "SELECT id,nip,name,alamat,no_hp,golongan FROM data_guru WHERE id=$1 AND delete_at IS NULL"
	rows, err := g_repo.DB.QueryContext(ctx, sql, GuruId)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&Guru.Id, &Guru.Nip, &Guru.Name, &Guru.Alamat, &Guru.No_hp, &Guru.Golongan)
		if err != nil {
			return
		}
		return
	} else {
		return Guru, errors.New("GURU NOT FOUND")
	}
}
func (g_repo *GuruRepoImpl) FindAllGuru(ctx context.Context) (Guru []model.Guru, err error) {
	sql := "SELECT id,nip,name,alamat,no_hp,golongan FROM data_guru WHERE delete_at IS NULL"
	rows, err := g_repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	gr := model.Guru{}
	for rows.Next() {
		rows.Scan(&gr.Id, &gr.Nip, &gr.Name, &gr.Alamat, &gr.No_hp, &gr.Golongan)
		Guru = append(Guru, gr)
	}
	return
}
func (g_repo *GuruRepoImpl) DeleleByIdGuru(ctx context.Context, GuruId uint64, update_at *time.Time) (err error) {
	sql := "UPDATE data_guru SET delete_at=$1 WHERE id=$2"
	_, err = g_repo.DB.ExecContext(ctx, sql, update_at, GuruId)
	if err != nil {
		return
	}
	return
}
