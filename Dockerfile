FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .
RUN go mod download
RUN go mod tidy
RUN go build cmd/main.go


FROM keymetrics/pm2:18-alpine

WORKDIR /app

COPY --from=builder  /app/main .
COPY templates ./templates
COPY views ./views
COPY ecosystem.config.js .
COPY app.env .
RUN chmod ++x /app/main

CMD ["pm2-runtime", "main"]