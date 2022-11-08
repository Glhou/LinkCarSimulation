package main

import (
	"fmt"
	"time"
)

func main() {
	NCoin, NCar := 20, 20
	N, M := 50, 50
	blockchain := createRandomBlockchain(NCoin, N, M)
	cars := createCars(NCar, N, M)
	i := 1
	for !cars.allCarsConnected() {
		blockchain = cars.connect(blockchain)
		//Blockchain.print()
		//Cars.print()
		blockchain = cars.doAuction(blockchain)
		//Cars.showAuctions()
		//Cars.showLastAuctionner()
		blockchain.showOwners()
		//Blockchain.showLastAuctionner()
		draw(blockchain, cars, N, M, i)
		i++
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Print(cars.allCarsConnected())
	fmt.Print("\n\n\n")
	cars.print()
	fmt.Print("\n\n\n")
	blockchain.print()
}
