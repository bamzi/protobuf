regenerate:
	go install github.com/bamzi/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/bamzi/protobuf/protobuf/ \
	--gogo_out=. events.proto
