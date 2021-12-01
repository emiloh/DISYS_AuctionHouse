package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/emiloh/DISYS_AuctionHouse/tree/simpler/Proto"
	"google.golang.org/grpc"
)

type frontendServer struct {
	Proto.UnimplementedAuctionHouseServer
}

var clients []Proto.AuctionHouseClient

func main() {
	listenPort := os.Args[1]
	ports := []string{os.Args[2], os.Args[3], os.Args[4]}

	go connectToRMs(ports)

	server := &frontendServer{Proto.UnimplementedAuctionHouseServer{}}

	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", listenPort)

	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	Proto.RegisterAuctionHouseServer(grpcServer,server)

	grpcServer.Serve(listener)
}

func (fs *frontendServer) Bid(ctx context.Context, offer *Proto.Offer) (*Proto.Ack, error) {
	//failed := []Proto.AuctionHouseClient
	for i := 0; i < len(clients) {
		
	}
	return &Proto.Ack{Response: Proto.Ack_SUCCES}, nil
}

func (fs *frontendServer) Result(ctx context.Context, info *Proto.Info) (*Proto.Details, error) {
	details, err := clients[0].Result(ctx, info)
	return details, err
}

func connectToRMs(ports []string){
	blocker := make(chan int)
	for i :=0; i < len(ports); i++ {
		conn, err := grpc.Dial(ports[i], grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect: %v", err)
		}
		defer conn.Close()
		clients = append(clients, Proto.NewAuctionHouseClient(conn))
	}
	<- blocker
}