# Go kurulum imajını başlangıç noktası olarak kullan
FROM golang:alpine as builder

# Çalışma dizinini ayarla
WORKDIR /app

# Go modül dosyalarını kopyala
COPY go.mod .
COPY go.sum .

# Bağımlılıkları indir
RUN go mod download

# Uygulama dosyalarını kopyala
COPY . .

# Swagger belgelerini oluştur
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

# Uygulamayı derle
RUN go build -o main .

# Çalışma aşaması
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Derlenmiş uygulamayı önceki aşamadan kopyala
COPY --from=builder /app/main .

# Portu aç
EXPOSE 8080

# Uygulamayı çalıştır
CMD ["./main"]
