package model

import "time"

type Siswa struct {
	Id        uint64
	Nim       uint64
	Name      string
	Alamat    string
	Jenkel    string
	Kelas     string
	Create_at *time.Time
	Update_at *time.Time
	Delete_at *time.Time
}
