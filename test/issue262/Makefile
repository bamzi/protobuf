regenerate:
	go install github.com/bamzi/protobuf/protoc-min-version
	go install github.com/bamzi/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/bamzi/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/bamzi/protobuf/types:. timefail.proto
