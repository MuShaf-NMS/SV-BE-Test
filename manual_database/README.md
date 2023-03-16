# Membuat Database Article secara manual

---

Di bawah ini adalah cara membuat database `article` dan table `posts` secara manual di database mysql

## Membuat Database Article

Sebelum membuat database baru di mysql kita harus terkoneksi ke database mysql terlebih dahulu menggunakan perintah di bawah.

```
mysql -u root -p
```

kemudia ketikan password jika diperlukan.

Setelah itu, untuk membuat database `article` gunakan perintah di bawah.

```
CREATE DATABASE artice;
```

Setelah berhasil membuat database `article` sekarang kita bisa menambahkan table ke dalamnya.

## Membuat Table Posts

Sebelum membuat table baru kita harus pastikan telah memilih database mana yang akan kita gunakan. Dalam kasus ini kita akan menggunakan database `article`.

Gunakan perintah di bawah untuk memilih database `article`.

```
USE artice;
```

Setelah database terpilih, sekarang kita bisa membuat table `posts` dengan menggunakan perintah di bawah.

```
CREATE TABLE posts (id INT KEY AUTO_INCREMENT, title VARCHAR(200), content TEXT, category VARCHAR(100), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, status VARCHAR(100));
```

Perintah di atas akan membuat table `posts` dengan rincian kolom sebai berikut:

- kolom `id` bertipe data int, sebaga primary key dan auto increment (`id` akan otomatis bertambah ketika data baru di masukkan ke dalam table)
- kolom `title` bertipe data varchar dengan panjang karakter maksimal 200
- kolom `content` bertipe data text
- kolom `category` bertipe data varchar dengan panjang karakter maksimal 100
- kolom `created_at` bertipe data timestamp dengan default value current_timestamp
- kolom `updated_at` bertipe data timestamp dengan default value current_timestamp dan ketika diupdate akan menjadi current_timestamp
- kolom `status` bertipe data varchar dengan panjang karakter maksimal 100

## Hasil Dump

Hasil dump database `article` dapat dilihat di file `manual_article.sql`
