FROM alpine

WORKDIR /app
COPY redis_tls_connect_check /app

ENTRYPOINT ["./redis_tls_connect_check"]
