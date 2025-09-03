#STAGE 1
FROM golang:1.24.6-alpine AS builder

RUN  apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# BUILD MAIN BINARY
RUN go build -o auth-service main.go

#STAGE 2

FROM alpine:latest

RUN apk add --no-cache ca-certificates && \
    apk add --no-cache wget

WORKDIR /root/

COPY --from=builder /app/auth-service .

EXPOSE 4000

CMD [ "./auth-service" ]
