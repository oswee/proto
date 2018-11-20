protoc -I . -I dms/v1 --go_out=plugins=grpc:%GOPATH%/src/ dms/v1/dms.proto
protoc -I . -I app/v1 --go_out=plugins=grpc:%GOPATH%/src/ app/v1/app.proto
protoc -I . -I crm/v1 --go_out=plugins=grpc:%GOPATH%/src/ crm/v1/crm.proto
protoc -I . -I metric/v1 --go_out=plugins=grpc:%GOPATH%/src/ metric/v1/metric.proto
protoc -I . --go_out=plugins=grpc:%GOPATH%/src/ protobuf.proto