package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/gorilla/websocket"
)

var kafkaTopic = "obudata"

func main() {
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.handleWs)
	http.ListenAndServe(":30000", nil)
}

type DataReceiver struct {
	msgchn chan types.ObuData
	conn   *websocket.Conn
	prod   Dataproducer
}

func NewDataReceiver() (*DataReceiver, error) {
	p, err := NewKafkaProducer()

	if err != nil {
		return nil, err
	}
	return &DataReceiver{
		msgchn: make(chan types.ObuData, 128),
		prod:   p,
	}, nil
}
func (dr *DataReceiver) produceData(data types.ObuData) error {
	return dr.prod.ProduceData(data)
}
func (dr *DataReceiver) handleWs(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn

	go dr.wsReceiveLoop()

}

//loop over websocket connection so we can keep ranging over message

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("New OBU connected client connected !")
	for {
		var data types.ObuData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			continue

		}

		fmt.Println("received message: ", data)
		if err := dr.produceData(data); err != nil {
			fmt.Println("kafka produce error:", err)
		}
	}

}
