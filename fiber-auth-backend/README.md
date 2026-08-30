# Fiber Auth API

Contoh backend Golang menggunakan framework **Fiber** untuk registrasi dan login user, dengan **PostgreSQL** (via GORM) dan **JWT** untuk autentikasi.

## Struktur Folder

```
fiber-auth/
├── main.go                      # Entry point
├── go.mod
├── .env.example
├── database/
│   └── database.go              # Koneksi DB + auto-migration
├── models/
│   └── user.go                  # Struct User & request/response
├── handlers/
│   └── auth_handler.go          # Logic register, login, me
├── middleware/
│   └── auth_middleware.go       # Middleware verifikasi JWT
├── routes/
│   └── routes.go                # Pendaftaran endpoint
└── utils/
    ├── jwt.go                   # Generate & parse JWT
    └── password.go              # Hash & compare password (bcrypt)
```

## Cara Menjalankan

1. **Install dependencies**
   ```bash
   go mod tidy
   ```

2. **Siapkan database PostgreSQL**, lalu copy env:
   ```bash
   cp .env.example .env
   ```
   Sesuaikan `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, dan ganti `JWT_SECRET` dengan string acak yang kuat.

3. **Jalankan server**
   ```bash
   go run main.go
   ```
   Server berjalan di `http://localhost:3000` (atau sesuai `PORT` di .env).

## Endpoint

### POST `/api/auth/register`
Registrasi user baru (otomatis login, langsung dapat token).

**Body:**
```json
{
  "name": "Fikri Sujud",
  "email": "fikri@example.com",
  "password": "rahasia123"
}
```

**Response (201):**
```json
{
  "token": "eyJhbGciOi...",
  "user": {
    "id": 1,
    "name": "Fikri Sujud",
    "email": "fikri@example.com",
    "created_at": "...",
    "updated_at": "..."
  }
}
```

### POST `/api/auth/login`
Login dengan email & password.

**Body:**
```json
{
  "email": "fikri@example.com",
  "password": "rahasia123"
}
```

**Response (200):** sama seperti register — `token` + `user`.

### GET `/api/auth/me`
Ambil data user yang sedang login. Perlu header:
```
Authorization: Bearer <token>
```

## Catatan Keamanan

- Password di-hash pakai **bcrypt** sebelum disimpan, tidak pernah disimpan plain text.
- Field `Password` di struct `User` diberi tag `json:"-"` agar tidak pernah ikut ter-serialize di response.
- Pesan error saat login sengaja dibuat generik ("Email atau password salah") agar tidak membocorkan apakah suatu email terdaftar atau tidak.
- `JWT_SECRET` **wajib** diganti dengan nilai yang kuat & rahasia sebelum production, jangan pernah commit `.env` ke git.
- `AllowOrigins: "*"` di CORS hanya untuk development — ganti dengan domain frontend spesifik saat production.

## Pengembangan Lanjutan (opsional)

- Tambahkan refresh token agar tidak perlu login ulang tiap token expired.
- Validasi input lebih ketat pakai library seperti `go-playground/validator`.
- Tambahkan rate limiting di endpoint login untuk mencegah brute force (Fiber punya middleware `limiter` bawaan).
- Kirim email verifikasi saat registrasi.
