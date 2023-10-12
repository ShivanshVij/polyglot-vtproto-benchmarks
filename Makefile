generate:
	- mkdir -p vtproto && protoc --go_out=vtproto --go-vtproto_out=vtproto bench.proto
	- mkdir -p polyglot && protoc --go-polyglot_out=polyglot bench.proto

benchmark:
	- go test -bench=. ./... -benchtime=100000000x