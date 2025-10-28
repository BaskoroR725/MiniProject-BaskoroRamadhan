package models

import "time"

type Transaksi struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	AlamatPengiriman  uint      `json:"alamat_pengiriman"`
	HargaTotal        float64   `json:"harga_total"`
	Invoice           string    `json:"invoice"`
	MetodeBayar       string    `json:"metode_bayar"`
	UserID            uint      `json:"user_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	User       User        `json:"user"`
	DetailTransaksi []DetailTransaksi `json:"detail_transaksi"`
}
