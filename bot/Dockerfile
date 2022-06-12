FROM golang:alpine
COPY . /overengineered-groupmebot
WORKDIR /overengineered-groupmebot
RUN go build -o bot/main bot/main.go 
CMD ["bot/main"]