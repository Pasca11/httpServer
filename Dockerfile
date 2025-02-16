FROM golang:1.23-alpine

WORKDIR /app

# Копируем файлы с зависимостями
COPY go.mod ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код
COPY *.go ./

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Используем минимальный образ для запуска
FROM alpine:latest

WORKDIR /app

# Копируем собранное приложение из предыдущего этапа
COPY --from=0 /app/main .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"] 