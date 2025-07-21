# Go RESTful API Платежной системы

<div align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
  <img src="https://img.shields.io/badge/API-FF6C37?style=for-the-badge&logo=postman&logoColor=white" alt="API">
</div>

## Описание проекта

Проект представляет собой микросервис для обработки транзакций с REST API, разработанный на Go. Он поддерживает лучшие практики, соответствующие принципам SOLID и DDD архитектуре. 
При первом запуске автоматически создаются 10 тестовых кошельков с начальным балансом 100.0 у.е. каждый.

**Ключевые требования:**
- Реализация на Go с использованием реляционной БД (PostgreSQL)
- Обеспечение безопасности транзакций
- Сохранение данных между перезапусками сервера
- Четкая структура кода и документация

Проект разработан в рамках тестового задания [infotecs](https://infotecs.ru) для стажера Go-разработчика и предоставляет следующий функционал:

- RESTful endpoints в общепринятом формате
- Стандартные CRUD операции с базой данной
- Миграция базы данных (gorm.Automigrate)
- Валидация данных
- Обработка ошибок с корректной генерацией ответ на ошибки

## Используемые технологии и пакеты

### Основные зависимости 
- [Echo](https://echo.labstack.com/) - HTTP фреймоврк 
- [GORM](https://gorm.io/) - ORM для работы с PostgreSQL
- [UUID](https://github.com/google/uuid) - генерация уникальных адресов
- [Decimal](https://github.com/shopspring/decimal) - точные денежные вычисления

### Вспомогательные пакеты
- [Viper](https://github.com/spf13/viper) - управление конфигурацией
- [Godotenv](https://github.com/joho/godotenv) - загрузка переменных окружения из .env файла
- [GoDoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc) - генерация документации

## Установка и запуск

### Требования

* Go 1.21+
* Docker 20.10+
* PostgreSQL 15+


### Локальная разработка
```bash
# установка проекта
git clone github.com/normalniydada/case_infotecs
cd case_infotecs

# запуск контейнера с PostgreSQL
docker-compose up db -d

# запуск проекта
go run cmd/main.go
```

### Cборка Docker-контейнера

Необходимо в файле `config/config.yaml` внести следующие изменения:
```
database:
  host: "localhost" -> заменить на "db"
```

Далее в терминале ввести команду:
```bash
# сборка контейнера с приложением
docker-compose up --build
```

## REST API 

Сервер RESTful API работает по адресу `http://127.0.0.1:8080`. Он предоставляют следующие endpoints:

### **`POST /api/send`**: перевод средств между кошельками
    
Пример запроса (json):  
```
{
    "from" : "e240d825d255af751f5f55af8d9671beabdf2236c0a3b4e2639b3e182d994c88", # <- кошелек отправителя
    "to" : "e240d825d255af751f5f55af8d9671beabdf2236c0a3b4e2639b3e182d994c89", # <- кошелек получателя
    "amount" : 3.50 # <- сумма перевода
}
```
Коды ответов: 
* `200 OK` - успешный перевод
* `400 Bad Request` - неверный формат запроса
* `404 Not Found` - кошелек отправителя/получателя не найден
* `500 Internal Server Error` - серверная ошибка  

### **`GET /api/transactions?count=N`**: просмотр истории последних N транзакций  
    
Параметры: 
* `count` - количество возвращаемых транзакций (N)
  
Коды ответов: 
* `200 OK` - успешный запрос
* `400 Bad Requset` - неверный параметр
* `500 Internal Server Error` - серверная ошибка  

### **`GET /api/wallet/{address}/balance`**: проверка баланса кошелька  

Параметры пути:
* `address` - идентификатор (адрес) кошелька
  
Коды ответов:
* `200 OK` - успешный запрос
* `400 Not Found` - кошелек не найден
* `500 Internal Server Error` - серверная ошибка

















