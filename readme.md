###### Microservice example "cart-service" using gRPC

This modules requires pb compiler plugin
install from 
go get -u github.com/golang/pb/protoc-gen-go


compile the definiton by 
protoc --go_out=plugins=grpc:. ./pb/*.proto
