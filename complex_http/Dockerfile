FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go .

# at buildtime
RUN CGO_ENABLED=0 GOOS=linux go build -o /httpserver

# add here
# create count.txt
# volume
# environment variable
# ports

# at runtime
CMD [ "/httpserver" ]