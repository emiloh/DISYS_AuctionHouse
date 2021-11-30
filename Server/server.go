package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/emiloh/DISYS_AuctionHouse/tree/Auction/Proto"
	"google.golang.org/grpc"
)

type Item struct {
	id       int64
	name     string
	price    int64
	user     string
	timeleft int64
}
type Connection struct {
	id     string
	stream Proto.AuctionHouse_RegisterServer
	error  chan error
}

type RMServer struct {
	Proto.UnimplementedAuctionHouseServer
	connections map[string]*Connection
	items       map[int64]*Item
}

func main() {
	port := os.Args[1]
	connections := make(map[string]*Connection)
	items := make(map[int64]*Item)
	items = seed(items)

	server := &RMServer{Proto.UnimplementedAuctionHouseServer{}, connections, items}

	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	Proto.RegisterAuctionHouseServer(grpcServer, server)

	grpcServer.Serve(listener)
}

func (rms *RMServer) Bid(ctx context.Context, offer *Proto.Offer) (*Proto.Acknowledgement, error) {
	return &Proto.Acknowledgement{}, nil
}

func (rms *RMServer) Result(ctx context.Context, info *Proto.Info) (*Proto.Details, error) {
	return &Proto.Details{}, nil
}

func (rms *RMServer) View(ctx context.Context, user *Proto.User) (*Proto.DetailsList, error) {
	ids := []int64{rms.items[1].id, rms.items[2].id, rms.items[3].id}
	names := []string{rms.items[1].name, rms.items[2].name, rms.items[3].name}
	timelefts := []int64{rms.items[1].timeleft, rms.items[2].timeleft, rms.items[3].timeleft}
	prices := []int64{rms.items[1].price, rms.items[2].price, rms.items[3].price}
	users := []string{rms.items[1].user, rms.items[2].user, rms.items[3].user}
	details := &Proto.DetailsList{
		Id:       ids,
		Name:     names,
		Timeleft: timelefts,
		Amount:   prices,
		User:     users,
	}
	response := &Proto.Response{DetailList: details}
	rms.connections["0"].stream.Send(response)
	return details, nil
}

func (rms *RMServer) Register(request *Proto.RegisterRequest, stream Proto.AuctionHouse_RegisterServer) error {
	conn := &Connection{
		id:     request.Id,
		stream: stream,
		error:  make(chan error),
	}
	rms.connections[conn.id] = conn
	return <-conn.error
}

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
