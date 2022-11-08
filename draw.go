package main

import (
	"math"
	"strconv"

	"github.com/fogleman/gg"
)

func draw(blockchain Blockchain, cars Cars, N int, M int, i int) {
	const S = 1024
	var LatPas = float64(S / N)
	var LongPas = float64(S / M)
	var radius = math.Min(LatPas, LongPas)
	var lineweight = math.Min(LatPas, LongPas) / 10
	dc := gg.NewContext(S, S)
	for _, coin := range blockchain {
		// draw Coins in blue
		dc.DrawCircle(coin.latitude*LatPas, coin.longitude*LongPas, radius)
		dc.SetRGB(0, 0, 1)
		dc.Fill()
		// draw Coin line with their owner
		if coin.owner != "" {
			dc.SetLineWidth(lineweight)
			dc.SetRGB(0, 1, 0)
			dc.DrawLine(coin.latitude*LatPas, coin.longitude*LongPas, cars.contactCarByIp(coin.owner).latitude*LatPas, cars.contactCarByIp(coin.owner).longitude*LongPas)
			dc.Stroke()
		}
	}
	for _, car := range cars {
		// draw Cars in red
		dc.DrawCircle(car.latitude*LatPas, car.longitude*LongPas, radius)
		dc.SetRGB(1, 0, 0)
		dc.Fill()
		// draw Car dashed line with their auctionning coin
		if car.auctionning != nil {
			dc.SetLineWidth(lineweight)
			dc.SetRGB(1, 1, 0)
			dc.SetDash(10, 10)
			dc.DrawLine(car.latitude*LatPas, car.longitude*LongPas, car.auctionning.latitude*LatPas, car.auctionning.longitude*LatPas)
			dc.Stroke()
			dc.SetDash(0, 0)
		}
	}
	if i != 0 {
		dc.SavePNG("result/out" + strconv.Itoa(i) + ".png")
	}
	dc.SavePNG("out.png")
}
