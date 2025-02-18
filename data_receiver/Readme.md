🚗 Toll Calculator - Data Receiver Service

This project is responsible for receiving GPS data from vehicles (OBUs) via WebSocket and sending it to Kafka for further processing.

📌 Overview

OBUs (vehicles) send GPS data (latitude, longitude, OBU ID) to the WebSocket server.

WebSocket server (DataReceiver) receives the data and forwards it to Kafka.

Kafka Producer sends the data to the Kafka topic (obudata).

Another Kafka consumer (not part of this service) will later process the data.

🏗 Project Structure

📂 toll-calculator
├── 📄 main.go # WebSocket server & Kafka integration
├── 📄 producer.go # Kafka producer implementation
├── 📄 obu.go # OBU simulator (for testing)
├── 📄 types.go # Data structures for OBU data
├── 📄 README.md # Project documentation

🛠 Setup Instructions

1️⃣ Install Dependencies

go mod tidy

2️⃣ Start Kafka & Zookeeper (If not running)

zookeeper-server-start.sh config/zookeeper.properties
kafka-server-start.sh config/server.properties

3️⃣ Create Kafka Topic

kafka-topics.sh --create --topic obudata --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1

4️⃣ Run the Data Receiver Service

go run main.go

5️⃣ Run the OBU Simulator (To send test data)

go run obu.go

📝 Code Breakdown

1️⃣ WebSocket Server (main.go)

Starts an HTTP WebSocket server on port 30000.

When an OBU (vehicle) connects, it starts listening for GPS data.

The received data is sent to Kafka.

2️⃣ Kafka Producer (producer.go)

Responsible for sending GPS data to Kafka.

Uses Goroutines to check message delivery status.

3️⃣ OBU Simulator (obu.go)

Simulates vehicles sending random GPS coordinates.

Connects to the WebSocket server and sends messages at 1-second intervals.

4️⃣ Types (types.go)

Defines the ObuData struct:

type ObuData struct {
OBUID int `json:"obuId"`
Lat float64 `json:"lat"`
Long float64 `json:"long"`
}

❌ Common Issues & Fixes

🛑 Infinite Recursion in produceData()

Problem: Function was calling itself instead of Kafka Producer.
Fix: Update main.go

func (dr \*DataReceiver) produceData(data types.ObuData) error {
return dr.prod.ProduceData(data) // Correct function call
}

🛑 Kafka Not Running

Problem: kafka-server-start.sh is not started.
Fix: Run Kafka before starting the service.

🛑 WebSocket Connection Refused

Problem: WebSocket server is not running.
Fix: Make sure main.go is running on port 30000.

🚀 Future Improvements

Implement a Kafka Consumer to process and store toll calculations.

Add Docker support for easy deployment.

Implement authentication for WebSocket connections.

🎯 Conclusion

This service receives GPS data from OBUs and sends it to Kafka, where it can be processed further. It is designed using WebSockets and Kafka for real-time, scalable toll calculation. 🚀

🔹 Authors:
🔹 License: MIT
