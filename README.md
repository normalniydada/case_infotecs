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

Проект разработан в рамках тестового задания для стажера Go-разработчика infotecs.
