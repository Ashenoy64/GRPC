FROM golang:1.23.0-bullseye

WORKDIR /app

COPY . .

RUN go mod download

RUN apt update
RUN apt install iproute2 -y

RUN go build -o main .

EXPOSE 50051

CMD ["./main"]