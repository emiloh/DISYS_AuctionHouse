package frontend

import (
	"context"

	"github.com/emiloh/DISYS_AuctionHouse/tree/Auction/Proto"
)

type Connection struct {
	id string
	stream proto.AuctionHouse_RegisterServer
	error chan error
}
type frontendServer struct {
	proto.UnimplementedAuctionHouseServer
	connetions map[string]*Connection
}

func main() {

}

func (fs *frontendServer) Bid(ctx context.Context, offer *proto.Offer) (*proto.Acknowledgement, error) {

	return &proto.Acknowledgement{Status: proto.Acknowledgement_SUCCES}, nil
}

func (fs *frontendServer) Result(ctx context.Context, info *proto.Info) (*proto.Offer, error) {

	return &proto.Offer{}, nil
}

func (fs *frontendServer) View(ctx context.Context, empty *proto.EmptyRequest) (*proto.InfoList, error) {

	return &proto.InfoList{}, nil
}

func (fs *frontendServer) Register(request *proto.RegisterRequest, stream proto.AuctionHouse_RegisterServer) error {
	conn := &Connection{
		id: request.Id,
		stream: stream,
		error: make(chan error),
	}
	fs.connetions[conn.id] = conn
	return <- conn.error
}