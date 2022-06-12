FROM golang:1.16-alpine
COPY . /overengineered-groupmebot
WORKDIR /overengineered-groupmebot
RUN go build -o bot/main bot/main.go 
CMD ["bot/main"]