package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// the amount of bidders we have at our auction
const bidderCount = 10

// initial wallet value for all bidders
const walletAmount = 250

// items is the map of auction items
var items = []string{
	"The \"Best Gopher\" trophy",
	"The \"Learn Go with Adelina\" experience",
	"Two tickets to a Go conference",
	"Signed copy of \"Beautiful Go code\"",
	"Vintage Gopher plushie",
}

// bid is a type that pairs the bidder id and the amount they want to bid
type bid struct {
	bidderID string
	amount   int
}

// auctioneer receives bids and announces winners
type auctioneer struct {
	bidders map[string]*bidder
}

// runAuction and manages the auction for all the items to be sold
// Change the signature of this function as required
func (a *auctioneer) runAuction(bidCh <-chan bid, open chan<- struct{}) {
	for _, item := range items {
		var maxBid bid

		log.Printf("Opening bids for %s!\n", item)
		a.openBids(open)
		for i := 0; i < bidderCount; i++ {
			//open <- struct{}{}
			currBid := <-bidCh
			if currBid.amount > maxBid.amount {
				maxBid = currBid
				log.Printf("%s has set the current bid to $%d\n", maxBid.bidderID, maxBid.amount)
			}
		}
		log.Printf("Going once..Going Twice... Sold! To %s for $%d!\n", maxBid.bidderID, maxBid.amount)
		a.bidders[maxBid.bidderID].payBid(maxBid.amount)
	}
}

func (a *auctioneer) openBids(open chan<- struct{}) {
	for i := 0; i < bidderCount; i++ {
		open <- struct{}{}
	}
}

// bidder is a type that holds the bidder id and wallet
type bidder struct {
	id     string
	wallet int
}

// placeBid generates a random amount and places it on the bids channels
// Change the signature of this function as required
func (b *bidder) placeBid(bidCh chan<- bid, open <-chan struct{}) {
	for range items {
		<-open
		maxBid := b.wallet
		bidPlaced := bid{b.id, rand.Intn(maxBid + 1)}
		//log.Printf("%s: %s has placed bid of %d", item, bidPlaced.bidderID, bidPlaced.amount)
		bidCh <- bidPlaced
	}
}

// payBid subtracts the bid amount from the wallet of the auction winner
func (b *bidder) payBid(amount int) {
	b.wallet -= amount
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Welcome to the LinkedIn Learning auction.")
	bidders := make(map[string]*bidder, bidderCount)
	bidCh := make(chan bid, bidderCount)
	openBid := make(chan struct{})
	for i := 0; i < bidderCount; i++ {
		id := fmt.Sprint("Bidder ", i)
		b := bidder{
			id:     id,
			wallet: walletAmount,
		}
		bidders[id] = &b
		go b.placeBid(bidCh, openBid)
	}
	a := auctioneer{
		bidders: bidders,
	}
	a.runAuction(bidCh, openBid)
	log.Println("The LinkedIn Learning auction has finished!")
}

// getRandomAmount generates a random integer amount up to max
func getRandomAmount(max int) int {
	return rand.Intn(int(max))
}
