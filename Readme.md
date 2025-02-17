a# Toll Management System for Trucks

This project is a microservices-based system designed to manage toll infrastructure for trucks in Belgium. The system calculates toll fees based on GPS coordinates sent by onboard units (OBUs) in trucks, considering factors such as distance traveled, vehicle weight, and emissions. The system is built incrementally, with each microservice handling a specific task, ensuring scalability, reliability, and fault tolerance.

---

## System Overview

The system is composed of the following microservices:

1. **Receiver Service**:

   - Collects GPS coordinates from OBUs in trucks.
   - Publishes the data to Kafka for reliable storage and reprocessing.

2. **Distance Calculator Service**:

   - Pulls GPS data from Kafka.
   - Calculates the distance traveled on toll roads.
   - Factors in toll road charges based on distance, weight, and emissions.

3. **Invoice Service**:
   - Generates invoices based on the calculated distance and toll fees.
   - Stores invoices in a database.
   - Supports additional features like pre-checking toll fees for transportation companies.

---

## Key Features

- **Incremental Development**: The system is built step-by-step, allowing for continuous learning and adaptation.
- **Fault Tolerance**: Kafka ensures reliable messaging and the ability to replay events in case of missing or incorrect data.
- **Scalability**: Each microservice is designed to perform a single task efficiently, enabling horizontal scaling.
- **Monitoring**: Prometheus and Grafana are used for metrics collection and visualization.

---

## Tech Stack

- **Programming Language**: Go (Golang) for building microservices.
- **Containerization**: Docker for packaging and deploying services.
- **Messaging System**: Kafka for reliable data streaming and event replay.
- **Monitoring**: Prometheus for metrics collection and Grafana for visualization.
- **Database**: A time-series database for storing toll data (e.g., InfluxDB or TimescaleDB).

---

## Getting Started

### Prerequisites

- Docker and Docker Compose installed.
- Go installed (if running services locally).
- Kafka and Zookeeper set up (can be done via Docker).

### Installation
