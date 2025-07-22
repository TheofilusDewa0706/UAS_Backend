# Aplikasi CRUD Komik

Aplikasi ini adalah sistem manajemen komik berbasis web yang dibangun dengan **Golang (Gin)** untuk backend dan **ReactJS** untuk frontend. Aplikasi mendukung autentikasi, manajemen data komik, komentar, dan fitur real-time update menggunakan WebSocket.

---

## 🛠 Teknologi yang Digunakan

### Backend:
- Golang + Gin
- JWT Authentication
- WebSocket (untuk update stok real-time)
- Swagger (dokumentasi API)
- GORM (ORM untuk MySQL/PostgreSQL)

### Frontend:
- ReactJS + React Router
- Protected Routes berbasis Role (Admin/User)
- Context API untuk autentikasi

---

## 🔐 Role dan Hak Akses
- **Admin (Role ID: 1)**:
  - CRUD Komik
  - Melihat & Menghapus komentar
- **User (Role ID: 2)**:
  - Membeli komik (Checkout)
  - Membuat, mengedit, dan menghapus komentar sendiri

---

## 📁 Struktur Proyek (Backend)

- `/controllers`: Logika bisnis API (Login, Komik, Komentar)
- `/models`: Struktur tabel database (Komik, Komentar, User)
- `/routes`: Routing API dan middleware role
- `/config`: Koneksi database
- `/websocket`: Handler WebSocket untuk update stok komik
- `main.go`: Entry point server

---

## 🔄 Rute API Penting

### Komik
- `GET /komik` - Semua komik (Admin/User)
- `POST /komik` - Tambah komik (Admin)
- `PUT /komik/:id` - Ubah komik (Admin)
- `DELETE /komik/:id` - Hapus komik (Admin)
- `GET /komik/updates` - WebSocket update stok

### Komentar
- `GET /comments` - Lihat komentar
- `POST /comments` - Tambah komentar (User)
- `PUT /comments/:id` - Edit komentar (User)
- `DELETE /comments/:id` - Hapus komentar (User/Admin)

### Auth
- `POST /login` - Login dan mendapatkan JWT Token

---

## 🖥️ Rute Frontend

- `/` - Login Page
- `/admin` - Dashboard Admin (CRUD Komik)
- `/user` - Dashboard User (Beli komik, Komentar)
- `/checkout` - Halaman checkout (User)
- `*` - 404 Not Found

---

## 🚀 Menjalankan Aplikasi

### Backend
```bash
cd backend
go mod tidy
go run main.go
```

### Frontend
```bash
cd frontend
npm install
npm run dev
```

---

## ✉️ Kontak

Jika ada pertanyaan silakan hubungi tim pengembang di: **support@example.com**

---

Lisensi: MIT
