protoc --go_out=./proto --go_opt=paths=import \
    --go-grpc_out=./proto --go-grpc_opt=paths=import \
    assets/inventory.proto assets/bank.proto
