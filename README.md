# ping-makra

Minimal Go backend with three routes:

- `GET /ping` returns `pong`
- `GET /schema` returns `authentication required`
- `POST /extract` returns `authentication required`

Run it with:

```bash
go run .
```

Run it with Docker:

```bash
docker build -t ping-server .
docker run -p 8080:8080 ping-server
```
