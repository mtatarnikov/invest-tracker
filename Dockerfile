# Golang как базовый образ
FROM golang:latest

# Установить netcat-openbsd (нужен для wait-for-it.sh)
RUN apt-get update && apt-get install -y netcat-openbsd

# Установить Goose (migrations)
RUN go install github.com/pressly/goose/cmd/goose@latest

# Установить рабочую директорию
WORKDIR /app

# Копировать файлы проекта
COPY . .

# Копировать config.json
COPY config.json /app/config.json

# Установить переменную окружения для пути к конфигурационному файлу
ENV CONFIG_PATH /app/config.json

# Скомпилировать приложение
RUN go build -o main ./cmd/web/main.go

# Порт контейнера
EXPOSE 80

# Добавить скрипт ожидания для базы данных
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Запуск через entrypoint
RUN chmod +x entrypoint.sh
ENTRYPOINT ["/app/entrypoint.sh"]

# Выполнение миграций и запуск приложения
CMD ["./wait-for-it.sh", "db:5432", "--", "./main"]
