package frontend

import (
	"context"

	"github.com/emiloh/DISYS_AuctionHouse/tree/Auction/Proto"
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

}

func (fs *frontendServer) Bid(ctx context.Context, offer *Proto.Offer) (*Proto.Acknowledgement, error) {

	return &Proto.Acknowledgement{Status: Proto.Acknowledgement_SUCCES}, nil
}

func (fs *frontendServer) Result(ctx context.Context, info *Proto.Info) (*Proto.Offer, error) {

	return &Proto.Offer{}, nil
}

func (fs *frontendServer) View(ctx context.Context, empty *Proto.EmptyRequest) (*Proto.InfoList, error) {

	return &Proto.InfoList{}, nil
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
