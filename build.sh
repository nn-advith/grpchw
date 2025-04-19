#!/bin/bash

command=$1

case "$command" in
    proto)
        protoc --proto_path=proto --go_out=proto --go-grpc_out=proto proto/*.proto
    ;;
    server)
        go run server/server.go
    ;;
    client)
        go run client/client.go
    ;;
    *)
        echo "Usage: ./build.sh [proto|server|client]"
    ;;
esac