# 🛍️ Evermos Mini Project - Backend Golang

## 📌 Tentang Proyek

Proyek ini merupakan tugas akhir dari **Virtual Internship Backend Golang** yang diselenggarakan oleh **Rakamin** bekerjasama dengan **Evermos**.

Proyek ini adalah sistem backend REST API untuk aplikasi e-commerce sederhana yang mencakup fitur manajemen user, produk, kategori, toko, dan transaksi dengan autentikasi JWT.

**Dibuat oleh:** Baskoro Ramadhan  
**Program:** Virtual Internship Backend Golang - Evermos x Rakamin

---

## 📁 Struktur Folder Proyek

```
evermos-mini/
│
├── config/
│   └── config.go              # Konfigurasi koneksi database
│   └── seeder.go              # Konfigurasi seeder category database
│
├── controllers/
│   ├── auth_controller.go       # Handler autentikasi (register, login)
│   ├── produk_controller.go     # Handler CRUD produk
│   └── transaksi_controller.go  # Handler transaksi pembelian
│
├── middleware/
│   └── jwt_middleware.go        # Middleware validasi JWT token
│
├── models/
│   ├── user.go                  # Model data user
│   ├── toko.go                  # Model data toko
│   ├── category.go              # Model data kategori produk
│   ├── produk.go                # Model data produk
│   ├── log_produk.go            # Model log perubahan produk
│   ├── alamat.go                # Model data alamat
│   ├── transaksi.go             # Model data transaksi
│   └── detail_transaksi.go      # Model detail item transaksi
│
├── routes/
│   ├── routes.go                # Router utama
│   ├── auth_route.go            # Route autentikasi
│   ├── user_route.go            # Route user
│   ├── produk_route.go          # Route produk
│   └── transaksi_route.go       # Route transaksi
│
├── utils/
│   ├── hash_utils.go            # Utility hashing password
│   ├── jwt_utils.go             # Utility generate & validate JWT
│   ├── log_utils.go             # Utility logging sistem
│   └── validator_utils.go        # Utility validasi input
│
├── .env                         # File konfigurasi environment
├── go.mod                       # Dependencies management
├── go.sum                       # Checksum dependencies
├── main.go                      # Entry point aplikasi
└── README.md                    # Dokumentasi proyek
```

---

## ⚙️ Instalasi dan Menjalankan Proyek

### 1️⃣ Clone Repository

```bash
git clone https://github.com/BaskoroR725/MiniProject-BaskoroRamadhan.git
cd MiniProject-BaskoroRamadhan
```

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

### 3️⃣ Konfigurasi Database

Buat file `.env` di root folder proyek dengan isi:

```env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASS=
DB_NAME=evermos_mini
APP_PORT=:8080
JWT_SECRET=supersecretkey
```

**Catatan:** Sesuaikan nilai `DB_USER`, `DB_PASS`, dan `DB_NAME` dengan konfigurasi MySQL Anda.

### 4️⃣ Buat Database

Buat database MySQL dengan nama sesuai `.env` (default: `evermos_mini`):

```sql
CREATE DATABASE evermos_mini;
```

### 5️⃣ Jalankan Aplikasi

```bash
go run main.go
```

Aplikasi akan berjalan di `http://localhost:8080`

---

## 🧠 Fitur Utama

| Modul                | Fitur                 | Deskripsi                                                                       |
| -------------------- | --------------------- | ------------------------------------------------------------------------------- |
| **Auth**             | Register, Login       | Validasi input, hashing password dengan bcrypt, generate JWT Token              |
| **User**             | Get & Update Profile  | Mengambil dan memperbarui data profil user yang login                           |
| **Kategori**         | CRUD Kategori         | Mengelola kategori produk                                                       |
| **Toko**             | Manajemen Toko        | Setiap user dapat memiliki toko                                                 |
| **Produk**           | CRUD Produk           | Menambah, melihat, memperbarui, dan menghapus produk                            |
| **Log Produk**       | Logging Otomatis      | Tercatat otomatis setiap kali produk dibuat atau diubah                         |
| **Transaksi**        | CRUD Transaksi        | Membuat pesanan, mengurangi stok produk secara otomatis                         |
| **Security**         | Middleware JWT        | Validasi token untuk semua endpoint yang memerlukan autentikasi                 |
| **Soft/Hard Delete** | Penghapusan Fleksibel | Opsi menghapus produk beserta log (hard delete) atau hanya produk (soft delete) |

---

## 🔐 Autentikasi & Otorisasi

Sistem menggunakan **JWT (JSON Web Token)** untuk autentikasi:

1. User melakukan **register** dengan username, email, dan password
2. User melakukan **login** dan mendapatkan JWT token
3. Token digunakan di header `Authorization: Bearer <token>` untuk akses endpoint yang dilindungi
4. Middleware JWT akan memvalidasi token sebelum mengakses resource

---

## 🧪 Testing API dengan Postman

### Urutan Pengujian yang Disarankan:

1. **Auth - Register**  
   Buat akun user baru dengan endpoint `/api/auth/register`

2. **Auth - Login**  
   Login dengan kredensial yang sudah dibuat, salin token JWT yang diberikan

3. **Simpan Token**  
   Masukkan token ke Postman environment variable `{{token}}`

4. **Category - Tambah Kategori**  
   Buat kategori produk terlebih dahulu

5. **Toko - Buat Toko** (jika ada)  
   Buat toko untuk user yang login

6. **Produk - Create Produk**  
   Tambahkan produk baru dengan kategori yang sudah dibuat

7. **Produk - Get All / Get by ID**  
   Lihat daftar produk atau detail produk

8. **Produk - Update**  
   Ubah data produk (otomatis tercatat di log_produk)

9. **Transaksi - Create**  
   Buat transaksi pembelian (stok produk berkurang otomatis)

10. **Transaksi - Get**  
    Lihat riwayat transaksi

11. **Produk - Delete**  
    Hapus produk (hard delete menghapus produk beserta log-nya)

### 📥 Import Collection Postman

Gunakan file `EvermosMiniAPI_TestCollection.postman_collection.json` untuk import collection ke Postman. Untuk tes cepat API pada aplikasi ini. 

### 📥 Cara Import Postman Collection

1. Download/Clone repository ini
2. Buka Postman aplikasi
3. Klik tombol Import di pojok kiri atas
4. Pilih file `EvermosMiniAPI_TestCollection.postman_collection.json`
5. Collection akan muncul di sidebar Postman Anda

---

## 🗑️ Penghapusan Produk

### Hard Delete (Menghapus Produk & Log)

Menghapus produk beserta semua log perubahannya dari database.

### Soft Delete (Hanya Menghapus Produk, Log Tetap Ada)

Untuk menyimpan riwayat log produk meskipun produk sudah dihapus, gunakan fungsi alternatif berikut di `controllers/produk_controller.go`:

```go
func DeleteProdukOnly(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Unscoped().Delete(&models.Produk{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal menghapus produk tanpa log",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Produk berhasil dihapus (log tetap disimpan)",
	})
}
```

**Gunakan versi ini jika ingin menyimpan riwayat produk lama di tabel `log_produks`.**

---

## 🧰 Teknologi yang Digunakan

- **[Golang](https://go.dev/)** v1.22+ - Bahasa pemrograman utama
- **[Fiber v2](https://gofiber.io/)** - Web framework yang cepat dan ekspresif
- **[GORM](https://gorm.io/)** - ORM untuk Go yang powerful
- **[MySQL](https://www.mysql.com/)** - Database relasional
- **[JWT](https://jwt.io/)** - Token-based authentication
- **[Go Validator v10](https://github.com/go-playground/validator)** - Validasi input data
- **[godotenv](https://github.com/joho/godotenv)** - Load environment variables
- **Modular Architecture** - Struktur kode yang terorganisir dan scalable

---

## 📚 Dokumentasi API (Endpoint)

### Auth Endpoints

- `POST /api/auth/register` - Registrasi user baru
- `POST /api/auth/login` - Login dan dapatkan JWT token

### User Endpoints (Protected)

- `GET /api/user/profile` - Ambil data profil user
- `PUT /api/user/profile` - Update profil user

### Kategori Endpoints (Protected)

- `GET /api/category` - List semua kategori
- `POST /api/category` - Tambah kategori baru
- `GET /api/category/:id` - Detail kategori
- `PUT /api/category/:id` - Update kategori
- `DELETE /api/category/:id` - Hapus kategori

### Produk Endpoints (Protected)

- `GET /api/produk` - List semua produk
- `POST /api/produk` - Tambah produk baru
- `GET /api/produk/:id` - Detail produk
- `PUT /api/produk/:id` - Update produk
- `DELETE /api/produk/:id` - Hapus produk (hard delete)

### Transaksi Endpoints (Protected)

- `GET /api/transaksi` - List transaksi user
- `POST /api/transaksi` - Buat transaksi baru
- `GET /api/transaksi/:id` - Detail transaksi

---

## 🧩 Tahap Opsional: Integrasi API Wilayah Indonesia (Emsifa)

Proyek ini dapat diintegrasikan dengan API publik wilayah Indonesia untuk fitur alamat pengiriman.

**API Publik:** https://www.emsifa.com/api-wilayah-indonesia/

### Implementasi

Buat controller untuk mengambil data wilayah di `controllers/alamat_controller.go`:

```go
func GetProvinces(c *fiber.Ctx) error {
    resp, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data wilayah"})
    }
    defer resp.Body.Close()

    var provinces []map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&provinces)

    return c.JSON(fiber.Map{"status": true, "data": provinces})
}
```

Tambahkan route opsional di `routes/alamat_route.go`:

```go
func SetupAlamatRoutes(router fiber.Router) {
    router.Get("/provinsi", controllers.GetProvinces)
    router.Get("/kota/:provinceId", controllers.GetCities)
    router.Get("/kecamatan/:cityId", controllers.GetDistricts)
    router.Get("/kelurahan/:districtId", controllers.GetVillages)
}
```

### Kegunaan

Endpoint ini dapat digunakan untuk:

- Menampilkan dropdown provinsi saat user mengisi alamat pengiriman
- Validasi alamat pengiriman
- Integrasi dengan sistem ongkir/ekspedisi
- Melengkapi data toko atau user

---

## 📝 Catatan Pengembangan

### Fitur Log Produk

Setiap kali produk dibuat atau diubah, sistem otomatis mencatat perubahan di tabel `log_produks`. Log ini berguna untuk:

- Tracking perubahan harga
- Audit trail produk
- History perubahan data produk

### Manajemen Stok Otomatis

Ketika transaksi dibuat, sistem akan:

1. Validasi ketersediaan stok produk
2. Mengurangi stok produk secara otomatis
3. Mencatat detail transaksi

---

## 🧠 Catatan Teknis

### Slug Produk

- **Slug produk tidak diperbarui otomatis** saat nama produk diubah
- Ini didesain untuk **SEO stability** agar URL produk tetap konsisten
- Jika ingin update slug, harus dilakukan manual

### Auto Increment ID

- **ID produk tidak selalu berurutan** karena rollback atau gagal insert tidak mengurangi `auto_increment`
- Ini adalah behavior normal MySQL dan tidak mempengaruhi fungsi aplikasi

### Stok Produk

- **Transaksi otomatis mengurangi stok** produk yang dibeli
- Pastikan validasi stok dilakukan sebelum transaksi dibuat
- Stok tidak bisa negatif (ada validasi)

### Log Produk pada Transaksi

- Tabel `log_produks` menyimpan **snapshot produk** saat transaksi berlangsung
- Berguna untuk tracking harga historis
- Meskipun produk dihapus, log transaksi tetap tersimpan

---

## 🤝 Kontribusi

Proyek ini merupakan tugas akhir virtual internship. Untuk saran dan masukan, silakan hubungi:

**Baskoro Ramadhan**  
Email: [baskorowebdev@gmail.com]  

---

## 📄 Lisensi

Proyek ini dibuat untuk keperluan edukasi dalam program Virtual Internship Rakamin x Evermos.

---

## 🙏 Acknowledgments

Terima kasih kepada:

- **Rakamin Academy** - Penyelenggara program virtual internship
- **Evermos** - Partner perusahaan yang memberikan kesempatan belajar
- **Mentor & Fasilitator** - Atas bimbingan selama program berlangsung

---

**Made by Baskoro Ramadhan**
