# Go RESTful API Платежной системы

<div align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
  <img src="https://img.shields.io/badge/API-FF6C37?style=for-the-badge&logo=postman&logoColor=white" alt="API">
</div>

## Описание проекта

Проект представляет собой микросервис для обработки транзакций с REST API, разработанный на Go. Он поддерживает лучшие практики, соответствующие принципам SOLID и DDD архитектуре. Основные функции системы:

- **Перевод средств** между кошельками через POST /api/send
- **Просмотр истории** последних транзакций через GET /api/transactions?count=N
- **Проверка баланса** кошелька через GET /api/wallet/{address}/balance

При первом запуске автоматически создаются 10 тестовых кошельков с начальным балансом 100.0 у.е. каждый.

**Ключевые требования:**
- Реализация на Go с использованием реляционной БД (PostgreSQL)
- Обеспечение безопасности транзакций
- Сохранение данных между перезапусками сервера
- Четкая структура кода и документация

Проект разработан в рамках тестового задания для стажера Go-разработчика [infotecs](https://infotecs.ru) и предоставляет следующий функционал:

- RESTful endpoints в общепринятом формате
- Стандартные CRUD операции с базой данной
- Управление конфигурацией приложени, зависящее от среды
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

## Запуск

### Установка проекта
```bash
# установка проекта
git clone github.com/normalniydada/case_infotecs
cd case_infotecs

# запуск контейнера с PostgreSQL
docker-compose up db -d

# запуск проекта
go run cmd/main.go
```














