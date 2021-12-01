package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/emiloh/DISYS_AuctionHouse/tree/simpler/Proto"
	"google.golang.org/grpc"
)

var lock sync.Mutex
var id int64 = 0
var name string = "Tamagotchi"
var price int64 = 500
var user string = ""
var timeleft int64 = 120
var sold = false

type RMServer struct {
	Proto.UnimplementedAuctionHouseServer
}

func main() {
	port := os.Args[1]

	server := &RMServer{Proto.UnimplementedAuctionHouseServer{}}

	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen on port, %s: %v", port, err)
	}

	Proto.RegisterAuctionHouseServer(grpcServer, server)

	log.Println("DISYS Auction House is online.")
	go startCountDown()
	
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve incoming connecitons on listener: %v", err)
	}
}

func (rms *RMServer) Bid(ctx context.Context, offer *Proto.Offer) (*Proto.Ack, error) {
	log.Printf("Recieved bid from %s. Checking bid against current bid.", offer.User)
	if !sold {
		lock.Lock()
		defer lock.Unlock()

		if offer.Amount > price {
			log.Printf("%s's bid was higher than current bid. Accepting bid. Current offer is now %d", offer.User, offer.Amount)
			price = offer.Amount
			user = offer.User
			return &Proto.Ack{Response: Proto.Ack_SUCCES}, nil
		}else if offer.Amount <= price {
			log.Printf("%s's bid was too low. Rejecting bid.", offer.User)
			return &Proto.Ack{Response: Proto.Ack_FAIL}, nil
		}
		return &Proto.Ack{Response: Proto.Ack_EXCEPTION}, errors.New("wtf")
	}else{
		log.Println("The bidding session is over and the item has been sold.")
		return &Proto.Ack{Response: Proto.Ack_SOLD}, nil
	}
}

func (rms *RMServer) Result(ctx context.Context, info *Proto.Info) (*Proto.Details, error) {
	details := &Proto.Details{
		Id: id,
		Name: name,
		Timeleft: timeleft,
		Amount: price,
		User: user,
	}
	return details, nil
}

func startCountDown() {
	log.Printf("The bidding session for %s has begun!", name)
	tickerS := time.NewTicker(1 * time.Second)
	ticker10S := time.NewTicker(10 * time.Second)
	blocker := make(chan bool)

	go func (){ 
		for{
			select{
			case <- blocker:
				return
			case <- ticker10S.C:
				log.Printf("%d seconds remaining of the auction.", timeleft)
			case <- tickerS.C:
				lock.Lock()
				timeleft--
				lock.Unlock()
			}
		}
	}()

	time.Sleep(120 * time.Second)
	ticker10S.Stop()
	tickerS.Stop()
	sold = true
	blocker <- true
	log.Println("The auction has offically ended")
}
