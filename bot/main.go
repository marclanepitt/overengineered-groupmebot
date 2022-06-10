package main

import (
	"context"
	"log"
	"net/http"
	"overengineered-groupmebot/protos"
	"overengineered-groupmebot/services/mvp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	http.HandleFunc("/botRequest", routeRequest)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func routeRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := mvp.NewMvpClient(conn)
	message, _ := client.GetMvpMessage(context.TODO(), &protos.Player{
		Name: "test",
	})

	log.Println(message.GetBody())
}
