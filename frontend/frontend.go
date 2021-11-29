package main

import (
	"context"
	"log"
	"net"

	"github.com/emiloh/DISYS_AuctionHouse/tree/Auction/Proto"
	"google.golang.org/grpc"
)

type Connection struct {
	id     string
	stream Proto.AuctionHouse_RegisterServer
	error  chan error
}
type frontendServer struct {
	Proto.UnimplementedAuctionHouseServer
	connetions map[string]*Connection
}

func main() {
	connections := make(map[string]*Connection)

	server := &frontendServer{Proto.UnimplementedAuctionHouseServer{}, connections}

	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	Proto.RegisterAuctionHouseServer(grpcServer,server)

	grpcServer.Serve(listener)
}

func (fs *frontendServer) Bid(ctx context.Context, offer *Proto.Offer) (*Proto.Acknowledgement, error) {

	return &Proto.Acknowledgement{Status: Proto.Acknowledgement_SUCCES}, nil
}

func (fs *frontendServer) Result(ctx context.Context, info *Proto.Info) (*Proto.Details, error) {

	return &Proto.Details{}, nil
}

func (fs *frontendServer) View(ctx context.Context, empty *Proto.EmptyRequest) (*Proto.DetailsList, error) {

	return &Proto.DetailsList{}, nil
}

func (fs *frontendServer) Register(request *Proto.RegisterRequest, stream Proto.AuctionHouse_RegisterServer) error {
	conn := &Connection{
		id:     request.Id,
		stream: stream,
		error:  make(chan error),
	}
	fs.connetions[conn.id] = conn
	return <-conn.error

}
