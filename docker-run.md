Tentu, berikut adalah cara untuk menjalankan image Docker Anda menggunakan `docker-compose` beserta konfigurasi untuk MongoDB sebagai replica set.

### Struktur Proyek

Pastikan Anda memiliki struktur file berikut:

```
.
├── docker-compose.yml
└── mongo-init.js
```

### 1. File `mongo-init.js`

Buat file ini untuk menginisialisasi MongoDB replica set.

```javascript
// mongo-init.js
rs.initiate({
  _id: "rs0",
  members: [{ _id: 0, host: "mongo:27017" }],
});
```

### 2. File `docker-compose.yml`

File ini akan mendefinisikan dan menghubungkan semua service Anda.

```yaml
version: "3.8"

services:
  # MongoDB Service with Replica Set
  mongo:
    image: mongo:4.4
    container_name: mongo
    ports:
      - "27017:27017"
    volumes:
      - ./mongo-init.js:/docker-entrypoint-initdb.d/mongo-init.js:ro
      - mongo-data:/data/db
    command: ["--replSet", "rs0", "--bind_ip_all"]
    networks:
      - app-net

  # Gateway Service
  gateway-service:
    image: embapge/hacktiv8-p3-w1-gc-embapge:gateway-service-1.0
    container_name: gateway-service
    ports:
      - "80:80" # Asumsi gateway berjalan di port 80
    depends_on:
      - shopping-service
      - payment-service
    networks:
      - app-net

  # Payment Service
  payment-service:
    image: embapge/hacktiv8-p3-w1-gc-embapge:payment-service-1.0
    container_name: payment-service
    environment:
      # Sesuaikan nama database jika berbeda
      - MONGO_URI=mongodb://mongo:27017/paymentdb?replicaSet=rs0
    depends_on:
      - mongo
    networks:
      - app-net

  # Shopping Service
  shopping-service:
    image: embapge/hacktiv8-p3-w1-gc-embapge:shopping-service-1.0
    container_name: shopping-service
    environment:
      # Sesuaikan nama database jika berbeda
      - MONGO_URI=mongodb://mongo:27017/shoppingdb?replicaSet=rs0
    depends_on:
      - mongo
    networks:
      - app-net

volumes:
  mongo-data:

networks:
  app-net:
    driver: bridge
```

### Cara Menjalankan

1.  Pastikan Anda berada di direktori yang sama dengan file `docker-compose.yml` dan `mongo-init.js`.
2.  Buka terminal atau command prompt Anda.
3.  Jalankan perintah berikut untuk membangun dan menjalankan semua container di background:

        ```bash
        docker-compose up -d
        ```

4.  Untuk menghentikan dan menghapus semua container, network, dan volume yang terkait, jalankan:

        ```bash
        docker-compose down -v
        ```
