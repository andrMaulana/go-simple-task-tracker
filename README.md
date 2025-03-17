# Task Tracker CLI 📝

Aplikasi Command-Line Interface (CLI) sederhana untuk melacak dan mengelola tugas sehari-hari. Dibangun dengan bahasa Go, aplikasi ini memungkinkan Anda menambahkan, memperbarui, menghapus, dan memantau status tugas dengan mudah.

---

## Fitur Utama ✨

- **Tambahkan Tugas**: Simpan tugas baru dengan deskripsi.
- **Perbarui Tugas**: Ubah deskripsi tugas yang sudah ada.
- **Hapus Tugas**: Hapus tugas yang tidak diperlukan.
- **Tandai Status**:
  - `todo`: Tugas belum dikerjakan.
  - `in-progress`: Tugas sedang dikerjakan.
  - `done`: Tugas selesai.
- **Filter Tampilan**: Tampilkan tugas berdasarkan status (`done`, `todo`, `in-progress`).
- **Penyimpanan Lokal**: Data tugas disimpan dalam file `tasks.json`.

---

## Cara Menggunakan 🛠️

### 1. Instalasi

Pastikan Anda telah menginstal Go (versi 1.20+).  
Clone repositori ini dan masuk ke direktori proyek:

```bash
git clone https://github.com/andrMaulana/go-simple-task-tracker.git
cd go-simple-task-tracker
```

### 2. Command Utama

- Menambahkan Tugas

```bash
go run cmd/task-cli/main.go add "Beli bahan makanan"
```

output:

```bash
Tugas berhasil ditambahkan (ID: 1)
```

- Mengupdate Tugas

```bash
go run cmd/task-cli/main.go update 1 "Beli bahan makanan dan masak malam"
```

output:

```bash
Tugas #1 berhasil diperbarui
```

- Menghapus Tugas

```bash
go run cmd/task-cli/main.go delete 1
```

output:

```bash
Tugas #1 berhasil dihapus
```

- Menandai Status Tugas

```bash
go run cmd/task-cli/main.go mark-in-progress 1  # Ubah status ke "in-progress"
go run cmd/task-cli/main.go mark-done 1        # Ubah status ke "done"
```

- Menampilkan Daftar Tugas

```bash
go run cmd/task-cli/main.go list          # Tampilkan semua tugas
go run cmd/task-cli/main.go list done     # Tampilkan tugas yang sudah selesai
go run cmd/task-cli/main.go list todo     # Tampilkan tugas yang belum dikerjakan
```

---

File `tasks.json`
File ini akan dibuat otomatis di direktori root proyek. Contoh struktur:

```json
{
  "tasks": [
    {
      "id": 1,
      "description": "Beli bahan makanan",
      "status": "in-progress",
      "createdAt": "2023-10-01T12:00:00Z",
      "updatedAt": "2023-10-01T12:00:00Z"
    }
  ]
}
```

---

Struktur Proyek 📂

```
task-tracker/
├── cmd/                # Entry point aplikasi
│   └── task-cli/
│       └── main.go     # CLI handler
├── internal/
│   ├── application/    # Business logic
│   ├── domain/         # Entitas dan model data
│   └── infrastructure/ # Implementasi penyimpanan (JSON)
├── tasks.json          # File penyimpanan data tugas
└── README.md           # Dokumentasi ini
```

---
