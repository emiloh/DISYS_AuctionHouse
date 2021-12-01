package main

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

var uid string
var io *bufio.Reader
var client Proto.AuctionHouseClient

func main() {
	uid =  os.Args[1]
	io = bufio.NewReader(os.Stdin)
	welcome()

	//done := make(chan int)

	conn, err := grpc.Dial(":1400", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client = Proto.NewAuctionHouseClient(conn)

	//register(uid)

	for {
		text, _ := io.ReadString('\n')
		commands := strings.Fields(text)
		switch commands[0] {
		case "\\bid":
			amount, _ := strconv.ParseInt(commands[1], 10, 64)
			bid(amount)
		case "\\help":
			help()
		case "\\result":
			id, _ := strconv.ParseInt(commands[1], 10, 64)
			result(id)
		case "\\leave":
			break
		default:
			fmt.Println("Please use one of the commands. If you are unsure about them, type \\help.")
		}
	}

}

/*func register(id string) {
	registerRequest := &Proto.RegisterRequest{Id: id}

	stream, err := client.Register(context.Background(), registerRequest)

	if err != nil {
		log.Fatalf("Conenction failed: %v", err)
	}

	go func(str Proto.AuctionHouse_RegisterClient) {
		for {
			msg, msgerr := str.Recv()
			if msgerr != nil {
				log.Fatalf("Failed to receieve message: %v", msgerr)
			}
			if msg.Details != nil {
				log.Printf("Current offer from %s on %s with id %d is %d with a remaining time of %d seconds", msg.Details.User, msg.Details.Name, msg.Details.Id, msg.Details.Amount, msg.Details.Timeleft)
			} else if msg.DetailList != nil {
				ids := msg.DetailList.Id
				for i := 0; i < len(ids); i++ {
					log.Printf("Current offer on %s with id %d is %d with a remaining time of %d seconds", msg.DetailList.Name[i], msg.DetailList.Id[i], msg.DetailList.Amount[i], msg.DetailList.Timeleft[i])
				}
			} else if msg.Acknowledgement != nil {
				log.Printf("Your bid led to %d", msg.Acknowledgement.Status)
			}
		}

	}(stream)
}*/

func bid(amount int64) {
	offer := &Proto.Offer{
		Amount: amount,
		User: uid,
	}
	client.Bid(context.Background(), offer)
}

func result(id int64) {
	info := &Proto.Info{Uid: uid}
	client.Result(context.Background(), info)
}

func help() {
	fmt.Println("____________________ Commands _________________________")
	fmt.Println("											")
	fmt.Println("Following commands are available:		")
	fmt.Println("\\bid <AMOUNT>    (Bids 'amount' on the item)")
	fmt.Println("\\show                 (Lists all items on the auction house)")
	fmt.Println("\\help                 (Lists the available commands)")
	fmt.Println("\\leave                (Leaves the auction house)")
	fmt.Println("_______________________________________________________")
	fmt.Println()
}

func welcome() {
	fmt.Println("____________________ Auction House ____________________")
	fmt.Println("											")
	fmt.Println("Welcome to DISYS Auction House, " + uid + "!		")
	fmt.Println("Here you have the possiblity to bid on different ")
	fmt.Println("items you may like. To aqquire further assist on ")
	fmt.Println("the bidding tool cosider using the command \\help.	")
	fmt.Println("Best of luck and happy bidding, " + uid + "!  	   ")
	fmt.Println("_______________________________________________________")
	fmt.Println()
}
