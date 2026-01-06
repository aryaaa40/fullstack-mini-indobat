Mini Indobat Backend
Aplikasi backend sederhana untuk manajemen produk dan transaksi apotek (obat) yang dibangun menggunakan Go, Gin Gonic, dan GORM. Aplikasi ini dilengkapi dengan mekanisme Database Locking untuk menangani race condition pada saat transaksi konkuren.

Fitur Utama
Master Produk: Kelola data obat, stok, dan harga.
- Transaksi Order: Pembelian obat dengan perhitungan diskon otomatis.
- Concurrency Safety: Menggunakan SELECT FOR UPDATE (Row Locking) untuk mencegah stok minus saat banyak request masuk bersamaan.
- Database Migration: Menggunakan GORM AutoMigrate.

Persiapan Database (PostgreSQL)
- Buat database baru dengan nama mini_indobat.
- Pastikan PostgreSQL Anda berjalan di port default 5432.
- Konfigurasi kredensial (username/password) akan diatur melalui file .env.

Instalasi dan Konfigurasi
Install Dependencies
- go mod tidy

Konfigurasi Environment
Buat file .env di root project dengan konfigurasi berikut:
- DB_HOST=localhost
- DB_PORT=5432
- DB_USER=postgres
- DB_PASS=arya123
- DB_NAME=mini_indobat

Jalankan Aplikasi
- go run main.go
- Aplikasi akan berjalan di http://localhost:8080. Database akan otomatis bermigrasi saat aplikasi dijalankan pertama kali.

Tech Stack
- Go: Programming language
- Gin Gonic: Web framework
- GORM: ORM library
- PostgreSQL: Database
