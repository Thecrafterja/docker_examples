FROM golang:1.23

WORKDIR /app

COPY . .

# At buildtime
RUN CGO_ENABLED=0 GOOS=linux go build -o ./compilexample

# At runtime
CMD [ "./compilexample" ]