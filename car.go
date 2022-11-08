package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Car struct {
	ip             string
	latitude       float64
	longitude      float64
	energyLevel    int
	connected      bool
	auctionning    *Coin
	lastAuctionner string // ip of the auctionner before, if first own ip
}

// Coin methods included in car.go

func (coin Coin) distance(latitude float64, longitude float64) float64 {
	return math.Sqrt(math.Pow(coin.latitude-latitude, 2) + math.Pow(coin.longitude-longitude, 2))
}

func (coin Coin) cost(latitude float64, longitude float64, energyLevel int) float64 {
	return coin.appraisedValue + coin.distance(latitude, longitude) + 10*float64(energyLevel/100)
}

// Car methods

func (car *Car) findCheepestCoin(blockchain Blockchain) *Coin {
	cheepestCoin := &Coin{id: -1, appraisedValue: math.MaxFloat64}
	for i, coin := range blockchain {
		if coin.cost(car.latitude, car.longitude, car.energyLevel) < cheepestCoin.cost(car.latitude, car.longitude, car.energyLevel) && coin.owner == "" {
			cheepestCoin = blockchain[i]
		}
	}
	return cheepestCoin
}

func (car *Car) queueForCoin(coin *Coin) {
	// queue for coin
	if coin.auctionStartTime == (time.Time{}) {
		coin.auctionStartTime = time.Now()
	}
	if coin.lastAuctionner != "" && coin.lastAuctionner != car.ip {
		car.lastAuctionner = coin.lastAuctionner
	} else {
		car.lastAuctionner = ""
	}
	coin.lastAuctionner = car.ip
	car.auctionning = coin
}

func (car *Car) getAuctionners(cars Cars) []*Car {
	// get all auctionners
	auctionners := []*Car{}
	auctionner := car
	for auctionner != nil {
		auctionners = append(auctionners, auctionner)
		auctionner = cars.contactCarByIp(auctionner.lastAuctionner)
	}
	return auctionners
}

func (car *Car) auction(cars Cars) {
	if car.readyForAuction() && car.isLastAuctionner() {
		coin := car.auctionning
		auctionners := car.getAuctionners(cars)
		winnerId := 0
		for i, auctionner := range auctionners {
			if auctionner.energyLevel < auctionners[winnerId].energyLevel && coin.distance(auctionner.latitude, auctionner.longitude) <= coin.distance(auctionners[winnerId].latitude, auctionners[winnerId].longitude) {
				winnerId = i
			}
		}
		if *auctionners[winnerId] != (Car{}) && *coin != (Coin{}) {
			auctionners[winnerId].connected = true
			coin.owner = auctionners[winnerId].ip
			coin.purchasedTime = time.Now()
			coin.auctionStartTime = time.Time{}
			cars.purgeAuctionnedCoin(*coin)
		}
	}
}

func (car Car) isLastAuctionner() bool {
	if car.auctionning.lastAuctionner == car.ip {
		return true
	} else {
		return false
	}
}

func (car *Car) readyForAuction() bool {
	if !car.connected && car.auctionning != nil && time.Since(car.auctionning.auctionStartTime) > 500*time.Millisecond {
		return true
	} else {
		return false
	}
}

func (car Car) print() {
	fmt.Printf("%+v\n", car)
}

type Cars map[int]*Car

// Cars methods

func (cars Cars) print() {
	for _, car := range cars {
		fmt.Printf("%+v\n\n", car)
	}
}

func (cars Cars) showAuctions() {
	for _, car := range cars {
		if !car.connected && car.auctionning != nil {
			fmt.Printf("Car %s is auctionning coin %d\n", car.ip, car.auctionning.id)
		}
	}
}

func (cars Cars) showLastAuctionner() {
	for _, car := range cars {
		if !car.connected {
			fmt.Printf("Car %s last auctionner is %s\n", car.ip, car.lastAuctionner)
		}
	}
}

func (cars *Cars) connect(blockchain Blockchain) Blockchain {
	// connect all cars asynchroniously to energy
	for i, car := range *cars {
		if car.auctionning == nil && car.connected == false {
			car.queueForCoin(car.findCheepestCoin(blockchain))
			(*cars)[i] = car
		}
	}
	return blockchain
}

func (cars *Cars) purgeAuctionnedCoin(coin Coin) {
	for i, car := range *cars {
		if car.auctionning != nil && car.auctionning.id == coin.id {
			car.auctionning = nil
			car.lastAuctionner = ""
			(*cars)[i] = car
		}
	}
}

func (cars *Cars) updateByIp(ip string, car Car) {
	for i, c := range *cars {
		if c.ip == ip {
			(*cars)[i] = &car
		}
	}
}

func (cars *Cars) doAuction(blockchain Blockchain) Blockchain {
	// launch cars auctions
	for _, car := range *cars {
		car.auction(*cars)
	}
	return blockchain
}

func (cars Cars) allCarsConnected() bool {
	for _, car := range cars {
		if !car.connected {
			return false
		}
	}
	return true
}

func (cars Cars) contactCarByIp(ip string) *Car {
	for _, car := range cars {
		if car.ip == ip {
			return car
		}
	}
	return nil
}

// initialize functions

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func createCars(n int, x int, y int) Cars {
	cars := make(Cars)
	for i := 0; i < n; i++ {
		cars[i] = &Car{
			ip:          "192.168.0." + strconv.Itoa(i),
			latitude:    float64(randInt(0, x)),
			longitude:   float64(randInt(0, y)),
			connected:   false,
			energyLevel: randInt(0, 100),
			auctionning: nil,
		}
	}
	return cars
}
