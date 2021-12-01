package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/emiloh/DISYS_AuctionHouse/tree/simpler/Proto"
	"google.golang.org/grpc"
)

var uid string
var io *bufio.Reader
var client Proto.AuctionHouseClient

func main() {
	uid =  os.Args[1]
	dialPort := os.Args[2]
	io = bufio.NewReader(os.Stdin)
	welcome()

	conn, err := grpc.Dial(dialPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client = Proto.NewAuctionHouseClient(conn)

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
			result()
		case "\\leave":
			os.Exit(0)
		default:
			fmt.Println("Please use one of the commands. If you are unsure about them, type \\help.")
		}
	}

}

func bid(amount int64) {
	offer := &Proto.Offer{
		Amount: amount,
		User: uid,
	}
	ack, err := client.Bid(context.Background(), offer)
	if err != nil {
		log.Fatalf("Failed to bid on the auction: %v", err)
	}
	if ack.Response == Proto.Ack_SUCCES {
		log.Println("Your bid was accepted.")
	}else if ack.Response == Proto.Ack_SOLD{
		log.Println("The auction has ended. It is not possible to bid anymore.")
	}else {
		log.Println("Unfortunatly, your bid was not high enough.")
	}
}

func result() {
	info := &Proto.Info{Uid: uid}
	details, err := client.Result(context.Background(), info)
	if err != nil {
		log.Fatalf("Failed to retrieve current status of auction: %v", err)
	}
	if details.Timeleft == 0 {
		log.Printf("The acution is over. %s won %s with the highest bid of %d", details.User, details.Name, details.Amount)
	}else{
		log.Printf("%s has bid %d on %s. %d seconds left.", details.User, details.Amount, details.Name, details.Timeleft)
	}
	
}

func help() {
	fmt.Println("____________________ Commands _________________________")
	fmt.Println("											")
	fmt.Println("Following commands are available:		")
	fmt.Println("\\bid <AMOUNT>    (Bids 'amount' on the item)")
	fmt.Println("\\result                 (Shows the current status of the auction)")
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
