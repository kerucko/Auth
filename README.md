# Auth

Сервис для аутентификации и авторизации пользователей.

Взаимодействие происходит с помощью gRPC запросов.

# gRPC методы

Этот сервис предоставляет два gRPC метода для регистрации пользователей и входа в систему.

## Register
Метод `Register` позволяет пользователям зарегистрироваться в системе. Он принимает запрос `RegisterRequest` и возвращает ответ `RegisterResponse`.

### RegisterRequest
- `email` (string): Email пользователя.
- `password` (string): Пароль пользователя.

### RegisterResponse
- `user_id` (int64): Уникальный идентификатор пользователя.

## Login
Метод `Login` используется для входа в систему. Он принимает запрос `LoginRequest` и возвращает ответ `LoginResponse`.

### LoginRequest
- `email` (string): Email пользователя.
- `password` (string): Пароль пользователя.
- `app_id` (int32): Идентификатор приложения.

### LoginResponse
- `token` (string): Токен аутентификации пользователя.
