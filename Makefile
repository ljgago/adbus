TARGET = adbus
PROTO_FILES = adbus.proto
PROTO_PATH = api/proto/v1
BUILD_OUT = pkg/api/v1
SWAGGER_OUT = api/swagger/v1
# RPC_BUILD_OUT = pkg/rpc

all: clean build

clean:
	rm -rf $(TARGET)

build:
	go build -o $(TARGET) main.go

run:
	go run .

proto:
	protoc -I$(PROTO_PATH) \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		-Ithird_party/gogo \
		--gogofaster_out=plugins=grpc,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:$(BUILD_OUT) \
	$(PROTO_FILES)
	protoc -I$(PROTO_PATH) \
		-I${GOPATH}/src \
		-Ithird_party/gogo \
		--grpc-gateway_out=allow_patch_feature=false,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:$(BUILD_OUT) \
		$(PROTO_FILES)
	protoc -I$(PROTO_PATH) \
		-I${GOPATH}/src \
		-Ithird_party/gogo \
		--swagger_out=logtostderr=true:$(SWAGGER_OUT) \
		$(PROTO_FILES)
	sed -i.bak "s/empty.Empty/types.Empty/g" $(BUILD_OUT)/$(TARGET).pb.gw.go && rm $(BUILD_OUT)/$(TARGET).pb.gw.go.bak
