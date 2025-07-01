FROM golang:buster as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

##########################

FROM debian:buster as prod

WORKDIR /app
COPY --from=builder /app/app /app/app
COPY --from=builder /app/static /app/static
EXPOSE 8080
CMD ["./app"]