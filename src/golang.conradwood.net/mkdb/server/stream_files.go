package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	pb "golang.conradwood.net/apis/mkdb"

	//	"golang.conradwood.net/mkdb/creator"
	"strings"

	"golang.conradwood.net/go-easyops/utils"
)

func (e *echoServer) CreateDBFiles(req *pb.CreateDBRequest, srv pb.MKDB_CreateDBFilesServer) error {
	ctx := srv.Context()
	extra_files := make(map[string][]byte)

	filename, err := utils.FindFile("extra/stdfiles/dbpkg.go")
	if err != nil {
		return err
	}
	dir := filepath.Dir(filename)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		fname := f.Name()
		if filepath.Ext(fname) != ".go" {
			continue
		}
		ffname := dir + "/" + fname
		fmt.Printf("File: \"%s\" (%s)\n", fname, ffname)
		ct, err := utils.ReadFile(ffname)
		if err != nil {
			return err
		}
		extra_files[fname] = ct
	}

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
	for k, v := range extra_files {
		sd := &pb.FileStream{Filename: k, Data: v}
		err = srv.Send(sd)
		if err != nil {
			return err
		}
	}
	return nil
}
