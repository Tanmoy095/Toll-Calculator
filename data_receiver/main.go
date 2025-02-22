package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/gorilla/websocket"
)

var kafkaTopic = "obudata"

func main() {
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal("Error initializing data receiver:", err)
	}

	http.HandleFunc("/ws", recv.handleWs)

	fmt.Println("WebSocket server starting on port 30000...")
	err = http.ListenAndServe(":30000", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

type DataReceiver struct {
	msgchn chan types.OBUData
	conn   *websocket.Conn
	prod   Dataproducer
}

// In NewDataReceiver() function
func NewDataReceiver() (*DataReceiver, error) {
	var (
		p   Dataproducer
		err error
	)

	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		p, err = NewKafkaProducer()
		if err == nil {
			break
		}
		log.Printf("Kafka connection attempt %d failed: %v\n", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to Kafka after %d retries: %v", maxRetries, err)
	}

	p = NewLogMiddleware(p)
	return &DataReceiver{
		msgchn: make(chan types.OBUData, 128),
		prod:   p,
	}, nil
}
func (dr *DataReceiver) produceData(data types.OBUData) error {
	return dr.prod.ProduceData(data)
}

func (dr *DataReceiver) handleWs(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	dr.conn = conn
	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("New OBU client connected!")
	defer func() {
		if dr.conn != nil {
			dr.conn.Close()
		}
	}()

	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("WebSocket read error:", err)
			break // Exit the loop on error
		}

		fmt.Println("Received message:", data)
		if err := dr.produceData(data); err != nil {
			log.Println("Kafka produce error:", err)
		}
	}
}
