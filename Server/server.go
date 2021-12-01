package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/emiloh/DISYS_AuctionHouse/tree/simpler/Proto"
	"google.golang.org/grpc"
)

/*type Item struct {
	id       int64
	name     string
	price    int64
	user     string
	timeleft int64
}*/

var id int64 = 0
var name string = "Tamagotchi"
var price int64 = 500
var user string = ""
var timeleft int64 = 7000

type RMServer struct {
	Proto.UnimplementedAuctionHouseServer
	//items       map[int64]*Item
}

func main() {
	port := os.Args[1]
	//items := make(map[int64]*Item)
	//items = seed(items)

	server := &RMServer{Proto.UnimplementedAuctionHouseServer{}}

	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	Proto.RegisterAuctionHouseServer(grpcServer, server)

	grpcServer.Serve(listener)
}

func (rms *RMServer) Bid(ctx context.Context, offer *Proto.Offer) (*Proto.Ack, error) {
	return &Proto.Ack{}, nil
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



/*
func seed(items map[int64]*Item) map[int64]*Item {
	Item1 := &Item{
		id:       1,
		name:     "Damaged shoes",
		price:    10,
		user:     "",
		timeleft: 700,
	}
	Item2 := &Item{
		id:       2,
		name:     "Tamagotchi",
		price:    100,
		user:     "",
		timeleft: 7000,
	}
	Item3 := &Item{
		id:       3,
		name:     "Fast bike",
		price:    500,
		user:     "",
		timeleft: 10000,
	}

	items[1] = Item1
	items[2] = Item2
	items[3] = Item3

	return items
}

*/
