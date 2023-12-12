package main

import (
	"fmt"
	pb "golang.conradwood.net/apis/mkdb"
	//	"golang.conradwood.net/mkdb/creator"
	"strings"
)

func (e *echoServer) CreateDBFiles(req *pb.CreateDBRequest, srv pb.MKDB_CreateDBFilesServer) error {
	ctx := srv.Context()
	res, err := e.CreateDBFile(ctx, req)
	if err != nil {
		return err
	}
	for _, s := range strings.Split(res.GoFile, "\n") {
		fname := fmt.Sprintf("db-%s.go", req.Message)
		s = s + "\n"
		sd := &pb.FileStream{Filename: fname, Data: []byte(s)}
		err = srv.Send(sd)
		if err != nil {
			return err
		}
	}
	return nil
}



