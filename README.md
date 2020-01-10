# Protocol Buffer Tutorial
## From - https://developers.google.com/protocol-buffers/docs/gotutorial

### Setup
`go get -u github.com/golang/protobuf/protoc-gen-go`

#### (Re)compiling go code from `.proto` files

`protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto`

`protoc -I ./proto --go_out=./pkg/addressbook ./proto/addressbook.proto`
