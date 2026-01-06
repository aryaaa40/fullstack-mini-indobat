# Mini Indobat Backend

Aplikasi backend sederhana untuk manajemen produk dan transaksi apotek (obat) yang dibangun menggunakan Go, Gin Gonic, dan GORM. Aplikasi ini dilengkapi dengan mekanisme Database Locking untuk menangani race condition pada saat transaksi konkuren.

## Fitur Utama

- **Master Produk**: Kelola data obat, stok, dan harga.
- **Transaksi Order**: Pembelian obat dengan perhitungan diskon otomatis.
- **Concurrency Safety**: Menggunakan SELECT FOR UPDATE (Row Locking) untuk mencegah stok minus saat banyak request masuk bersamaan.
- **Database Migration**: Menggunakan GORM AutoMigrate.

## Persiapan Database (PostgreSQL)

1. Buat database baru dengan nama `mini_indobat`.
2. Pastikan PostgreSQL Anda berjalan di port default `5432`.
3. Konfigurasi kredensial (username/password) akan diatur melalui file `.env`.

## Instalasi dan Konfigurasi

### Install Dependencies
```bash
go mod tidy
```

### Konfigurasi Environment

Buat file `.env` di root project dengan konfigurasi berikut:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=arya123
DB_NAME=mini_indobat
```

### Jalankan Aplikasi
```bash
go run main.go
```

Aplikasi akan berjalan di `http://localhost:8080`. Database akan otomatis bermigrasi saat aplikasi dijalankan pertama kali.

## API Endpoints

### 1. Get All Products
`GET /products`

<img width="1916" height="957" alt="Screenshot 2026-01-06 224317" src="https://github.com/user-attachments/assets/9d2e84ae-4817-4a4c-8a2d-d99cf39f51c3" />


### 2. Create Product
`POST /products`

<img width="1907" height="912" alt="Screenshot 2026-01-06 224254" src="https://github.com/user-attachments/assets/3087ece7-2c19-4eef-b236-2c777b431bb3" />


### 3. Create Order
`POST /orders`

<img width="1912" height="961" alt="Screenshot 2026-01-06 224826" src="https://github.com/user-attachments/assets/9d0ae7d2-1487-4d6a-b348-6661ec192c01" />


### 4. Concurrency Test
`Testing multiple concurrent requests`

<img width="1913" height="921" alt="Screenshot 2026-01-06 224337" src="https://github.com/user-attachments/assets/519e7326-cf76-4c76-93c9-b220e2c43f25" />


## Tech Stack

- **Go**: Programming language
- **Gin Gonic**: Web framework
- **GORM**: ORM library
- **PostgreSQL**: Database
