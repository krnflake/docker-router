docker-router
=======

Simple HTTP proxy that supports WebSocket. It's designed to route traffic to docker containers by parsing http headers.

### Architecture

```
    _______       _______________       _________________
    |Nginx| ----> |docker-router| ----> |docker upstream|
    -------       ---------------       -----------------
```

Example nginx settings:

```
server {
  listen 80;

  server_name example.org;

  location / {
    proxy_pass http://127.0.0.1:8080;
    proxy_set_header DockerID c788c689dd72;
    proxy_set_header Port 80;
  }
}
```

### License

>Copyright (C) krnflake - All Rights Reserved  
>Unauthorized copying of this file, via any medium is strictly prohibited  
>Proprietary and confidential  
>Written by krnflake <hello@krnflake.ru>, August 2014
