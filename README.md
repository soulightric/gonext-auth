# Deploy ke VPS — Fiber Auth + Next.js Frontend

Panduan ini mengikuti pola yang sudah kamu pakai di VPS (`ilkomith.cloud`): Docker multi-container + Nginx reverse proxy + Certbot.

## Struktur folder di VPS

Ikuti pola kamu yang sudah ada:

```
/home/projects/2026/nama-project/
├── fiber-auth/              # source code backend (copy dari project ini)
├── nextjs-auth-frontend/    # source code frontend (copy dari project ini)
├── docker-compose.yml
└── .env
```

## Langkah-langkah

### 1. Push code ke server
Copy folder `fiber-auth/` dan `nextjs-auth-frontend/` (yang sudah ada `Dockerfile` masing-masing) ke VPS, taruh sejajar dengan `docker-compose.yml` ini. Bisa pakai `git clone`, `scp`, atau `rsync`.

### 2. Siapkan .env
```bash
cp .env.example .env
nano .env
```
Isi dengan kredensial database asli dan `JWT_SECRET` yang kuat (generate misalnya dengan `openssl rand -base64 32`). Ganti juga `PUBLIC_API_URL` dan `PUBLIC_APP_URL` sesuai domain yang akan dipakai.

### 3. Build & jalankan stack
```bash
docker compose up -d --build
```
Ini akan menjalankan 3 container: `db` (PostgreSQL), `backend` (Fiber, port 3000 di localhost host), `frontend` (Next.js, port 3001 di localhost host). Ketiganya hanya bind ke `127.0.0.1`, **tidak langsung exposed ke publik** — akses publik hanya lewat Nginx.

### 4. Setup schema permission (sekali saja)
Container `db` fresh biasanya sudah otomatis owner sesuai `POSTGRES_USER`, jadi masalah `permission denied for schema public` yang kamu alami di local biasanya **tidak terjadi** di setup Docker ini (beda dari install native yang bikin role terpisah dari owner). Kalau tetap terjadi, exec ke container db dan jalankan grant yang sama seperti sebelumnya:
```bash
docker compose exec db psql -U <DB_USER> -d <DB_NAME> -c "GRANT ALL ON SCHEMA public TO <DB_USER>;"
```

### 5. Setup Nginx
Pakai contoh di `nginx-example.conf` — sesuaikan `server_name` dengan domain/subdomain project ini, lalu:
```bash
sudo cp nginx-example.conf /etc/nginx/sites-available/namaproject.ilkomith.cloud
sudo ln -s /etc/nginx/sites-available/namaproject.ilkomith.cloud /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

### 6. Aktifkan SSL
```bash
sudo certbot --nginx -d namaproject.ilkomith.cloud
```
Certbot otomatis update config Nginx untuk HTTPS dan redirect HTTP → HTTPS.

### 7. Verifikasi
```bash
curl https://namaproject.ilkomith.cloud/api/auth/login -X POST \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"salah"}'
```
Kalau balikin JSON error `"Email atau password salah"` (bukan error koneksi), berarti backend sudah bisa diakses lewat domain publik.

## Kenapa strukturnya begini

- **DB tidak exposed ke publik** — cuma bisa diakses dari network internal Docker (`internal`), jadi nggak bisa diserang langsung dari luar.
- **Backend & frontend bind ke `127.0.0.1` saja** — akses publik wajib lewat Nginx, sehingga Nginx bisa handle SSL, rate limiting, dan logging terpusat.
- **`NEXT_PUBLIC_API_URL` di-inject saat build (`ARG`)**, bukan runtime — karena Next.js meng-inline environment variable berprefix `NEXT_PUBLIC_` langsung ke JS bundle saat build time, bukan dibaca ulang saat container jalan.
- **`ALLOWED_ORIGIN` di backend** diarahkan ke domain publik frontend — kalau nanti ada mobile app juga manggil API ini, cukup ubah logic CORS di `main.go` untuk terima multiple origin, atau matikan CORS check untuk request tanpa header `Origin` (mobile app native biasanya tidak mengirim header ini).

## Update/redeploy

```bash
git pull   # atau update source code
docker compose up -d --build
```
Docker akan rebuild image yang berubah saja dan restart container terkait, minim downtime.

## Monitoring dasar

```bash
docker compose logs -f backend     # log backend real-time
docker compose logs -f frontend    # log frontend real-time
docker compose ps                  # status semua container
```
