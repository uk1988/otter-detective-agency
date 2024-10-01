# Earthfile
VERSION 0.8

# Define the base stage to build the Go binaries
FROM golang:1.23.1
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .

# Target to build the player service binary
build-playerservice:
    RUN CGO_ENABLED=0 go build -o playerservice cmd/playerservice/main.go
    SAVE ARTIFACT playerservice

# Target to build the case service binary
build-caseservice:
    RUN CGO_ENABLED=0 go build -o caseservice cmd/caseservice/main.go
    SAVE ARTIFACT caseservice

# Target to build the evidence service binary
build-evidenceservice:
    RUN CGO_ENABLED=0 go build -o evidenceservice cmd/evidenceservice/main.go
    SAVE ARTIFACT evidenceservice

# Target to build the interrogation service binary
build-interrogationservice:
    RUN CGO_ENABLED=0 go build -o interrogationservice cmd/interrogationservice/main.go
    SAVE ARTIFACT interrogationservice

# Target to build the deduction service binary
build-deductionservice:
    RUN CGO_ENABLED=0 go build -o deductionservice cmd/deductionservice/main.go
    SAVE ARTIFACT deductionservice

# Target to build the game service binary
build-gameservice:
    RUN CGO_ENABLED=0 go build -o gameservice cmd/gameservice/main.go
    SAVE ARTIFACT gameservice

# Target to build all binaries
build-all:
    BUILD +build-playerservice
    BUILD +build-caseservice
    BUILD +build-evidenceservice
    BUILD +build-interrogationservice
    BUILD +build-deductionservice
    BUILD +build-gameservice

# Target to create the final player service image
docker-playerservice:
    FROM alpine:3.20
    RUN apk --no-cache add bash
    WORKDIR /app
    COPY +build-playerservice/playerservice .
    COPY deploy/docker/scripts/wait-for-it.sh .
    RUN chmod +x /app/wait-for-it.sh
    CMD ["/app/playerservice"]
    SAVE IMAGE vfiftyfive/oda-playerservice:latest

# Target to create the final case service image
docker-caseservice:
    FROM alpine:3.20
    RUN apk --no-cache add bash
    WORKDIR /app
    COPY +build-caseservice/caseservice .
    COPY deploy/docker/scripts/wait-for-it.sh .
    RUN chmod +x /app/wait-for-it.sh
    CMD ["/app/caseservice"]
    SAVE IMAGE vfiftyfive/oda-caseservice:latest

# Target to create the final evidence service image
docker-evidenceservice:
    FROM alpine:3.20
    RUN apk --no-cache add bash
    WORKDIR /app
    COPY +build-evidenceservice/evidenceservice .
    COPY deploy/docker/scripts/wait-for-it.sh .
    RUN chmod +x /app/wait-for-it.sh
    CMD ["/app/evidenceservice"]
    SAVE IMAGE vfiftyfive/oda-evidenceservice:latest

# Target to create the final interrogation service image
docker-interrogationservice:
    FROM alpine:3.20
    RUN apk --no-cache add bash
    WORKDIR /app
    COPY +build-interrogationservice/interrogationservice .
    COPY deploy/docker/scripts/wait-for-it.sh .
    RUN chmod +x /app/wait-for-it.sh
    CMD ["/app/interrogationservice"]
    SAVE IMAGE vfiftyfive/oda-interrogationservice:latest

# Target to create the final deduction service image
docker-deductionservice:
    FROM alpine:3.20
    RUN apk --no-cache add bash
    WORKDIR /app
    COPY +build-deductionservice/deductionservice .
    COPY deploy/docker/scripts/wait-for-it.sh .
    RUN chmod +x /app/wait-for-it.sh
    CMD ["/app/deductionservice"]
    SAVE IMAGE vfiftyfive/oda-deductionservice:latest

# Target to create the final game service image
docker-gameservice:
    FROM alpine:3.20
    WORKDIR /app
    COPY +build-gameservice/gameservice .
    CMD ["/app/gameservice"]
    SAVE IMAGE vfiftyfive/oda-gameservice:latest

# Target to build and save all images
docker-all:
    BUILD +docker-playerservice
    BUILD +docker-caseservice
    BUILD +docker-evidenceservice
    BUILD +docker-interrogationservice
    BUILD +docker-deductionservice
    BUILD +docker-gameservice

# Target to build for multiple platforms
multi:
    BUILD --platform=linux/amd64 --platform=linux/arm64 +docker-all
