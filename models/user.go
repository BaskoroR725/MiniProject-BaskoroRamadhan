package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	NamaUser     string    `json:"nama_user"`
	KataSandi    string    `json:"-"` // disembunyikan dari JSON
	NoTelp       string    `gorm:"unique" json:"no_telp"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Email        string    `gorm:"unique" json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Role string `json:"role" gorm:"default:user"`

	Toko        Toko         `json:"toko"`
	Alamat      []Alamat     `json:"alamat"`
	Transaksi   []Transaksi  `json:"transaksi"`
}
