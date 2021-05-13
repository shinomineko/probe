# probe

| URI    | Method | Description                      | Response          |
|--------|--------|----------------------------------|-------------------|
| /hello | GET    | Response 200 OK                  | {"status":"up"}   |
| /hello | GET    | Response 503 Service Unavailable | {"status":"down"} |
| /start | GET    | Set /hello to return 200         | {"status":"ok"}   |
| /stop  | GET    | Set /hello to return 503         | {"status":"ok"}   |
