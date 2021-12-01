package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/emiloh/DISYS_AuctionHouse/Proto"
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
	marked := -1
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
				marked = i
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
	if marked != -1 {
		clients = removeServer(clients, marked)
	}
	if succes >= len(clients) / 2 + 1{
		log.Printf("Bid was succesful. Passing it to %s.", offer.User)
		return &Proto.Ack{Response: Proto.Ack_SUCCES}, nil
	}else {
		log.Printf("Bid was unsuccesful. Passing it to %s.", offer.User)
		return &Proto.Ack{Response: Proto.Ack_FAIL}, nil
	}
	
}

func (fs *frontendServer) Result(ctx context.Context, info *Proto.Info) (*Proto.Details, error) {
	index := rand.Intn(len(clients))
	log.Printf("Sending %s's request to show current status on auction to RM%d", info.Uid, index)
	var err error
	var details *Proto.Details
	timeout := time.After(5 * time.Second)
	ticker := time.NewTicker(500 * time.Millisecond)
	marked := -1
	check:
	for{
		select{
			// Reached five seconds of trying with no luck
		case <- timeout:
			log.Printf("RM %d is not responding. Assuming it is dead. Removing it from list over RMs", index)
			err = errors.New("Server down")
			marked = index
			break check
		case <- ticker.C:
			if  details != nil {
				err = nil
				log.Printf("Received details from RM%d. Passing it to %s", index, info.Uid)
				break check
			}
			details, _ = clients[index].Result(ctx, info)
		}
	}
	if marked != -1 {
		clients = removeServer(clients, index)
	}
	if err != nil {
		details, err = clients[0].Result(ctx, info)	
		log.Printf("Received details from RM%d. Passing it to %s", index, info.Uid)
	}
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

func removeServer(p []Proto.AuctionHouseClient, index int) []Proto.AuctionHouseClient {
	return append(p[:index], p[index+1:]...)
}