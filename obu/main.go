package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

const wsEndpoint = "ws://localhost:30000/ws"

var sendInterval = time.Second * 5

// Generate random coordinates
func genCord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

// Generate random latitude and longitude
func Location() (float64, float64) {
	return genCord(), genCord()
}

// Generate random OBU IDs
func generateOBUID(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(999999)
	}
	return ids
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate OBU IDs
	obuIDs := generateOBUID(20)

	// Connect to the WebSocket server
	var conn *websocket.Conn
	var err error

	// Retry connection logic
	for i := 0; i < 5; i++ { // Retry 5 times
		conn, _, err = websocket.DefaultDialer.Dial(wsEndpoint, nil)
		if err == nil {
			break
		}
		log.Printf("Connection attempt %d failed: %v\n", i+1, err)
		time.Sleep(time.Second * 2) // Wait before retrying
	}
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server after retries: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to WebSocket server")

	// Continuously send data to the server
	for {
		for i := 0; i < len(obuIDs); i++ {
			lat, long := Location()
			data := struct {
				OBUID     int     `json:"obuID"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			}{
				OBUID:     obuIDs[i],
				Latitude:  lat,
				Longitude: long,
			}

			// Print the data being sent
			fmt.Printf("Sending data: %+v\n", data)

			// Send the data as JSON over the WebSocket connection
			if err := conn.WriteJSON(data); err != nil {
				log.Fatalf("Failed to send data: %v", err)
			}
		}

		// Wait before sending the next batch of data
		time.Sleep(sendInterval)
	}
}
