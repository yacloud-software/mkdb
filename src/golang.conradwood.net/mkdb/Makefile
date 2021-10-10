.PHONY: client test server
all: client server
client:
	cd client && go install mkdb-client.go
server:
	cd server && go install mkdb-server.go
test: client
	${GOBIN}/mkdb-client >test/thing.go
	cd test && go install  mkdb-test.go thing.go

