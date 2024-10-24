# Pelatihan KMTETI - Backend with Go

1. Buat sebuah file `.env`
```bash
MONGODB="mongodb_connection_string"
```
2. Install package yang dibutuhkan
```bash
go mod download
go mod tidy
```
3. Jalankan program
```bash
go run src/main.go
```

## Penjelasan Struktur
### 1. src/
Merupakan folder yang berisi kode program. Di dalamnya terdapat main function yang menjadi entry point dari program Go. 
### 2. db/db.go
Merupakan kode yang digunakan untuk menghubungkan server dengan database. Function di dalam db.go akan me-load `.env` yang berisi connection string ke MongoDB
### 3. handler/
Merupakan folder yang berisi file dengan handler function. Handler Function adalah fungsi khusus yang dijalankan jika sebuah API endpoint dipanggil oleh client.
### 4. model/
Merupakan folder yang berisi struktur dokumen dari MongoDB. Pendefinisian struct disesuaikan dengan field yang ada pada data di MongoDB.