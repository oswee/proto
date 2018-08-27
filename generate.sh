protoc -I customer/ customer/customer.proto --go_out=plugins=grpc:customer/go
protoc -I application/ application/application.proto --go_out=plugins=grpc:application/go
protoc -I shipping/ shipping/shipping.proto --go_out=plugins=grpc:shipping/go
protoc -I metric/ metric/metric.proto --go_out=plugins=grpc:metric/go
protoc -I v1/ v1/os.proto --go_out=plugins=grpc:v1

protoc -I routeguide/ routeguide/route_guide.proto --go_out=plugins=grpc:routeguide

protoc --go_out=D:/Code/Go/src dms/v1/dms.proto
protoc --go_out=D:/Code/Go/src app/v1/app.proto
protoc --go_out=D:/Code/Go/src crm/v1/crm.proto
protoc --go_out=D:/Code/Go/src metric/v1/metric.proto
protoc --go_out=D:/Code/Go/src protobuf.proto