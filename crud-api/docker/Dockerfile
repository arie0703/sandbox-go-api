FROM golang:alpine3.15

WORKDIR /api

COPY . .

RUN go mod tidy

CMD ["go", "run", "./main.go"]
