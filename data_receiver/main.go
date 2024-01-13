package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/gorilla/websocket"
)

func main() {
	recv := NewDataReceiver()
	http.HandleFunc("/ws", recv.handleWs)
	http.ListenAndServe(":3000", nil)
}

type DataReceiver struct {
	msgchn chan types.ObuData
	conn   *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgchn: make(chan types.ObuData, 128),
	}
}

func (dr *DataReceiver) handleWs(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
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

		fmt.Printf("received OBU data from [%d] :: <lat %.2f, long %.2f> \n", data.OBUID, data.Lat, data.Long)
		dr.msgchn <- data
	}

}
