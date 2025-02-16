## Level 1: CRUD Sederhana dengan MySQL atau POSTGRESQL

### Kasus: Buat aplikasi sederhana untuk mengelola daftar buku menggunakan Golang dan MySQL dan PostgreSQL.

## Spesifikasi:

### Buat tabel books dengan kolom:

```
id (primary key, auto-increment)
title (varchar, 255)
author (varchar, 255)
year (int)
```

### Implementasikan CRUD menggunakan Gin dan GORM:

```
Create: Tambah buku baru
Read: Tampilkan daftar semua buku
Update: Ubah data buku berdasarkan id
Delete: Hapus buku berdasarkan id
```

### Buat koneksi database menggunakan GORM

📌 Tantangan tambahan: Buat validasi sederhana, misalnya title dan author tidak boleh kosong.
