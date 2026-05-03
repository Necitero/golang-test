# Golang Test / Backend

This is the backend part of this repository.

## 🧪 Testing

1. Start the go service

```bash
cd golang-test/backend
go run .
```

2. Send HTTP requests (example: curl)

```bash
# POST
curl -H "Content-Type: application/json" -d '{"headline": "Clean room", "note": "Urgent","due_date":"2026-05-03T14:00:00Z","status": "open"}' 0.0.0.0:8080/api/todo

# GET
curl 0.0.0.0:8080/api/todo/[id]
```
