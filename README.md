# Ozon Fintech: тестовое задание

### Задача: Сервис сокращения ссылок с возможностью выбора способа хранилищ(database, inmemory)
---
#### Запуск:
    1. git clone https://github.com/oonufg/ozon_fintech_test.git
    2. docker-compose up
#### Отключение:
    1. docker-compose down
---
#### API:
    Существует Swagger файл:
    Лежит: docs/api
### HTTP Методы
* ### Создать сокращённую ссылку:
      POST: /api/v1/urls
      {
        "fullUrl": {string value}
      }
      RESPONSE:
        CODE: 200
        {
          "compressedUrl": {string value}
        }
* ### Получить полную ссылку по сокращённой:
      GET: /api/v1/urls/{CompressedURL}
      RESPONSE:
        CODE: 200
        {
          "fullUrl": {string value}
        }
        CODE: 404
