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

### License
>Copyright (C) krnflake - All Rights Reserved  
>Unauthorized copying of this file, via any medium is strictly prohibited  
>Proprietary and confidential  
>Written by krnflake <hello@krn.ovh>, October 2014
