# Membuat Database Article dengan migrate dan api-service

---

Disini saya menggunakan <a href="https://gorm.io/">Gorm</a> sebagai ORM dan <a href="https://gin-gonic.com/">gin</a> sebagai frameworknya. Untuk model/entity `Post` didefinisikan di directory `entity`, kemudian dimigrate menggunakan `AutoMigrate`. Kemudian saya membuat `repository` sebagai jembatan antara database dengan `service`, sehingga `service` hanya fokus ke business logic dan tidak perlu melakukan query langsung ke database. Ketika ada `request` masuk, `handler` akan menentukan `service` mana yang akan digunakan. Sebelum aplikasi dijalankan `handler` harus diregstrasikan dengan `route`. Di sini saya menggunakan fungsi `InitializeRoute` untuk meregister `handler` pada `route` yang telah didefinisikan. `DTO` sebagai media transfer data dari `request body` ke `service`. Di `DTO` ini juga di definisikan validasi menggunakan tag <a href="https://pkg.go.dev/github.com/go-playground/validator/v10">Validate<a/> yang akan divalidasi di `hanlder` sebelum diteruskan ke `service` tujuan.
Sebelum menjalankan aplikasi ini, kita harus membuat file `.env` terlebih dahulu. Contoh file `.env` bisa dilihat di file `example.env`. Saya menggunakan <a href="https://pkg.go.dev/github.com/joho/godotenv">godotenv</a> untuk memuat file `.env`. `Variable` yang dimuat dari file `.env` akan disimpan sebagai config. Aplikasi dijalankan menggunakan fungsi `server.Run` disertai dengan `config.App_Port` sebagai parameter.

Perintah untuh menjalankan aplikasi:

```
go run .
```
