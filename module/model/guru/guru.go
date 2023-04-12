package model

import "time"

type Guru struct {
	Id        uint64
	Nip       uint64
	Name      string
	Alamat    string
	No_hp     string
	Golongan  string
	Create_at *time.Time
	Update_at *time.Time
	Delete_at *time.Time
}
