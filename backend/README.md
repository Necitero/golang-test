# Golang Test / Backend

This is the backend part of this repository.

## 🧪 Testing

1. Start the docker container service

```bash
docker-compose build backend && docker-compose up
```

2. Send HTTP requests (example: curl)

```bash
# POST
curl -H "Content-Type: application/json" -d '{"headline": "Clean room", "note": "Urgent","due_date":"2026-05-03T14:00:00Z","status": "open"}' 0.0.0.0:8080/api/todo

# GET
curl 0.0.0.0:8080/api/todo/[id]
```

3. To restart the container after changes, run

```bash
docker-compose down && docker-compose build backend && docker-compose up
```

> [!WARNING]
> While it is possible to run the go backend on its own, it is not recommended as the sql instance requires to run in a docker container.
>
> Alternatively, run an sql container independently and use the following to only run the backend:

```bash
    cd golang-test/backend
    go run .
```
