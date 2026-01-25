# Kasir API

Kasir API adalah backend RESTful API sederhana yang dibangun menggunakan Go (Golang) untuk mengelola data produk dan kategori, cocok untuk sistem kasir atau Point of Sales (POS) sederhana.

## 🚀 Fitur

- **Manajemen Produk (CRUD)**: Create, Read, Update, Delete data produk.
- **Manajemen Kategori (CRUD)**: Create, Read, Update, Delete data kategori.
- **Health Check**: Endpoint untuk memantau status server.
- **Swagger Documentation**: Dokumentasi API interaktif yang terintegrasi.
- **Centralized Logging**: Pencatatan log Info dan Error yang terstruktur.
- **Standardized Response**: Format respons JSON yang konsisten untuk sukses dan error.

## 🛠️ Teknologi yang Digunakan

- **Language**: [Go](https://go.dev/)
- **API Documentation**: [Swagger (Swaggo)](https://github.com/swaggo/swag)
- **Live Reload**: [Air](https://github.com/air-verse/air) (untuk pengembangan)

## 📂 Struktur Project

```
kasir-api/
├── docs/           # Generated Swagger documentation
├── dto/            # Data Transfer Objects (Request structs)
├── handlers/       # HTTP Handlers (Controllers)
├── models/         # Data Models (Structs & In-memory data)
├── pkg/            # Utility packages (Logger, Response Helper)
├── routes/         # Router setup
├── main.go         # Entry point application
└── go.mod          # Go module definitions
```

## 📦 Instalasi & Cara Menjalankan

### Prasyarat
- Go terinstall di komputer Anda.

### 1. Clone Repository
```bash
git clone https://github.com/username/kasir-api.git
cd kasir-api
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Menjalankan Server

#### Mode Biasa
```bash
go run main.go
```

#### Mode Development (dengan Hot Reload)
Pastikan Anda sudah menginstall `air`:
```bash
# Install Air (jika belum)
go install github.com/air-verse/air@latest

# Jalankan dengan Air
air
```

Server akan berjalan di `http://localhost:8080`.

## 📚 Dokumentasi API

Dokumentasi lengkap API tersedia melalui Swagger UI.
Setelah server berjalan, buka browser dan akses:

```
http://localhost:8080/swagger/index.html
```

### Endpoint Utama

**Product**
- `GET /api/product` - Ambil semua produk
- `GET /api/product/{id}` - Ambil produk berdasarkan ID
- `POST /api/product` - Tambah produk baru
- `PUT /api/product/{id}` - Update produk
- `DELETE /api/product/{id}` - Hapus produk

**Category**
- `GET /api/category` - Ambil semua kategori
- `GET /api/category/{id}` - Ambil kategori berdasarkan ID
- `POST /api/category` - Tambah kategori baru
- `PUT /api/category/{id}` - Update kategori
- `DELETE /api/category/{id}` - Hapus kategori

**Health**
- `GET /health` - Cek status server

## 📝 Format Request JSON

**Create/Update Product**
```json
{
  "name": "Nama Produk",
  "harga": 10000,
  "stock": 50
}
```

**Create/Update Category**
```json
{
  "name": "Minuman",
  "description": "Kategori untuk berbagai jenis minuman"
}
```

## 🤝 Kontribusi

Silakan buat *pull request* untuk berkontribusi pada project ini.

## 📄 Lisensi

[MIT License](LICENSE)
