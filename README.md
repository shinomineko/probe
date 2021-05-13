# probe

| URI     | Method | Description                      | Response          |
|---------|--------|----------------------------------|-------------------|
| /status | GET    | Response 200 OK                  | {"status":"up"}   |
| /status | GET    | Response 503 Service Unavailable | {"status":"down"} |
| /start  | GET    | Set /status to return 200        | {"status":"ok"}   |
| /stop   | GET    | Set /status to return 503        | {"status":"ok"}   |
