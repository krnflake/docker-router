docker-router
=======
Simple HTTP proxy that supports WebSocket. It's designed to route traffic to docker containers by parsing http headers.

### Architecture

```
    __________       _______________       _________________
    |weberver| ----> |docker-router| ----> |docker upstream|
    ----------       ---------------       -----------------
```

Example nginx settings:
```
server {
  listen 80;

  server_name example.org;

  location / {
    proxy_pass http://127.0.0.1:8080;
    proxy_set_header docker c788c689dd72:80;
  }
}
```

