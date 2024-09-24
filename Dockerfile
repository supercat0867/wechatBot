FROM golang:1.20-alpine

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

RUN go build -o /app/main .

WORKDIR /app

EXPOSE 8080

CMD ["./main"]
