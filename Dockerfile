FROM golang:1.22.3 AS builder

WORKDIR /app

ADD . .

RUN CGO_ENABLED=0 go build -o app

FROM scratch

COPY --from=builder /app/app /app

CMD ["/app"]