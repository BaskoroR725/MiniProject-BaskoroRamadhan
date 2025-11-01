# 🛍️ Evermos Mini Project - Backend Golang

<div align="center">

![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-00ACD7?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white)

**RESTful API E-Commerce Backend**  
Dibangun dengan Golang, Fiber Framework, GORM ORM, dan MySQL

---

**Proyek Akhir Virtual Internship**  
**Rakamin Academy x Evermos**  
Backend Developer Program 2025

Dibuat oleh: **Baskoro Ramadhan**

[📦 Demo](#) • [📖 Dokumentasi](#-dokumentasi-api-endpoint) • [🚀 Quick Start](#%EF%B8%8F-instalasi-dan-menjalankan-proyek)

---

</div>

## 📌 Tentang Proyek

Evermos Mini Project adalah sistem backend REST API untuk platform e-commerce sederhana yang mencakup fitur manajemen user, toko, produk, kategori, alamat, dan transaksi dengan sistem autentikasi JWT dan role-based access control.

Proyek ini dikembangkan mengikuti spesifikasi **Rakamin Evermos Postman Collection** dengan menerapkan **Clean Architecture** dan best practices dalam pengembangan backend modern.

### 🎯 Tujuan Pembelajaran

- Implementasi REST API dengan Golang
- Penerapan Clean Architecture & Modular Design
- Autentikasi & Otorisasi menggunakan JWT
- Role-Based Access Control (RBAC)
- Database Design & Relationship Management
- File Upload Handling
- Pagination & Filtering
- Integration dengan External API

---

## ✨ Fitur Utama

### 🔐 Autentikasi & Otorisasi

- **Register & Login** dengan JWT Authentication
- **Role Management** (`user`, `admin`)
- **Token-based** authorization
- **Password hashing** dengan bcrypt

### 👤 User Management

- View & update profil user
- Manajemen alamat pengiriman (CRUD)
- Validasi email & nomor telepon unik
- User isolation (tidak dapat mengakses data user lain)

### 🏪 Toko (Store Management)

- **Auto-create** toko saat user register
- Update nama dan foto toko (file upload)
- Public endpoint untuk list & detail toko
- Owner-only access untuk manajemen toko

### 🗂️ Kategori Produk (Admin Only)

- CRUD kategori produk
- **Admin-only** endpoint protection
- Data kategori publik untuk user

### 🛍️ Produk (Product Management)

- CRUD produk oleh pemilik toko
- **Upload foto produk** (multipart form-data)
- **Pagination** (`page`, `limit`)
- **Filtering** (`category_id`, `nama_produk`, `toko_id`)
- Automatic slug generation untuk SEO
- Log produk otomatis saat transaksi

### 💳 Transaksi

- Proses pembelian produk
- **Automatic stock reduction**
- Log produk snapshot saat transaksi
- Status tracking (`Menunggu Pembayaran`, `Dikirim`, `Selesai`)
- Transaction history per user

### 📍 Alamat & Wilayah Indonesia

- CRUD alamat pengiriman user
- **Integrasi API Wilayah Indonesia** (emsifa.com)
- Data provinsi & kota real-time
- Validasi alamat pengiriman

---

## 🏗️ Struktur Folder

```
evermos-mini/
│
├── config/
│   ├── database.go              # Konfigurasi koneksi database
│   └── seeder.go                # Seeder kategori awal
│
├── controllers/
│   ├── auth_controller.go       # Handler register & login
│   ├── user_controller.go       # Handler profil user
│   ├── toko_controller.go       # Handler manajemen toko
│   ├── produk_controller.go     # Handler CRUD produk
│   ├── kategori_controller.go   # Handler CRUD kategori
│   ├── alamat_controller.go     # Handler CRUD alamat
│   ├── transaksi_controller.go  # Handler transaksi
│   └── provcity_controller.go   # Handler API wilayah
│
├── middleware/
│   ├── jwt_middleware.go        # Middleware validasi JWT
│   └── admin_middleware.go      # Middleware validasi admin role
│
├── models/
│   ├── user.go                  # Model data user
│   ├── toko.go                  # Model data toko
│   ├── produk.go                # Model data produk
│   ├── category.go              # Model data kategori
│   ├── transaksi.go             # Model data transaksi
│   ├── detail_transaksi.go      # Model detail item transaksi
│   ├── log_produk.go            # Model log snapshot produk
│   └── alamat.go                # Model data alamat
│
├── routes/
│   ├── routes.go                # Router utama & setup
│   ├── auth_route.go            # Route autentikasi
│   ├── user_route.go            # Route user & alamat
│   ├── toko_route.go            # Route toko
│   ├── product_route.go         # Route produk
│   ├── kategori_route.go        # Route kategori
│   ├── trx_route.go             # Route transaksi
│   └── provcity_route.go        # Route provinsi & kota
│
├── utils/
│   ├── jwt_utils.go             # Utility JWT token
│   ├── hash_utils.go            # Utility password hashing
│   ├── authorization_utils.go   # Utility validasi User
│   ├── validator_utils.go       # Utility validasi input
│   └── log_produk_utils.go      # Utility log produk
│
├── uploads/                     # Folder penyimpanan file upload
│
├── .env                         # Konfigurasi environment(buat sendiri)
├── .gitignore                   # Git ignore file
├── go.mod                       # Go modules
├── go.sum                       # Go dependencies checksum
├── main.go                      # Entry point aplikasi
├── Rakamin....collection.json   # Postman Collection untuk test API
└── README.md                    # Dokumentasi proyek
```

---

## ⚙️ Instalasi dan Menjalankan Proyek

### 📋 Prerequisites

Pastikan Anda sudah menginstall:

- **Go** v1.22 atau lebih baru ([Download](https://go.dev/dl/))
- **MySQL** v8.0 atau lebih baru ([Download](https://dev.mysql.com/downloads/))
- **Postman** untuk testing API ([Download](https://www.postman.com/downloads/))
- **Git** ([Download](https://git-scm.com/downloads))

### 1️⃣ Clone Repository

```bash
git clone https://github.com/BaskoroR725/MiniProject-BaskoroRamadhan.git
cd MiniProject-BaskoroRamadhan
```

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

### 3️⃣ Setup Database

Buat database MySQL baru:

```sql
CREATE DATABASE evermos_mini;
```

### 4️⃣ Konfigurasi Environment

Buat file `.env` di root folder dengan isi:

```env
# Database Configuration
DB_USER=root
DB_PASS=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=evermos_mini

# Application Configuration
APP_PORT=8080

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_here_change_this
```

**⚠️ PENTING:** Ganti `JWT_SECRET` dengan string random yang aman!

### 5️⃣ Jalankan Aplikasi

```bash
go run main.go
```

Jika berhasil, akan muncul pesan:

```
 Server running on http://localhost:8080
 Database connected successfully
 Auto migration completed
```

Aplikasi siap digunakan di: **http://localhost:8080**

---

## 🧪 Testing API dengan Postman

### 📥 Import Postman Collection

1. Buka aplikasi **Postman**
2. Klik tombol **Import** di pojok kiri atas
3. Pilih tab **File**
4. Upload file `Rakamin Evermos Virtual Internship.postman_collection.json`
5. Collection akan muncul di sidebar

### 🔧 Setup Environment Variable

Buat environment baru di Postman:

1. Klik icon **⚙️ (gear)** di pojok kanan atas
2. Klik **Add** untuk environment baru
3. Beri nama: `Evermos Local`
4. Tambahkan variable:

| Variable | Initial Value           | Current Value              |
| -------- | ----------------------- | -------------------------- |
| `local`  | `http://localhost:8080` | `http://localhost:8080`    |
| `token`  | (kosongkan dulu)        | (akan diisi setelah login) |

5. **Save** dan pilih environment `Evermos Local`

### 📋 Urutan Testing yang Disarankan

1. **Auth - Register** (`POST /auth/register`)

   - Buat akun user baru
   - Toko otomatis dibuat

2. **Auth - Login** (`POST /auth/login`)

   - Login dengan kredensial yang dibuat
   - **Copy token JWT** dari response
   - **Paste token** ke environment variable `{{token}}`

3. **User - Get Profile** (`GET /user`)

   - Test endpoint protected dengan JWT

4. **Category - Get All** (`GET /category`)

   - Lihat daftar kategori yang sudah di-seed

5. **Toko - Get My Toko** (`GET /toko/my`)

   - Lihat toko yang auto-generated

6. **Produk - Create** (`POST /product`)

   - Tambah produk baru (gunakan form-data untuk upload foto)

7. **Alamat - Create** (`POST /user/alamat`)

   - Tambah alamat pengiriman

8. **Province & City** (`GET /provcity/listprovincies`)

   - Test integrasi API wilayah

9. **Transaksi - Create** (`POST /trx`)

   - Buat transaksi pembelian

10. **Transaksi - Get All** (`GET /trx`)
    - Lihat history transaksi

### 💡 Tips Penggunaan Postman

- ✅ Pastikan environment `Evermos Local` sudah dipilih
- ✅ Token JWT akan **expired** setelah beberapa jam, login ulang jika dapat error 401
- ✅ Untuk upload file, gunakan **form-data** dengan key `photo` (type: file)
- ✅ Header `token: Bearer {{token}}` otomatis ditambahkan jika environment sudah diset
- ✅ Gunakan **pagination** dengan query `?page=1&limit=10`
- ✅ Gunakan **filtering** dengan query `?nama_produk=sepatu&category_id=1`

---

## 📚 Dokumentasi API (Endpoint)

### 🔐 Authentication

| Method | Endpoint         | Deskripsi                               | Auth |
| ------ | ---------------- | --------------------------------------- | ---- |
| `POST` | `/auth/register` | Registrasi user baru (auto create toko) | ❌   |
| `POST` | `/auth/login`    | Login dan dapatkan JWT token            | ❌   |

**Request Body Register:**

```json
{
  "nama": "Baskoro test",
  "kata_sandi": "password123",
  "no_telp": "081234567890",
  "tanggal_lahir": "2000-01-01"
}
```

**Request Body Login:**

```json
{
  "no_telp": "081234567890",
  "kata_sandi": "password123"
}
```

---

### 👤 User Management

| Method | Endpoint | Deskripsi               | Auth   |
| ------ | -------- | ----------------------- | ------ |
| `GET`  | `/user`  | Ambil profil user login | ✅ JWT |
| `PUT`  | `/user`  | Update profil user      | ✅ JWT |

---

### 📍 Alamat

| Method   | Endpoint           | Deskripsi              | Auth   |
| -------- | ------------------ | ---------------------- | ------ |
| `GET`    | `/user/alamat`     | List semua alamat user | ✅ JWT |
| `GET`    | `/user/alamat/:id` | Detail alamat by ID    | ✅ JWT |
| `POST`   | `/user/alamat`     | Tambah alamat baru     | ✅ JWT |
| `PUT`    | `/user/alamat/:id` | Update alamat          | ✅ JWT |
| `DELETE` | `/user/alamat/:id` | Hapus alamat           | ✅ JWT |

---

### 🏪 Toko

| Method | Endpoint         | Deskripsi                  | Auth   |
| ------ | ---------------- | -------------------------- | ------ |
| `GET`  | `/toko`          | List semua toko (public)   | ❌     |
| `GET`  | `/toko/:id_toko` | Detail toko by ID (public) | ❌     |
| `GET`  | `/toko/my`       | Toko milik user login      | ✅ JWT |
| `PUT`  | `/toko/:id_toko` | Update toko (nama & foto)  | ✅ JWT |

**Note:** Update toko menggunakan **form-data** dengan key `photo` untuk upload foto.

---

### 🗂️ Kategori (Admin Only)

| Method   | Endpoint        | Deskripsi             | Auth     |
| -------- | --------------- | --------------------- | -------- |
| `GET`    | `/category`     | List semua kategori   | ❌       |
| `GET`    | `/category/:id` | Detail kategori by ID | ❌       |
| `POST`   | `/category`     | Tambah kategori baru  | ✅ Admin |
| `PUT`    | `/category/:id` | Update kategori       | ✅ Admin |
| `DELETE` | `/category/:id` | Hapus kategori        | ✅ Admin |

**Cara menjadi Admin:**

```sql
UPDATE users SET role = 'admin' WHERE id = 1;
```

---

### 🛍️ Produk

| Method   | Endpoint       | Deskripsi                                 | Auth   |
| -------- | -------------- | ----------------------------------------- | ------ |
| `GET`    | `/product`     | List produk (support pagination & filter) | ❌     |
| `GET`    | `/product/:id` | Detail produk by ID                       | ❌     |
| `POST`   | `/product`     | Tambah produk baru                        | ✅ JWT |
| `PUT`    | `/product/:id` | Update produk                             | ✅ JWT |
| `DELETE` | `/product/:id` | Hapus produk (soft delete)                | ✅ JWT |

**Query Parameters:**

- `?page=1&limit=10` - Pagination
- `?nama_produk=sepatu` - Filter by nama
- `?category_id=1` - Filter by kategori
- `?toko_id=1` - Filter by toko

**Upload Foto Produk:**

- Gunakan **form-data**
- Key: `photo`
- Type: `file`
- Format: JPG, PNG, JPEG
- Max size: 5MB

---

### 💳 Transaksi

| Method | Endpoint   | Deskripsi              | Auth   |
| ------ | ---------- | ---------------------- | ------ |
| `GET`  | `/trx`     | List transaksi user    | ✅ JWT |
| `GET`  | `/trx/:id` | Detail transaksi by ID | ✅ JWT |
| `POST` | `/trx`     | Buat transaksi baru    | ✅ JWT |

**Request Body Create Transaksi:**

```json
{
  "alamat_kirim_id": 1,
  "detail_transaksi": [
    {
      "produk_id": 1,
      "kuantitas": 2
    },
    {
      "produk_id": 2,
      "kuantitas": 1
    }
  ]
}
```

**Note:** Stok produk akan otomatis berkurang dan log produk tersimpan.

---

### 🌍 Provinsi & Kota (External API)

| Method | Endpoint                            | Deskripsi             | Auth |
| ------ | ----------------------------------- | --------------------- | ---- |
| `GET`  | `/provcity/listprovincies`          | List semua provinsi   | ❌   |
| `GET`  | `/provcity/listcities/:prov_id`     | List kota by provinsi | ❌   |
| `GET`  | `/provcity/detailprovince/:prov_id` | Detail provinsi by ID | ❌   |
| `GET`  | `/provcity/detailcity/:city_id`     | Detail kota by ID     | ❌   |

**Data Source:** https://www.emsifa.com/api-wilayah-indonesia/

---

## 🔒 Keamanan & Authorization

### JWT Authentication

- Setiap endpoint protected memerlukan header: `token: Bearer <token>`
- Token expired setelah 24 jam (configurable)
- Token berisi payload: `user_id`, `role`, `exp`

### Role-Based Access Control (RBAC)

| Role      | Akses                                                       |
| --------- | ----------------------------------------------------------- |
| **User**  | CRUD data pribadi (profil, alamat, toko, produk, transaksi) |
| **Admin** | User access + CRUD kategori produk                          |

### Data Isolation

- ✅ User tidak dapat melihat/edit data user lain
- ✅ User hanya dapat mengelola produk di toko miliknya
- ✅ User hanya dapat melihat transaksi miliknya sendiri
- ✅ Admin dapat mengelola semua kategori

### Validasi Data

- ✅ Email user harus **unik** (tidak boleh duplikat)
- ✅ Nomor telepon harus **unik**
- ✅ Password minimal **8 karakter**
- ✅ Validasi format email, no telepon, tanggal lahir
- ✅ Sanitasi input untuk mencegah SQL Injection

---

## 🧠 Catatan Teknis

### Log Produk pada Transaksi

Saat transaksi dibuat, sistem otomatis:

1. Menyimpan **snapshot produk** ke tabel `log_produks`
2. Data yang disimpan: harga, nama, deskripsi, foto (saat transaksi terjadi)
3. Berguna untuk tracking **harga historis**
4. Meskipun produk dihapus/diubah, log transaksi tetap valid

### Manajemen Stok Otomatis

Ketika transaksi berhasil dibuat:

- ✅ Stok produk **otomatis berkurang** sesuai kuantitas
- ✅ Validasi ketersediaan stok sebelum transaksi
- ✅ Stok tidak bisa negatif (ada validasi)
- ✅ Rollback otomatis jika ada error

### Slug Produk untuk SEO

- Slug produk **tidak diperbarui otomatis** saat nama diubah
- Didesain untuk **SEO stability** (URL produk tetap konsisten)
- Update slug harus dilakukan manual jika diperlukan

### Auto Increment ID

- ID tidak selalu berurutan karena MySQL auto_increment behavior
- Rollback atau gagal insert tidak mengurangi counter
- Ini normal dan tidak mempengarugi fungsi aplikasi

### File Upload

- File disimpan di folder `/uploads/`
- Database hanya menyimpan **path file**, bukan binary
- Format yang didukung: JPG, PNG, JPEG
- Maksimal ukuran: 5MB per file

---

## 🧩 Integrasi API Wilayah Indonesia

Proyek ini terintegrasi dengan API publik [emsifa.com](https://www.emsifa.com/api-wilayah-indonesia/) untuk data wilayah Indonesia real-time.

### Kegunaan

- Dropdown provinsi saat user input alamat pengiriman
- Validasi alamat pengiriman
- Integrasi dengan sistem ongkir/ekspedisi
- Data akurat dan selalu update

### Implementasi

Data diambil secara real-time tanpa perlu menyimpan ke database, sehingga:

- ✅ Selalu up-to-date
- ✅ Tidak memberatkan database
- ✅ Mudah maintenance

---

## 🧰 Teknologi yang Digunakan

| Teknologi         | Versi  | Kegunaan                        |
| ----------------- | ------ | ------------------------------- |
| **Golang**        | 1.22+  | Bahasa pemrograman utama        |
| **Fiber v2**      | Latest | Web framework (Express-like)    |
| **GORM**          | Latest | ORM untuk database management   |
| **MySQL**         | 8.0+   | Relational database             |
| **JWT**           | Latest | Token-based authentication      |
| **Validator v10** | Latest | Input validation                |
| **Godotenv**      | Latest | Environment variable management |
| **Bcrypt**        | Latest | Password hashing                |

---

## ⚡ Performa & Optimasi

### Database Optimization

- ✅ Gunakan **index** pada kolom yang sering difilter (`category_id`, `user_id`, `nama_produk`)
- ✅ Gunakan **Preload()** secara selektif untuk menghindari N+1 query
- ✅ Pagination dengan `LIMIT` dan `OFFSET` untuk dataset besar

### Application Optimization

- ✅ Tidak ada raw SQL query (aman dari SQL Injection)
- ✅ Connection pooling untuk database
- ✅ Middleware caching untuk response yang sering diakses
- ✅ File upload hanya simpan path, bukan binary

---

## 📖 Referensi & Sumber Belajar

### Official Documentation

- [Golang Documentation](https://go.dev/doc/)
- [Fiber Framework](https://docs.gofiber.io/)
- [GORM Guide](https://gorm.io/docs/)
- [JWT Introduction](https://jwt.io/introduction)

### External Resources

- [API Wilayah Indonesia](https://www.emsifa.com/api-wilayah-indonesia/)
- [Rakamin Evermos Collection](https://github.com/Fajar-Islami/go-example-cruid/blob/master/Rakamin%20Evermos%20Virtual%20Internship.postman_collection.json)
- [MySQL Configuration Guide](https://levelup.gitconnected.com/build-a-rest-api-using-go-mysql-gorm-and-mux-a02e9a2865ee)

---

## 🤝 Kontribusi

Proyek ini merupakan tugas akhir virtual internship dan tidak menerima kontribusi publik. Namun, Anda bebas untuk:

- 🍴 Fork repository ini
- 📝 Membuat issue untuk bug report
- 💡 Memberi saran improvement

---

## 📄 Lisensi

Proyek ini dibuat untuk keperluan edukasi dalam program **Virtual Internship Rakamin x Evermos**.  
Silakan gunakan sebagai referensi pembelajaran dengan tetap mencantumkan kredit.

---

## 🙏 Acknowledgments

Terima kasih kepada:

- **Rakamin Academy** - Penyelenggara program virtual internship yang luar biasa
- **Evermos** - Partner perusahaan yang memberikan kesempatan belajar
- **Mentor & Fasilitator** - Atas bimbingan, feedback, dan dukungan selama program
- **Community** - Teman-teman peserta internship yang saling support

---

## 📞 Kontak

**Baskoro Ramadhan**  
Junior Backend Developer

📧 Email: baskorowebdev@gmail.com  
🐙 GitHub: [@BaskoroR725](https://github.com/BaskoroR725)  
💼 LinkedIn: [linkedin.com/in/baskoro-ramadhan](https://linkedin.com/in/baskoro-ramadhan)

---

<div align="center">

### 🌟 Jika project ini bermanfaat, jangan lupa kasih ⭐ di repository!

**Made with ☕ by Baskoro Ramadhan**

_Rakamin x Evermos Virtual Internship 2025_

</div>
