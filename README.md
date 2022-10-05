# probe

| URI     | Method | Description                      | Response           |
|---------|--------|----------------------------------|--------------------|
| /status | GET    | Response 200 OK                  | {"message":"up"}   |
| /status | GET    | Response 503 Service Unavailable | {"message":"down"} |
| /start  | POST   | Set /status to return 200        | {"message":"ok"}   |
| /stop   | POST   | Set /status to return 503        | {"message":"ok"}   |
