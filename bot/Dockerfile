FROM golang:1.16-alpine
COPY . /overengineered-groupmebot
WORKDIR /overengineered-groupmebot
ENV GO111MODULE=auto
RUN go build -o bot/main bot/main.go 
CMD ["bot/main"]