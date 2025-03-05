// main.go - Entry point for the WebSocket server

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/gorilla/websocket"
)

// The main function initializes the WebSocket server and listens for connections.
func main() {
	// Create a new DataReceiver instance
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}

	// Register WebSocket handler
	http.HandleFunc("/ws", recv.handleWS)

	// Start the HTTP server on port 30000
	http.ListenAndServe(":30000", nil)
}

// DataReceiver is responsible for receiving WebSocket messages and producing them to Kafka.
type DataReceiver struct {
	msgch chan types.OBUData // Channel to buffer incoming OBU data
	conn  *websocket.Conn    // WebSocket connection instance
	prod  DataProducer       // Producer interface for sending data to Kafka
}

// NewDataReceiver initializes a DataReceiver with a Kafka producer.
func NewDataReceiver() (*DataReceiver, error) {
	var (
		p          DataProducer
		err        error
		kafkaTopic = "obudata" // Kafka topic to send data
	)

	// Create a new Kafka producer instance
	p, err = NewKafkaProducer(kafkaTopic)
	if err != nil {
		return nil, err
	}

	// Wrap producer with logging middleware
	p = NewLogMiddleware(p)

	return &DataReceiver{
		msgch: make(chan types.OBUData, 128), // Buffered channel for message handling
		prod:  p,                             // Assign producer
	}, nil
}

// produceData sends the received OBUData to Kafka using the producer.
func (dr *DataReceiver) produceData(data types.OBUData) error {
	return dr.prod.ProduceData(data)
}

// handleWS upgrades an HTTP request to a WebSocket connection and starts listening for messages.
func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}

	// Upgrade to WebSocket connection
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn // Store connection in DataReceiver

	// Start a goroutine to continuously read messages
	go dr.wsReceiveLoop()
}

// wsReceiveLoop continuously listens for messages from WebSocket clients.
func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("New OBU client connected!")

	for {
		var data types.OBUData

		// Read and parse incoming JSON message
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			continue
		}

		// Log received message
		fmt.Println("Received message:", data)

		// Send data to Kafka producer
		if err := dr.produceData(data); err != nil {
			fmt.Println("Kafka produce error:", err)
		}
	}
}
