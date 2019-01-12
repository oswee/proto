build:
	protoc	-I.	-Idms/v1	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis	--go_out=plugins=grpc:$(GOPATH)/src/	dms/v1/dms.proto

# protoc	-I.	-Idms/v1	--go_out=plugins=grpc:$GOPATH/src/	dms/v1/dms.proto
# protoc	-I.	-Idms/v1	-I/usr/local/include	-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis	--go_out=plugins=grpc:$GOPATH/src/	dms/v1/dms.proto
# protoc	-I.	-Idms/v1	-I/usr/local/include	-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis	--grpc-gateway_out=logtostderr=true:$GOPATH/src/	dms/v1/dms.proto

# To run Makefile execute
# make build

# Common ERRORS
# - google/api/annotations.proto: File not found.
# This import path should be added to include it
# -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis