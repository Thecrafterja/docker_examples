FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o /httpserver

ENV TIME_ROUTE="time"

EXPOSE 10101

CMD [ "/httpserver" ]