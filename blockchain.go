package main

import (
	"fmt"
	"time"
)

type Coin struct {
	docType          string    `json:"docType`
	appraisedValue   float64   `json:"appraisedValue"`
	generatedTime    time.Time `json:"generatedTime"`
	purchasedTime    time.Time `json:"purchasedTime"`
	id               int       `json:"id"`
	largeCategory    string    `json:"largeCategory"`
	latitude         float64   `json:"latitude"`
	longitude        float64   `json:"longitude"`
	owner            string    `json:"owner"`
	producer         string    `json:"producer"`
	smallCategory    string    `json:"smallCategory"`
	status           string    `json:"status"`
	lastAuctionner   string    `json:"lastAuctionner"`
	auctionStartTime time.Time `json:"auctionStartTime"`
}

// Coin methods / technically cannot exist

func (coin Coin) print() {
	fmt.Printf("%+v\n", coin)
}

type Blockchain map[int]*Coin

// Blockchain methods / technically cannot exist

func (blockchain Blockchain) print() {
	for _, coin := range blockchain {
		fmt.Printf("%+v\n\n", coin)
	}
}

func (blockchain Blockchain) showOwners() {
	for _, coin := range blockchain {
		fmt.Print(coin.id, " is owned by : ", coin.owner, "\n")
	}
	fmt.Print("\n")
}

func (blockchain Blockchain) showLastAuctionner() {
	for _, coin := range blockchain {
		fmt.Print(coin.id, " was last auctionned by : ", coin.lastAuctionner, "\n")
	}
	fmt.Print("\n")
}

// initialize functions

func createRandomBlockchain(n int, x int, y int) Blockchain {
	blockchain := make(Blockchain)
	for i := 0; i < n; i++ {
		blockchain[i] = &Coin{
			docType:        "coin",
			appraisedValue: float64(randInt(0, 10)),
			generatedTime:  time.Now(),
			id:             i,
			latitude:       float64(randInt(0, x)),
			longitude:      float64(randInt(0, y)),
			producer:       "windmill",
			smallCategory:  "green",
		}
	}
	return blockchain
}

func createBlockchain() Blockchain {
	return Blockchain{
		0: &Coin{
			docType:        "coin",
			appraisedValue: 5,
			generatedTime:  time.Now(),
			id:             0,
			latitude:       0,
			longitude:      0,
			producer:       "windmill",
			smallCategory:  "green",
		},
		1: &Coin{
			docType:        "coin",
			appraisedValue: 2,
			generatedTime:  time.Now(),
			id:             1,
			latitude:       1,
			longitude:      0,
			producer:       "windmill",
			smallCategory:  "green",
		},
		2: &Coin{
			docType:        "coin",
			appraisedValue: 7,
			generatedTime:  time.Now(),
			id:             2,
			latitude:       0,
			longitude:      1,
			producer:       "windmill",
			smallCategory:  "green",
		},
	}
}
