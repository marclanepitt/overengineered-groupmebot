FROM golang:1.16-alpine
COPY . /overengineered-groupmebot
WORKDIR /overengineered-groupmebot
RUN go build -o services/mvp/server/server services/mvp/server/server.go 
CMD ["services/mvp/server/server"]