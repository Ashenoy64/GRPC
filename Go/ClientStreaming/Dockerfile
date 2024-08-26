FROM golang:1.23.0

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN apt update
RUN apt install iproute2 -y

RUN go build -o main Server.go

EXPOSE 50051

CMD ["./main"]
