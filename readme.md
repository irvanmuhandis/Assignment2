# Assignment 2 Golang Class

Ini adalah tugas kedua untuk kelas Golang yang diselenggarakan Hacktiv8 Indonesia.

## Deskripsi

Tugas ini mengenai pembuatan REST API dengan menggunakan bahasa golang. Pada projek ini terdapat 2 tabel, yaitu Order dan Item dimana terdapat relasi One to Many diantara keduanya. Relasi tersebut adalah satu Order bisa berisi banyak Item. Adapun API ini memiliki requirement :
- Create : Membuat Order
- Get : Melihat Order disertai Item nya (Eagerload)
- Update : Merubah Order beserta Item nya 
- Delete : Menghapus Order berdasarkan id nya

[Untuk selengkapnya bisa klik link ini](https://www.kode.id/courses/take/scalable-web-service-with-golang-kominfo/pdfs/38528096-assignment-2?authuser=1)


## Teknologi yang dipakai

- Gin
- GORM
- Golang
- Swagger
- Database Postgress

## Cara Menjalankan
1. Buka direktori projek
2. Pastikan Konfigurasi postgres sesuai dengan konfigurasi dibawah ini :
   ```
	const (
		host     = "localhost"
		port     = "5432"
		user     = "postgres"
		password = "1234"
		dbname   = "postgres"
	)
	```
	Bila berbeda maka ubah code di ``database/db.go``
3. Buka PgAdmin
4. Masuk server dan isikan password agar koneksi pada database bisa terhubung
5. Jalankan aplikasi dengan perintah
   ```
	go run main.go
	// atau
	go run .
   ```
6. Buka dokumentasi swagger di http://localhost:8080/swagger/index.html