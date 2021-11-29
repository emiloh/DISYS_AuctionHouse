package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var id string
var io *bufio.Reader

func main() {
	id = os.Args[1]
	io = bufio.NewReader(os.Stdin)
	welcome()
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
