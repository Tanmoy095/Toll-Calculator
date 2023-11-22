package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var sendInterval = time.Second

type ObuData struct {

	//we will send websocket as json
	OBUID int `json:"obuID"`
	//lotitude
	Lat float64 `json:"lat"`
	//longitude
	Long float64 `json:"long"`
}

func main() {
	obuID := generateOBUIDS(20)
	for {
		for i := 0; i < len(obuID); i++ {
			lat, long := genLatLong()
			data := ObuData{
				OBUID: obuID[i],
				Lat:   lat,
				Long:  long,
			}
			fmt.Printf("%+v\n", data)

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
