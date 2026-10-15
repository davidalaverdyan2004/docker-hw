# employees

Учебный проект: небольшое api для учёта сотрудников. На go,
данные хранятся в postgres запускается все через docker сompose

## умеет

- `GET /employees` — получить всех сотрудников
- `POST /employees` — добавить сотрудника
- `DELETE /employees/{id}` — удалить по id

Для POST отправляю такой JSON:

```json
{"full_name": "Давид Алавердян", "position": "разработчик"}
```

## Как запустить

сначала скопировать `.env.example` в `.env` и вписать свой пароль:

```bash
cp .env.example .env
```

потом:

```bash
make run-http-app
```

поднимутся два контейнера, приложение и база. API на `localhost:5050`

остановить:

```bash
make postgres-down
```

данные не пропадают после перезапуска, они лежат в docker volume

Если хочется запустить приложение без Docker (например, для отладки),
можно поднять только базу через `make postgres-up` и потом `make run`.

## Что использовал

goLang, net/http, pgx (pgxpool), PostgreSQL, Docker, Docker Compose, Make.