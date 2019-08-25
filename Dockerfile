FROM debian:latest
RUN mkdir -p /app
WORKDIR /app

ADD transavro-srv /app/transavro-srv

CMD ["./transavro-srv"]
