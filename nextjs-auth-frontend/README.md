# Next.js Auth Frontend

Frontend sederhana (Next.js App Router + TypeScript) yang terintegrasi dengan backend **Fiber Auth API**.

## Struktur

```
nextjs-auth-frontend/
├── app/
│   ├── layout.tsx          # Root layout
│   ├── page.tsx            # Redirect otomatis ke /login atau /dashboard
│   ├── globals.css         # Styling global
│   ├── login/page.tsx      # Halaman login
│   ├── register/page.tsx   # Halaman registrasi
│   └── dashboard/page.tsx  # Halaman protected, fetch data dari /api/auth/me
├── lib/
│   └── api.ts              # Client API + helper token (localStorage)
├── package.json
├── tsconfig.json
└── .env.local.example
```

## Cara Menjalankan

1. **Pastikan backend Fiber sudah jalan** di `http://localhost:3000` (lihat README project `fiber-auth`).

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Setup environment**
   ```bash
   cp .env.local.example .env.local
   ```
   Isi default sudah mengarah ke `http://localhost:3000/api`, sesuaikan kalau backend jalan di port/host lain.

4. **Jalankan dev server**
   ```bash
   npm run dev
   ```
   Frontend jalan di **`http://localhost:3001`** (sengaja dibedakan dari backend yang di port 3000).

5. Buka `http://localhost:3001` di browser — otomatis diarahkan ke halaman login.

## Alur Autentikasi

1. User isi form di `/register` atau `/login`
2. Frontend `POST` ke backend Fiber (`/api/auth/register` atau `/api/auth/login`)
3. Token JWT yang dibalikin backend disimpan di `localStorage`
4. Redirect ke `/dashboard`
5. Dashboard fetch `GET /api/auth/me` dengan header `Authorization: Bearer <token>` untuk ambil data user
6. Tombol "Keluar" menghapus token dari `localStorage` dan redirect ke `/login`

## Catatan Penting

- **CORS**: backend Fiber sudah di-set untuk hanya mengizinkan origin `http://localhost:3001`. Kalau frontend dijalankan di port/domain lain, update `AllowOrigins` di `main.go` backend.
- **Penyimpanan token**: contoh ini pakai `localStorage` demi kesederhanaan. Untuk production, lebih aman pakai **httpOnly cookie** (butuh Next.js API route sebagai proxy ke backend, supaya token nggak bisa diakses lewat JavaScript/XSS).
- **Validasi token expired**: kalau `fetchMe` gagal (token invalid/expired), dashboard otomatis hapus token dan redirect ke login.
- Belum ada refresh token — kalau JWT expired (default 72 jam, diatur di backend), user harus login ulang manual.

## Pengembangan Lanjutan (opsional)

- Ganti penyimpanan token ke httpOnly cookie via Next.js Route Handler sebagai proxy backend.
- Tambahkan React Context/hook `useAuth()` agar status login bisa diakses di banyak komponen tanpa fetch berulang.
- Tambahkan middleware Next.js (`middleware.ts`) untuk proteksi route di level server, bukan cuma client-side check.
- Tambahkan form validation yang lebih baik (misalnya pakai `zod` + `react-hook-form`).
