[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-2e0aaae1b6195c2367325f4f02e2d04e9abb55f0b24a779b69b11b9e10269abc.svg)](https://classroom.github.com/online_ide?assignment_repo_id=19928596&assignment_repo_type=AssignmentRepo)
[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/GMrD03Jz)

# Graded Challenge 1 - P3

Graded Challenge ini dibuat guna mengevaluasi pembelajaran pada Hacktiv8 Program Fulltime Golang khususnya pada pembelajaran MongoDB dan implementasi microservice Golang.

---

## Assignment Objectives

Graded Challenge 1 ini dibuat guna mengevaluasi pemahaman MongoDB dan Microservice sebagai berikut:

- Mampu memahami konsep Microservice
- Mampu memahami konsep MongoDB
- Mampu mengimplementasikan MongoDB ke REST API Golang

---

## Assignment Directions: RESTful API

Buatlah sebuah **REST API** menggunakan **Echo Golang** dan implementasikan database **MongoDB** sesuai kriteria berikut:

### Microservice yang harus dibangun:

#### **1. Payment Service**

- Endpoint:
  - `/payments` (POST) - Create a new payment

#### **2. Shopping Service**

- Endpoints:
  - `/transactions` (POST) - Create a new transaction (harus call Payment Service; jika payment gagal, transaksi gagal)
  - `/transactions` (GET) - Retrieve all transactions
  - `/transactions/{id}` (GET) - Retrieve transaction by ID
  - `/transactions/{id}` (PUT) - Update transaction
  - `/transactions/{id}` (DELETE) - Delete transaction
  - `/products` (POST) - Create new product
  - `/products` (GET) - Retrieve all products
  - `/products/{id}` (GET) - Retrieve product by ID
  - `/products/{id}` (PUT) - Update product
  - `/products/{id}` (DELETE) - Delete product

> **Catatan:**  
> Pada endpoint `/transactions (POST)` di Shopping Service, harus ada integrasi ke Payment Service (call endpoint `/payments` secara synchronous). Jika pembayaran gagal, transaksi juga gagal.

---

## Database Schema

Berikut adalah _usulan skema database_ (MongoDB Collection) berdasarkan kebutuhan endpoint dan constraint:

### **Products Collection**

| Field      | Type   | Description                        | Constraint          |
| ---------- | ------ | ---------------------------------- | ------------------- |
| id         | String | Unique identifier for product      | primary key, unique |
| name       | String | Name of the product                |                     |
| price      | Float  | Price of the product               |                     |
| stock      | Int    | Quantity in stock                  |                     |
| created_at | Date   | Timestamp when product was created |                     |

---

### **Transactions Collection**

| Field      | Type   | Description                            | Constraint          |
| ---------- | ------ | -------------------------------------- | ------------------- |
| id         | String | Unique identifier for transaction      | primary key, unique |
| product_id | String | Refers to product                      |                     |
| payment_id | String | Refers to payment                      |                     |
| quantity   | Int    | Number of product purchased            |                     |
| total      | Float  | Total price of the transaction         |                     |
| status     | String | Status: success/failed                 |                     |
| created_at | Date   | Timestamp when transaction was created |                     |

---

### **Payments Collection**

| Field      | Type   | Description                            | Constraint          |
| ---------- | ------ | -------------------------------------- | ------------------- |
| id         | String | Unique identifier for payment          | primary key, unique |
| email      | String | Payer email                            | unique index        |
| amount     | Float  | Payment amount                         |                     |
| status     | String | Payment status: success/failed/pending |                     |
| created_at | Date   | Timestamp when payment was created     |                     |

---

**Catatan:**

- **id** di semua collection harus unique (primary key).
- **email** di Payments harus unique (gunakan unique index pada field email).
- Field dapat dikembangkan/disesuaikan lagi sesuai kebutuhan aplikasi.

---

## Error Handling

- Gunakan function/utility error handler seperti pada Phase 2.
- **WAJIB** menangani semua error case, edge case, dan validasi input.

---

## Docker Kontainerisasi

- Siapkan **Dockerfile** di masing-masing service.
- Sertakan dokumentasi singkat cara build & run aplikasi dengan Docker.

---

## Cloud Deployment (GCP)

- Deploy aplikasi ke GCP (Google Cloud Run).
- **Aplikasi WAJIB dapat diakses publik**.
- Lampirkan dokumentasi deployment (langkah-langkah jelas dari build hingga running di GCP).

---

## Job Scheduling

- Implementasikan **job scheduling** pada proses yang membutuhkan, contoh:
  - Pembersihan data kadaluarsa
  - Update status otomatis
- Bisa menggunakan Cloud Scheduler (GCP) atau scheduler di dalam service.

---

## Unit Test

- Buat **unit test** untuk memastikan setiap fungsi/method utama berjalan dengan benar.
- Minimal test pada fitur core (produk, transaksi, payment).

---

## Expected Results

Aplikasi terdiri dari **2 service**:

- **Shopping Service** berjalan pada `http://localhost:8080` dengan endpoint:
  - `/transactions` (CRUD)
  - `/products` (CRUD)
- **Payment Service** berjalan pada `http://localhost:8081` dengan endpoint:
  - `/payments` (POST)

**Ketentuan tambahan:**

- Semua constraint id/email harus dipenuhi.
- Semua operasi CRUD & integrasi pembayaran harus berfungsi.
- Endpoint bisa diakses & didokumentasikan.
- Unit test tersedia & berjalan.
- Dokumentasi kode & deployment lengkap.

---

## RESTRICTION

- **id** = primary key dan harus unik.
- **email** (jika ada) wajib unik (gunakan unique index).
- **Tidak boleh hardcode sensitive value di kode (gunakan ENV/konfigurasi).**

---

## Assignment Submission

Push Assignment yang telah Anda buat ke akun Github Classroom Anda masing-masing.

---

## Assignment Rubrics

| Criteria                                           | Meet Expectations                     | Points |
| -------------------------------------------------- | ------------------------------------- | ------ |
| Problem Solving (API endpoint & flow)              | Implemented & working correctly       | 75 pts |
| Database Design                                    | MongoDB schema sesuai kebutuhan       | 10 pts |
| Database queries efficient & appropriately indexed | Query efisien, index sesuai kebutuhan | 5 pts  |
| Readability                                        | Code is well-documented & readable    | 5 pts  |
| Documentation/comments                             | Cukup jelas & lengkap                 | 5 pts  |

Total Points: **100**

---

## Notes:

- **Deadline:** W2D1 - 18.00 (telat = nilai 0)

---

## Final

- Url aplikasi [34.101.156.80:8000](34.101.156.80:8000)
- Swagger /swagger/ (GET)
