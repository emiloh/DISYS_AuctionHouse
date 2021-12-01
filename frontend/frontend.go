package main

import (
	"context"
	"log"
	"net"

	"github.com/emiloh/DISYS_AuctionHouse/tree/simpler/Proto"
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

var clients []Proto.AuctionHouseClient
var streams []Proto.AuctionHouse_RegisterClient
var feid string 

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	clients = append(clients, Proto.NewAuctionHouseClient(conn))

	register("0");


	connections := make(map[string]*Connection)

	server := &frontendServer{Proto.UnimplementedAuctionHouseServer{}, connections}

	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", ":1400")

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

func (fs *frontendServer) View(ctx context.Context, user *Proto.User) (*Proto.DetailsList, error) {
	clients[0].View(context.Background(), &Proto.User{})
	msg, err := streams[0].Recv()
	if err != nil {}
	fs.connetions["hah"].stream.Send(msg)
	detailsList := msg.DetailList
	
	return detailsList, nil
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


func register(id string)  {
	registerRequest := &Proto.RegisterRequest{Id: id}

	stream, err := clients[0].Register(context.Background(), registerRequest)
	
	if err != nil {}

	streams = append(streams, stream)
}