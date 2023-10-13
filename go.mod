module benchmark

go 1.20

//replace github.com/loopholelabs/polyglot => ../polyglot-go

require (
	github.com/loopholelabs/polyglot v1.1.3-0.20231013153026-7a37dc1667af
	google.golang.org/grpc v1.58.3
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect; indirect..
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231012201019-e917dd12ba7a // indirect
)
