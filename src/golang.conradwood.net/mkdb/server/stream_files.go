package main

import (
	pb "golang.conradwood.net/apis/mkdb"
	"golang.conradwood.net/mkdb/creator"
)

func (e *echoServer) CreateDBFiles(req *pb.CreateDBRequest, srv pb.MKDB_CreateDBFilesServer) error {
	return creator.CreateDBFiles(req, srv)
}
