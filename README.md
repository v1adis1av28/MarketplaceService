# MarketplaceService

Rest-api сервис с имплементацией авторизации/регистрацией и взаимодействия с объявлениями.

## Запуск проекта

### Клонирование и сборка
Склонируйте репозиторий. перейдите в рабочую директорию и запустите docker compose файл.

```bash
git clone https://github.com/v1adis1av28/MarketplaceService.git
cd MarketplaceService
docker-compose up --build
````
Сервис будет доступен по адресу: [http://localhost:8080](http://localhost:8080)

Документация Swagger по адресу: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)


## Эндпоинты

### Регистрация

`POST /auth/register`

```json
{
  "email": "newuser@example.com",
  "password": "pass1234"
}
```

**Ответ:**

```json
{
  "message": "you successfully registered!"
}
```

---

### Авторизация

`POST /auth/login`

```json
{
  "email": "user1@mail.com",
  "password": "1234"
}
```

**Ответ:**

```json
{
  "message": "you successfully authorized!",
  "token": "Bearer <JWT>"
}
```

---

### Создание объявления

`POST /api/advertisement`

**Защищенный эндпоинт, чтобы им воспользоваться необходимо авторизоваться**

```
Authorization: Bearer <JWT>
```

```json
{
  "header" : "test223",
  "description" : "descrip2t12ion3",
  "image" : "http://.pro/wbd-front/skillbox-static/skillbox3.jpg",
  "price" : 100
}
```

**Ответ:**

```json
{
  "message": "advertisement created successfully"
}
```

---

### Лента объявлений

`GET /advertisements?limit=10&offset=0&sort=price_asc`

* `limit` — количество на странице (по умолчанию: 10)
* `offset` — смещение (по умолчанию: 0)
* `sort` — `price_asc`, `price_desc`, `created_at_desc`

**Пример запроса:**

```
GET /advertisements?limit=5&offset=0&sort=created_at_desc
```

**Ответ:**

```json
[
  {
        "id": 1,
        "header": "Lada 2112",
        "description": "В идеальном состоянии.",
        "image": "https://example.com/image1.jpg",
        "price": 80000,
        "created_at": "2025-07-21T16:31:10.539409Z",
        "owner_email": "test1@example.com",
        "is_owner": false
    },
    {
        "id": 6,
        "header": "Ноутбук ASUS",
        "description": "Для учебы и работы.",
        "image": "https://example.com/image5.jpg",
        "price": 45000,
        "created_at": "2025-07-21T16:31:10.539409Z",
        "owner_email": "admin@example.com",
        "is_owner": false
        },
  ...
]
```


