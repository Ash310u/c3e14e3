# ping-makra

Minimal Go backend with three routes:

- `GET /ping` returns `pong`
- `GET /schema` returns `authentication required`
- `POST /extract` returns `authentication required`

Run it with:

```bash
go run .
```

Run it with Docker Compose (one command):

```bash
docker compose up --build
```
