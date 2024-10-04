# Earthfile
VERSION 0.8

# Define the base stage to build the Go binaries
FROM golang:1.23.1
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .

# Target to build the CSI service binary
build-csiservice:
    RUN CGO_ENABLED=0 go build -o csiservice cmd/csiservice/main.go
    SAVE ARTIFACT csiservice

# Target to build the game service binary
build-gameservice:
    RUN CGO_ENABLED=0 go build -o gameservice cmd/gameservice/main.go
    SAVE ARTIFACT gameservice

# Target to build all binaries
build-all:
    BUILD +build-csiservice
    BUILD +build-gameservice

# Target to create the final CSI service image
docker-csiservice:
    FROM alpine:3.20
    RUN apk --no-cache add bash
    WORKDIR /app
    COPY +build-csiservice/csiservice .
    COPY deploy/docker/scripts/wait-for-it.sh .
    RUN chmod +x /app/wait-for-it.sh
    CMD ["/app/csiservice"]
    SAVE IMAGE --push ghcr.io/vfiftyfive/oda-csiservicev2:latest

# Target to create the final game service image
docker-gameservice:
    FROM alpine:3.20
    WORKDIR /app
    COPY +build-gameservice/gameservice .
    CMD ["/app/gameservice"]
    SAVE IMAGE --push ghcr.io/vfiftyfive/oda-gameservicev2:latest

# Target to build and save all images
docker-all:
    BUILD +docker-csiservice
    BUILD +docker-gameservice

# Target to build for multiple platforms
multi:
    BUILD --platform=linux/amd64 --platform=linux/arm64 +docker-all
