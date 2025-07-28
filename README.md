# Real-time Notification Service

Layanan notifikasi real-time antara server dan client.  
Aplikasi ini memungkinkan server mengirim notifikasi ke client secara real-time melalui WebSocket. Jika client sedang offline, notifikasi akan disimpan di database dan dapat diambil kembali melalui REST API.

## Teknologi yang Digunakan
Golang dengan Gin Framework sebagai backend.  
PostgreSQL sebagai database utama (opsional SQLite untuk development).  
JWT (JSON Web Token) sebagai mekanisme autentikasi.  
Gorilla WebSocket untuk komunikasi real-time.  
GORM sebagai ORM untuk akses database.

## Langkah Instalasi dan Setup

Langkah 1: Pastikan prasyarat berikut sudah terpasang di sistem
- Golang versi minimal 1.21
- PostgreSQL atau SQLite
- Git


Langkah 2: Clone repository
`git clone https://github.com/username/notification-service.git`
`cd notification-service`


Langkah 3: Setup database
Jika menggunakan PostgreSQL, buat database baru dengan perintah:
`CREATE DATABASE notification_db;`

Lalu edit file konfigurasi `internal/config/config.go` agar sesuai dengan pengaturan database Anda:
DBHost: "localhost",
DBPort: "5432",
DBUser: "postgres",
DBPassword: "password",
DBName: "notification_db",

Jika menggunakan SQLite (opsional untuk mempermudah development), ubah kode di file `main.go` menjadi:
db, err := gorm.Open(sqlite.Open("notification.db"), &gorm.Config{})


Langkah 4: Jalankan migrasi database
Jika menggunakan PostgreSQL jalankan:
`psql -U postgres -d notification_db -f database/migrations/001_init_tables.sql`

Jika menggunakan SQLite, tabel akan dibuat otomatis ketika pertama kali server dijalankan.


Langkah 5: Install dependencies dan jalankan server
`go mod tidy`
`go run main.go`
Setelah server berjalan, aplikasi bisa diakses melalui http://localhost:8080

## Endpoint API

Endpoint Autentikasi:
- POST /api/register  digunakan untuk registrasi user baru
- POST /api/login     digunakan untuk login dan mendapatkan token JWT

Endpoint Notifikasi:
- GET /api/ws                        digunakan untuk membuka koneksi WebSocket agar bisa menerima notifikasi secara real-time
- POST /api/notifications/send       digunakan untuk mengirim notifikasi manual ke user
- GET /api/notifications             digunakan untuk mengambil daftar riwayat notifikasi

## Cara Pengujian Menggunakan curl atau Postman

1. Registrasi user
curl -X POST http://localhost:8080/api/register
-H "Content-Type: application/json"
-d '{"username": "testuser", "password": "password123"}'

2. Login dan ambil token JWT
curl -X POST http://localhost:8080/api/login
-H "Content-Type: application/json"
-d '{"username": "testuser", "password": "password123"}'
Simpan token JWT yang dikembalikan karena akan digunakan pada permintaan berikutnya.

3. Koneksi ke WebSocket
Gunakan Postman atau client WebSocket lainnya, lalu hubungkan ke:
`ws://localhost:8080/api/ws`
Tambahkan header:
Authorization: Bearer <TOKEN_JWT>

4. Kirim notifikasi ke user
curl -X POST http://localhost:8080/api/notifications/send
-H "Authorization: Bearer <TOKEN_JWT>"
-H "Content-Type: application/json"
-d '{"user_id": 1, "message": "Pesananmu sudah dikirim!"}'
Jika user dengan ID tersebut sedang online, notifikasi akan segera muncul melalui WebSocket. Jika offline, notifikasi akan disimpan di database.

5. Ambil notifikasi lama yang sudah tersimpan di database
curl -X GET http://localhost:8080/api/notifications
-H "Authorization: Bearer <TOKEN_JWT>"

## Troubleshooting

Jika mendapatkan error 401 Unauthorized:
- Pastikan token JWT masih valid dan belum kadaluarsa.
- Gunakan header Authorization dengan format: Authorization: Bearer <token>.

Jika notifikasi tidak muncul di WebSocket:
- Pastikan koneksi WebSocket sudah terbuka sebelum mengirim notifikasi.
- Periksa apakah notifikasi tersimpan di database dengan query:
`'SELECT * FROM notifications WHERE user_id = 1;'`

Jika ada error koneksi database:
- Pastikan PostgreSQL atau SQLite berjalan.
- Periksa kembali konfigurasi database pada file config.go.

## Contoh Penggunaan

Contoh use case:  
Sebuah aplikasi e-commerce mengirimkan notifikasi "Pesananmu sudah dikirim" kepada user. Jika user sedang online maka notifikasi akan dikirimkan langsung melalui WebSocket. Jika user sedang offline maka notifikasi disimpan di database, dan ketika user kembali online notifikasi dapat diambil kembali melalui endpoint REST API.
