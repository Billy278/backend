package model

import "time"

type Guru struct {
	Id        uint64 `json:"id"`
	Nip       uint64 `json:"nip"`
	Name      string `json:"name"`
	Alamat    string `json:"alamat"`
	No_hp     string `json:"no_hp"`
	Golongan  string `json:"golongan"`
	Create_at *time.Time
	Update_at *time.Time
	Delete_at *time.Time
}
