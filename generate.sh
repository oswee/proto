protoc -I customer/ customer/customer.proto --go_out=plugins=grpc:customer/go
protoc -I application/ application/application.proto --go_out=plugins=grpc:application/go
protoc -I shipping/ shipping/shipping.proto --go_out=plugins=grpc:shipping/go