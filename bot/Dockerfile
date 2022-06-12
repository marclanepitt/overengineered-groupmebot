FROM golang:alpine
COPY . /overengineered-groupmebot
WORKDIR /overengineered-groupmebot
ENV GO111MODULE=on
RUN go build -o bot/main bot/main.go 
CMD ["bot/main"]