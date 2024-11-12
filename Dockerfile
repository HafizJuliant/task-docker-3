FROM golang:1.20-alpine

# Setel working directory di dalam container
WORKDIR /myapp

# Salin file go.mod dan go.sum jika ada
# COPY go.mod go.sum ./

# Salin AUTHORS.md ke dalam image Docker
COPY AUTHORS.md .

# Salin semua file dari project ke container
COPY . .

# Build aplikasi
RUN go build -o my-go-app .

# Copy LINKS.md after building
COPY LINKS.md .

# Ekspos port 80 untuk aplikasi
EXPOSE 80

# Jalankan aplikasi saat container mulai
CMD ["./my-go-app"]
