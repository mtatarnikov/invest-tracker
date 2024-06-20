#!/bin/bash

# Ожидание доступности базы данных
/wait-for-it.sh db:5432 --timeout=30 -- echo "Database is up"

# Выполнение миграций
goose -dir migrations postgres "host=$DB_HOST user=$DB_USER password=$DB_PASS dbname=$DB_NAME sslmode=$DB_SSLMODE" up

# Запуск приложения
./main
