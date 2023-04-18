package model

import "time"

type Siswa struct {
	Id        uint64 `json:"id"`
	Nim       uint64 `json:"nim"`
	Name      string `json:"name"`
	Alamat    string `json:"alamat"`
	Jenkel    string `json:"jenkel"`
	Kelas     string `json:"kelas"`
	Create_at *time.Time
	Update_at *time.Time
	Delete_at *time.Time
}
