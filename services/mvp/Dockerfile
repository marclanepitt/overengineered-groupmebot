FROM golang:1.16-alpine
COPY . /overengineered-groupmebot
WORKDIR /overengineered-groupmebot
ENV GO111MODULE=on
ENV PATH=$PATH:$GOPATH/bin
ENV GOROOT=/usr/local/go
ENV PATH=$PATH:$GOROOT/bin
RUN go build -o services/mvp/server/server services/mvp/server/server.go 
CMD ["services/mvp/server/server"]