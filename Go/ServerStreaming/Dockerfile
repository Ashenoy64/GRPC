FROM golang:1.22.6-bullseye

WORKDIR /app

COPY . .

RUN go mod download

RUN apt update
RUN apt install iproute2 -y

RUN go build -o main .

EXPOSE 50051

CMD ["./main"]