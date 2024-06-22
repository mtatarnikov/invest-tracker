# Инструкция по установке проекта Invest-Tracker

## Предварительные требования
Убедитесь, что у вас уже установлены **Docker** и **Git**.

## Установка

1. Клонируйте репозиторий локально:
   ```
   git clone https://github.com/mtatarnikov/invest-tracker.git
   ```

2. Добавьте файл `config.json` в корень проекта со следующим содержанием:
   ```json
   {
       "database": {
           "host": "db",
           "port": 5432,
           "user": "postgres",
           "password": "p0stgreS",
           "dbname": "invest"
       },
       "htmlTemplatePath": "/app/ui/html/",
       "htmlUiStaticPath": "/app/"
   }
   ```

3. Перейдите в корневую папку проекта, где находятся `config.json`, `Dockerfile`, `docker-compose.yml` и другие необходимые файлы.

4. Соберите и запустите контейнеры с помощью Docker Compose:
   ```
   docker-compose up --build
   ```