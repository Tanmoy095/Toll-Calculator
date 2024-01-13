package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/gorilla/websocket"
)

const wsEndpoint = "ws://127.0.0.1:3000/ws"

var sendInterval = time.Second

func main() {
	obuID := generateOBUIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(obuID); i++ {
			lat, long := genLatLong()
			data := types.ObuData{
				OBUID: obuID[i],
				Lat:   lat,
				Long:  long,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}

		}
		time.Sleep(sendInterval)

	}
}
func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}
func genLatLong() (float64, float64) {
	return genCoord(), genCoord()
}
func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}
