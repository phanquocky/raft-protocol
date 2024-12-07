# Raft Consensus Algorithm Implementation

This repository contains an implementation of the Raft consensus algorithm in Go. The project is structured to run multiple server instances that communicate with each other to maintain a consistent state.

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go 1.23.3 or later

### Building and Running

1. **Clone the repository:**

    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. **Build and run the Docker containers:**

    ```sh
    docker-compose up --build
    ```

3. **Run the client:**
The client folder is used to interact with the servers. You can run the client to send RPC calls to the servers.

    ```sh
    docker exec -it raftserver1 /bin/bash
    cd client
    go run main.go
    ```

## Project Components

### Main Server

The main server is implemented in [main.go](main.go). It initializes the consensus handler, registers the RPC server, and connects to peers.

### Consensus Handler

The consensus handler is implemented in the [consensus/handler](consensus/handler) directory. It includes various components such as:

- `append_entries.go`
- `broadcast.go`
- `execute.go`
- `get_log.go`
- `handler.go`
- `leader_election.go`
- `request_vote.go`

### Peer Connections

Peer connections are managed in [consensus/external/peer/peer.go](consensus/external/peer/peer.go). It handles connecting to other peers and maintaining the connections.

### Docker Configuration

The Docker configuration is defined in [docker-compose.yml](docker-compose.yml) and [Dockerfile](Dockerfile). It sets up multiple server instances with predefined IP addresses.

## Usage

After starting the servers using Docker Compose, you can run the client to interact with the servers. The client sends RPC calls to execute commands and retrieve logs.

## Acknowledgements

- [Raft Consensus Algorithm](https://raft.github.io/)

This README file was generate by GitHub Copilot.
