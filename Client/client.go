package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/emiloh/DISYS_AuctionHouse/tree/Auction/Proto"
	"google.golang.org/grpc"
)

var id string
var io *bufio.Reader
var client proto.AuctionHouseClient

func main() {
	id = os.Args[1]
	io = bufio.NewReader(os.Stdin)
	welcome()

	//done := make(chan int)

	conn, err := grpc.Dial(":1400", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client = proto.NewAuctionHouseClient(conn)

	register(id)

	for {
		text, _ := io.ReadString('\n')
		commands := strings.Fields(text)
		switch commands[0] {
		case "\\bid":
			amount, _ := strconv.ParseInt(commands[2], 10, 64)
			bid(commands[1], int64(amount))
		case "\\show":
		case "\\help":
			help()
		case "\\leave":
		default:
			fmt.Println("Please use one of the commands. If you are unsure about them, type \\help.")
		}
	}

}

func register(id string) {
	registerRequest := &proto.RegisterRequest{Id: id}

	stream, err := client.Register(context.Background(), registerRequest)

	if err != nil {
		log.Fatalf("Conenction failed: %v", err)
	}

	go func(str proto.AuctionHouse_RegisterClient) {
		for {
			msg, msgerr := str.Recv()
			if msgerr != nil {
				log.Fatalf("Failed to receieve message: %v", msgerr)
			}
			if msg.Info != nil {
				log.Printf("Current offer on %s with id %d is %d with a remaining time of %d seconds", msg.Info.name, msg.Info.id, msg.Info.amount, msg.Info.timeleft)
			} else if msg.InfoList != nil {
				log.Printf("")
			}
		}

	}(stream)
}

func bid(id string, amount int64) {

}

func show() {

}

func help() {
	fmt.Println("____________________ Commands _________________________")
	fmt.Println("											")
	fmt.Println("Following commands are available:		")
	fmt.Println("\\bid <ID> <AMOUNT>    (Bids 'amount' on item with 'id')")
	fmt.Println("\\show                 (Lists all items on the auction house)")
	fmt.Println("\\help                 (Lists the available commands)")
	fmt.Println("\\leave                (Leaves the auction house)")
	fmt.Println("_______________________________________________________")
	fmt.Println()
}

func leave() {

}

func welcome() {
	fmt.Println("____________________ Auction House ____________________")
	fmt.Println("											")
	fmt.Println("Welcome to DISYS Auction House, " + id + "!		")
	fmt.Println("Here you have the possiblity to bid on different ")
	fmt.Println("items you may like. To aqquire further assist on ")
	fmt.Println("the bidding tool cosider using the command \\help.	")
	fmt.Println("Best of luck and happy bidding, " + id + "!  	   ")
	fmt.Println("_______________________________________________________")
	fmt.Println()
}
