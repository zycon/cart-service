###### Microservice example "cart-service" using gRPC

This modules requires protobuf compiler plugin
install from 
go get -u github.com/golang/protobuf/protoc-gen-go


compile the definiton by 
protoc --go_out=plugins=grpc:. ./protobuf/*.proto
