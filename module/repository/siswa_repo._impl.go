package module

import (
	model "backend/module/model/siswa"
	"context"
	"database/sql"
	"errors"
	"time"
)

type SiswaRepoImpl struct {
	DB *sql.DB
}

func NewSiswaRepoImpl(db *sql.DB) SiswaRepo {
	return &SiswaRepoImpl{
		DB: db,
	}
}

func (s_repo *SiswaRepoImpl) CreateSiswa(ctx context.Context, siswaIn model.Siswa) (siswa model.Siswa, err error) {
	sql := "INSERT INTO data_siswa(nim,name,alamat,jenkel,kelas,create_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING id"
	rows, err := s_repo.DB.QueryContext(ctx, sql, siswaIn.Nim, siswaIn.Name, siswaIn.Alamat, siswaIn.Jenkel, siswaIn.Kelas, siswaIn.Create_at)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&siswa.Id)
		if err != nil {
			return
		}
	}
	return
}
func (s_repo *SiswaRepoImpl) UpdateSiswa(ctx context.Context, siswaIn model.Siswa) (siswa model.Siswa, err error) {
	sql := "UPDATE data_siswa SET nim=$1,name=$2,alamat=$3,jenkel=$4,kelas=$5,update_at=$6 WHERE id=$7 RETURNING id"
	rows, err := s_repo.DB.QueryContext(ctx, sql, siswaIn.Nim, siswaIn.Name, siswaIn.Alamat, siswaIn.Jenkel, siswaIn.Kelas, siswaIn.Update_at, siswaIn.Id)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&siswa.Id)
		if err != nil {
			return
		}
	}
	return
}
func (s_repo *SiswaRepoImpl) FindByIdSiswa(ctx context.Context, SiswaId uint64) (siswa model.Siswa, err error) {
	sql := "SELECT id,nim,name,alamat,jenkel,kelas FROM data_siswa WHERE id=$1 AND delete_at IS NULL"
	rows, err := s_repo.DB.QueryContext(ctx, sql, SiswaId)
	if err != nil {
		return
	}
	if rows.Next() {
		err = rows.Scan(&siswa.Id, &siswa.Nim, &siswa.Name, &siswa.Alamat, &siswa.Jenkel, &siswa.Kelas)
		if err != nil {
			return
		}
		return
	} else {
		return siswa, errors.New("SISWA NOT FOUND")
	}
}
func (s_repo *SiswaRepoImpl) FindAllSiswa(ctx context.Context) (siswa []model.Siswa, err error) {
	sql := "SELECT id,nim,name,alamat,jenkel,kelas FROM data_siswa WHERE delete_at IS NULL"
	rows, err := s_repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	sw := model.Siswa{}
	for rows.Next() {
		rows.Scan(&sw.Id, &sw.Nim, &sw.Name, &sw.Alamat, &sw.Jenkel, &sw.Kelas)
		siswa = append(siswa, sw)
	}
	return
}
func (s_repo *SiswaRepoImpl) DeleleByIdSiswa(ctx context.Context, siswaId uint64, update_at *time.Time) (err error) {
	sql := "UPDATE data_siswa SET delete_at=$1 WHERE id=$2"
	_, err = s_repo.DB.ExecContext(ctx, sql, update_at, siswaId)
	if err != nil {
		return
	}
	return
}
