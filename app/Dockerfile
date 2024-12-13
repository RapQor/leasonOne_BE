# Gunakan image resmi Go
FROM golang:alpine

# Set working directory di dalam container
WORKDIR /app

# Salin semua file dari folder lokal `app/` ke dalam container
COPY . .

# Jalankan `go mod tidy` untuk memastikan dependensi sudah diunduh
RUN go mod tidy

# Ekspos port 8080 (sesuai konfigurasi server)
EXPOSE 8080

# Perintah untuk menjalankan aplikasi
CMD ["go", "run", "main.go"]
