FROM golang:1.23-alpine AS builder
WORKDIR /app

# Copia os arquivos atuais
COPY . .

# Atualiza e limpa as dependências com o Go correto
RUN go mod tidy

# Compila o executável
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
