.PHONY: client test server

all: client server
client:
	cd client && go install ${LDFLAGS} mkdb-client.go  `ls -1 *.go|grep -v mkdb-client.go`
server:
	cd server && go install ${LDFLAGS} mkdb-server.go `ls -1 *.go|grep -v mkdb-server.go`
test: client
	${GOBIN}/mkdb-client >test/thing.go
	cd test && go install ${LDFLAGS}  mkdb-test.go thing.go

