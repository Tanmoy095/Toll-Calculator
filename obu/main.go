package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/gorilla/websocket"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

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
			if err := sendData(conn, data); err != nil {
				log.Println("Trying to reconnect...")
				conn, _, err = websocket.DefaultDialer.Dial(wsEndpoint, nil)
				if err != nil {
					log.Println("Reconnection failed:", err)
					time.Sleep(2 * time.Second) // Wait before retrying
					continue
				}
			}
		}
		time.Sleep(sendInterval)
	}

}
func sendData(conn *websocket.Conn, data types.ObuData) error {
	err := conn.WriteJSON(data)
	if err != nil {
		log.Println("WebSocket error:", err)
		conn.Close()
		return err
	}
	return nil
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
		ids[i] = rand.Intn(100000) // Generate smaller OBUIDs
	}
	return ids
}
