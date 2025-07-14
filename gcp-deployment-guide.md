# Panduan Deployment Aplikasi di Google Cloud Platform (GCP)

Berikut adalah langkah-langkah lengkap untuk men-deploy aplikasi Anda di Google Cloud VM menggunakan Docker dan Docker Compose.

## Langkah 1: Persiapan di Komputer Lokal

Satu-satunya hal yang perlu Anda siapkan adalah **Git Repository**. Pastikan Anda memiliki repository (misalnya di GitHub) yang berisi file-file berikut:

1.  `docker-compose.yml` (yang akan kita siapkan)
2.  `mongod.conf` (yang akan kita siapkan)
3.  Direktori `.env` untuk setiap layanan jika diperlukan (misalnya, `shopping-service.env`, `payment-service.env`). **PENTING**: Jangan push file `.env` langsung ke Git jika berisi data sensitif. Sebaiknya, buat file-file ini secara manual di VM nanti.

## Langkah 2: Membuat dan Mengkonfigurasi VM di GCP

1.  **Buka Google Cloud Console** dan navigasi ke **Compute Engine** > **VM instances**.
2.  Klik **CREATE INSTANCE**.
3.  **Beri nama** VM Anda (misalnya, `app-server`).
4.  **Pilih Region dan Zone** yang paling dekat dengan pengguna Anda.
5.  **Pilih Machine type**. `e2-medium` adalah pilihan awal yang baik.
6.  Pada bagian **Boot disk**, pastikan sistem operasinya adalah **Debian** atau **Ubuntu**.
7.  Pada bagian **Firewall**, centang **Allow HTTP traffic** dan **Allow HTTPS traffic**. Ini akan membuka port 80 dan 443.
8.  Klik **Create**.

## Langkah 3: Terhubung ke VM dan Instalasi

1.  Setelah VM dibuat, klik tombol **SSH** untuk membuka terminal di browser.
2.  **Update package list**:
    ```bash
    sudo apt-get update
    ```
3.  **Install Git, Docker, dan Docker Compose**:

    ```bash
    # Install Git
    sudo apt-get install -y git

    # Install Docker
    sudo apt-get install -y docker.io
    sudo usermod -aG docker $USER

    # Install Docker Compose
    sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
    ```

4.  **PENTING**: Terapkan perubahan grup Docker. Anda bisa **log out dan log in kembali** ke SSH, atau jalankan perintah `newgrp docker` di sesi terminal Anda saat ini.

## Langkah 4: Clone Project dan Setup Konfigurasi

1.  **Clone repository Anda** ke dalam VM:

    ```bash
    git clone https://github.com/username/repository-name.git
    cd repository-name
    ```

    Perintah ini akan secara otomatis membuat semua file yang dibutuhkan, termasuk `docker-compose.yml` dan direktori `mongo-config`.

2.  **(Opsional tapi Direkomendasikan) Buat file `.env`**: Jika aplikasi Anda membutuhkan variabel lingkungan yang sensitif, buat file-file tersebut sekarang.
    ```bash
    # Contoh untuk shopping-service
    nano ./shopping-service/.env
    ```
    Isi file tersebut dengan variabel yang dibutuhkan, lalu simpan (`CTRL+X`, `Y`, `Enter`).

## Langkah 5: Jalankan Aplikasi

1.  Sekarang Anda berada di direktori yang benar dan semua file konfigurasi sudah ada. Jalankan Docker Compose:

    ```bash
    # Perintah ini akan menarik (pull) image dari Docker Hub dan menjalankan semua container
    docker-compose up -d
    ```

2.  **Verifikasi**: Cek apakah semua container berjalan dengan baik:
    ```bash
    docker-compose ps
    ```
    Anda seharusnya melihat semua layanan dalam keadaan `Up`.

Aplikasi Anda sekarang sudah berjalan di Google Cloud dan dapat diakses melalui IP eksternal VM Anda.
