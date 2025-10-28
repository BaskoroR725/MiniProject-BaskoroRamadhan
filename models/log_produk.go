package models

import "time"

type LogProduk struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	ProdukID      uint      `json:"produk_id"` // foreign key ke produk
	NamaProduk    string    `json:"nama_produk"`
	Slug          string    `json:"slug"`
	HargaReseller float64   `json:"harga_reseller"`
	HargaKonsumen float64   `json:"harga_konsumen"`
	Deskripsi     string    `json:"deskripsi"`
	TokoID        uint      `json:"toko_id"`
	CategoryID    uint      `json:"category_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Produk   Produk   `json:"produk"`   // relasi balik ke Produk
	Toko     Toko     `json:"toko"`
	Category Category `json:"category"`
}
