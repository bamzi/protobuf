regenerate:
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/bamzi/protobuf/protobuf/ \
	--gogo_out=. issue312.proto
	(cd events && make regenerate)
