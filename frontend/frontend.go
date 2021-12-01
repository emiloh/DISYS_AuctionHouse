package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

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
	log.Printf("Sending %s's bid of %d to RMs", offer.User, offer.Amount)
	succes := 0
	fails := 0
	for i := 0; i < len(clients); i++ {
		var ack *Proto.Ack
		timeout := time.After(5 * time.Second)
		ticker := time.NewTicker(500 * time.Millisecond)
		check:
		for{
			select{
				// Reached five seconds of trying with no luck
			case <- timeout:
				log.Printf("RM %d is not responding. Assuming it is dead. Removing it from list over RMs", i)
				//implement removal
				break check
			case <- ticker.C:
				if ack != nil {
					if ack.Response == Proto.Ack_SUCCES {
						succes++
					}else{
						fails++
					}
					break check
				}
				ack, _ = clients[i].Bid(ctx, offer)
			}
		}
	}
	if succes >= len(clients) / 2 + 1{
		return &Proto.Ack{Response: Proto.Ack_SUCCES}, nil
	}else {
		return &Proto.Ack{Response: Proto.Ack_FAIL}, nil
	}
	
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