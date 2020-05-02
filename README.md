# grpc-test

## To run existing code
- `go get ./...`
- `go run main.go`

## To create new message or service using protobuf

#### Install protoc (You can skip if it is already installed)

##### Download zip file
`curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip`
##### Unzip protoc
`unzip protoc-3.11.4-linux-x86_64.zip -d $HOME/.local`
##### Add $PATH
`export PATH="$PATH:$HOME/.local/bin"`
##### Now check if it is properly installed or not
`protoc --version`

#### go get some packages
- `go get github.com/golang/protobuf/protoc-gen-go`
- `go get google.golang.org/grpc@v1.28.1`
