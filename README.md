
# 🚀 Registration Smart Reward

Sistem registrasi member untuk Smart Reward menggunakan Golang (Gin) + MongoDB + Next.js (App Router) + Tailwind CSS (FE).

---

## 📦 Struktur Project

```bash
registration-smart-reward/
├── BE/                  # Backend - Gin + MongoDB
│   ├── main.go
│   ├── internal/
│   └── seeder/
├── FE/                  # Frontend - Next.js App Router
│   ├── src/
│   ├── public/
│   └── package.json
```

---

## 🛠️ Backend (Go Gin)

### ✅ Prasyarat
- Golang ≥ 1.20
- MongoDB berjalan (lokal atau remote)
- `go` & `air` (opsional) untuk hot reload

### 📥 Install Dependency
```bash
cd BE
go mod tidy
```

### ⚙️ Setup `.env`
Buat file `.env` di root `BE/`:

```env
PORT=8080
MONGODB_URI=mongodb://localhost:27017
DB_NAME=registration_smart_reward
JWT_SECRET=supersecretkey
```

### 🌱 Jalankan Seeder Paket
```bash
go run seeder/paket/paket_seeder.go
```

### ▶️ Jalankan Server
```bash
go run main.go
```

Atau jika pakai air:
```bash
air
```

API akan berjalan di <http://localhost:8080>

---

## 💻 Frontend (Next.js + Tailwind CSS)

### ✅ Prasyarat
- Node.js ≥ 18.x
- Yarn (bukan pnpm)

### 📥 Install Dependency
```bash
cd FE
yarn install
```

### ▶️ Jalankan Frontend
```bash
yarn dev
```

Akses di browser:
```
http://localhost:3000
```

---

## ✨ Fitur yang Tersedia

- [x] Register Member Wizard (5 step)
- [x] Validasi & Alert (Formik + Yup + React Hot Toast)
- [x] Login + JWT Auth (token disimpan di localStorage)
- [x] Dashboard dengan:
  - Total Member Card
  - List Member (pagination + search)
  - Tambah Paket (modal)
  - Tree Graph Sponsor & Upline (dummy)
- [x] CORS sudah di-handle di backend

---

## 🧪 Dummy Login Credential

```bash
Username: admin
Password: admin123
```

---

## 📮 API Endpoint Penting

| Endpoint            | Method | Keterangan              |
|---------------------|--------|--------------------------|
| `/login`            | POST   | Login dan dapatkan token |
| `/member`           | POST   | Register member baru     |
| `/paket`            | GET    | Ambil list paket aktif   |
| `/api/members`      | GET    | List member (protected)  |
| `/api/paket`        | POST   | Tambah paket (protected) |

---

## 👨‍💻 Author
Built with ❤️ by Ali Syihab
