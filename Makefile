generate:
	- mkdir -p vtproto && protoc --go_out=vtproto --go-vtproto_out=vtproto bench.proto
	- mkdir -p polyglot && protoc --go-polyglot_out=polyglot bench.proto

benchmark:
	- go test -bench=.  ./...

leaks:
	- go test -bench=. -gcflags="-m=2" ./...