package main

import (
	"context"
	"enigmacamp.com/omzetcln/api"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	var conn *grpc.ClientConn
	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_PORT")
	serverInfo := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(serverInfo, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewOmzetClient(conn)
	//omzet, err := c.SubmitOmzet(context.Background(), &api.OmzetRequestMessage{
	//	Period: "Januari 2021",
	//	Outlet: "0209",
	//	Omzet:  10000,
	//})
	res, err := c.ClearOmzet(context.Background(), &api.OmzetRequestMessage{Outlet: "0209"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)

}
