FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go .

# builds go executable
RUN CGO_ENABLED=0 GOOS=linux go build -o ./httpserver

RUN mkdir data
RUN touch data/count.txt

# set possible directory for a volume
VOLUME [ "/app/data/" ]

# sets default value for environment variable
ENV TIME_ROUTE="time"

# exposes port of container to the host
EXPOSE 10101

CMD [ "./httpserver" ]