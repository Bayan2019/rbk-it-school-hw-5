# Домашнее задание 5

## Database Schema

[![rbk](./rbk-weather-api.svg "rbk DataBase")](https://dbdiagram.io/d/rbk-weather-api-69ef55e7c6a36f9c1b925ccc)

## REST API

### user

#### Register

```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"bayan@example.com","password":"tramp",
        "first_name": "Bayan", "last_name": "User",
        "is_active": true}'
```

#### Login

```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"bayan@example.com","password":"tramp"}'
```

```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"admin123"}'
```

#### Profile

```bash
curl http://localhost:8080/users/me \
  -H "Authorization: Bearer $TOKEN"
```

### cities

#### Добавить город пользователю

```bash
curl -X POST http://localhost:8080/cities \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "city": "Almaty"
  }'
```

#### Список городов пользователя

```bash
curl -X GET http://localhost:8080/cities \
-H "Authorization: Bearer $TOKEN"
```

#### Удалить город

```bash
curl -X DELETE http://localhost:8080/cities/3 \
-H "Authorization: Bearer $TOKEN"
```

### weather

#### Получить погоду по всем городам пользователя

```bash
curl -X GET http://localhost:8080/weather \
-H "Authorization: Bearer $TOKEN"
```

#### Получить историю с фильтрацией

```bash
curl -X GET http://localhost:8080/weather/history?city=astana \ 
-H "Authorization: Bearer $TOKEN"
```