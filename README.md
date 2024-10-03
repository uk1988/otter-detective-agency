# 🦦 Otter Detective Agency: A gRPC Microservices Game

## 🎯 Purpose

The Otter Detective Agency is an educational project designed to demonstrate the implementation of gRPC, Protocol Buffers, and microservices architecture using PostgreSQL as a persistent store. This application showcases how gRPC services interact with a database in a real-world scenario.

## 🏗️ Architecture

The application is composed of several microservices:

1. PlayerService: Manages player data and progress
2. CaseService: Handles case management and assignment
3. EvidenceService: Manages evidence related to cases
4. InterrogationService: Handles suspect interrogations
5. DeductionService: Processes case solutions
6. GameService: Orchestrates the game flow and user interactions

These services communicate via gRPC, with Protocol Buffers defining the service interfaces.

### Service Relationships

- The **Game Service** communicates with all other services to orchestrate the game flow.
- Each service interacts with the **PostgreSQL** database to persist and retrieve data.

### Branches

- **main**: Contains the current stable version of the application.
- **csi**: A v2 version of the application featuring the **CSI Service**, which adds forensic analysis capabilities.

## 🚀 Getting Started

### Prerequisites

- **Docker** and **Docker Compose** installed on your machine.
- **Go** (if you plan to run the services manually).
- **wscat** (WebSocket cat) installed globally for testing WebSocket connections.

### Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/otter-detective-agency.git
cd otter-detective-agency
```

## 🎮 Usage

### Starting the Application with Docker Compose

Start all services using Docker Compose:

```bash
cd deploy/docker
docker-compose up --build
```

This command builds and starts all the services defined in the `docker-compose.yml` file.

### Starting a New Game

#### Step 1: Get the WebSocket URL

Run the following command in your terminal to start a new game and retrieve the WebSocket URL:

```bash
NEW_GAME_URL=$(curl -s http://localhost:8080/startnewgame | sed -n 's/.*\(ws:\/\/[^ ]*\)/\1/p')
```

#### Step 2: Connect to the Game

Use `wscat` to connect to the game:

```bash
wscat --connect $NEW_GAME_URL
```
#### Step 3: Play the Game

In the Otter Detective Agency game, you take on the role of an otter detective solving intriguing cases. The game is played entirely through the terminal using WebSocket connections.

### Gameplay Overview

- **Start the Game**: Upon connecting, you'll be prompted to enter your detective name.
- **Case Assignment**: You'll receive a case file with details about the mystery you need to solve.
- **Exploration**:
  - **Locations**: Type `locations` to see a list of places you can investigate.
  - **Examine Locations**: Use `examine [location]` to look for clues in a specific area.
- **Evidence Collection**:
  - **List Evidence**: After examining a location, you'll see available evidence.
  - **Investigate Evidence**: Use `investigate [evidence name]` to inspect items closely.
  - Optionally, you can use the **CSI Service** to analyze evidence further. This will give you 7/10 chances to get a new clue.
- **Suspect Interaction**:
  - **List Suspects**: Type `suspects` to see potential suspects.
  - **Interrogate Suspects**: Use `interrogate [suspect name]` to question individuals.
  - **Ask Questions**: During interrogation, type `ask [question number]` to ask a specific question.
- **Solving the Case**:
  - **Make a Deduction**: When ready, type `solve` to attempt to solve the case.
  - **Submit Solution**: You'll be prompted to name the culprit. Enter your guess to see if you're correct.

### Navigation Tips

- **Commands**: Follow on-screen prompts and use the commands provided to navigate through the game.
- **Help**: If you get stuck, the game will often remind you of available actions.
- **Exit the Game**: You can quit the game at any time by pressing `Ctrl+C`.

## 🌿 Branches

- **main**: The primary development branch.
- **csi**: Contains the v2 version with the **CSI Service** for advanced forensic features.

Switch to the `csi` branch to explore the new CSI features:

```bash
git checkout csi
```

## 🐳 Deployment

### Docker Compose

The application can be easily deployed using Docker Compose. All services are defined in the `docker-compose.yml` file, which sets up the necessary networking and environment variables.

### Kubernetes

*Coming Soon!* A Kubernetes deployment configuration will be provided in future updates.

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/my-feature`).
3. Commit your changes (`git commit -am 'Add my feature'`).
4. Push to the branch (`git push origin feature/my-feature`).
5. Open a Pull Request.

Please make sure to update tests as appropriate.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
