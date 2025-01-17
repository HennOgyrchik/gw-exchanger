FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o gw-exchanger ./cmd/

FROM scratch
WORKDIR /app
COPY --from=builder ./app/gw-exchanger .
COPY --from=builder ./app/config.env .
COPY --from=builder ./app/internal/storages/migrations ./migrations
EXPOSE 9090
CMD ["./gw-exchanger"]

#docker build -t gw-exchanger .
#docker run --name gw-exchanger -d gw-exchanger