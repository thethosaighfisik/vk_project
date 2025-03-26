# Используем официальный Go образ как базовый
FROM golang:1.22

# Устанавливаем необходимые зависимости
RUN apt-get update && apt-get install -y \
    pkg-config \
    libssl-dev \
    libssl3

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта в рабочую директорию контейнера
COPY . .

# Загружаем зависимости Go
RUN go mod download

# Компилируем приложение
RUN go build -o /app/main .

# Открываем порты
EXPOSE 8080

# Запускаем приложение
CMD ["/app/main"]

