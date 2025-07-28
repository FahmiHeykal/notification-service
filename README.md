# Real-time Notification Service

Layanan notifikasi real-time antara server dan client.

Aplikasi ini memungkinkan server mengirim notifikasi ke client secara real-time melalui WebSocket. Jika client offline, notifikasi akan disimpan di database dan bisa diambil melalui REST API.

## Teknologi yang Digunakan
- Backend: Golang (Gin Framework)
- Database: PostgreSQL (bisa SQLite untuk development)
- Authentication: JWT (JSON Web Token)
- WebSocket: Gorilla WebSocket
- ORM: GORM

## Instalasi dan Setup

1. Prasyarat
- Golang 1.21+
- PostgreSQL atau SQLite
- Git

2. Clone Repository
git clone https://github.com/username/notification-service.git
cd notification-service

3. Setup Database

PostgreSQL:
CREATE DATABASE notification_db;

Update konfigurasi di internal/config/config.go:
DBHost: "localhost"
DBPort: "5432"
DBUser: "postgres"
DBPassword: "password"
DBName: "notification_db"

SQLite (opsional untuk development):
Ganti konfigurasi di main.go:
db, err := gorm.Open(sqlite.Open("notification.db"), &gorm.Config{})

4. Jalankan Migrasi Database
Jika menggunakan PostgreSQL:
psql -U postgres -d notification_db -f database/migrations/001_init_tables.sql

Jika menggunakan SQLite tabel akan dibuat otomatis saat pertama kali dijalankan.

5. Install Dependencies dan Jalankan Server
go mod tidy
go run main.go

Server akan berjalan di http://localhost:8080

## Endpoint API

Autentikasi
POST /api/register -> Registrasi user baru
POST /api/login -> Login dan mendapatkan token JWT

Notifikasi
GET /api/ws -> WebSocket untuk real-time notifications
POST /api/notifications/send -> Kirim notifikasi manual
GET /api/notifications -> Ambil riwayat notifikasi

## Cara Testing dengan Postman

1. Registrasi User
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}'

2. Login dan Dapatkan Token
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}'

Simpan token JWT yang diberikan.

3. Connect WebSocket
Buka WebSocket di Postman dengan URL:
ws://localhost:8080/api/ws

Tambahkan header:
Authorization: Bearer <TOKEN_JWT>
Klik Connect.

4. Kirim Notifikasi
curl -X POST http://localhost:8080/api/notifications/send \
  -H "Authorization: Bearer <TOKEN_JWT>" \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "message": "Pesananmu sudah dikirim!"}'

Jika user online notifikasi akan muncul di WebSocket.

5. Ambil Notifikasi Lama
curl -X GET http://localhost:8080/api/notifications \
  -H "Authorization: Bearer <TOKEN_JWT>"

## Troubleshooting

Error 401 Unauthorized
- Pastikan token JWT valid dan belum expired
- Gunakan format header: Authorization: Bearer <token>

Notifikasi tidak muncul di WebSocket
- Pastikan WebSocket terhubung sebelum mengirim notifikasi
- Cek apakah notifikasi tersimpan di database:
SELECT * FROM notifications WHERE user_id = 1;

Error database
- Pastikan PostgreSQL atau SQLite berjalan
- Periksa konfigurasi koneksi di config.go

## Contoh Penggunaan
Contoh use case: aplikasi e-commerce mengirim notifikasi "Pesananmu dikirim" ke user. Jika user offline, notifikasi akan disimpan dan bisa dibaca nanti.
