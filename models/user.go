package models

import "time"

type User struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	NamaUser      string    `json:"nama_user"`
	KataSandi     string    `json:"kata_sandi"`
	NoTelp        string    `json:"no_telp"`
	TanggalLahir  time.Time `json:"tanggal_lahir"`
	JenisKelamin  string    `json:"jenis_kelamin"`
	Email         string    `json:"email"`
	Pekerjaan     string    `json:"pekerjaan"`
	IDProvinsi    string    `json:"id_provinsi"`
	IDKota        string    `json:"id_kota"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Toko *Toko `json:"toko" gorm:"foreignKey:UserID"`

	Alamat     []Alamat     `json:"alamat,omitempty"`
	Transaksi  []Transaksi  `json:"transaksi,omitempty"`
}
