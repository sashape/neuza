# Используем базовый образ Golang
FROM golang:latest

# Устанавливаем переменную окружения для работы Go модулей
ENV GO111MODULE=on

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum и загружаем зависимости
COPY go.mod .
COPY go.sum .
COPY air.toml .
RUN go mod download

# Устанавливаем Air для автоматической перезагрузки приложения
RUN go install github.com/cosmtrek/air@latest

# Копируем все файлы проекта в рабочую директорию
COPY . .

# Команда по умолчанию для запуска приложения с помощью Air
CMD ["air", "-c", "air.toml"]