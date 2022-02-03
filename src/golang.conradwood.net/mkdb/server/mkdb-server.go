package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/emicklei/proto"
	pb "golang.conradwood.net/apis/mkdb"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/mkdb/lib"
	"golang.conradwood.net/mkdb/linux"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
	"strings"
)

var (
	port      = flag.Int("port", 10009, "The grpc server port")
	debug     = flag.Bool("debug", false, "debug mode")
	save_fail = flag.Bool("save_failed", false, "save go code which fails to compile (for analysis)")
)

type echoServer struct {
}

func main() {
	flag.Parse()
	sd := server.NewServerDef()
	sd.Port = *port
	sd.Register = server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterMKDBServer(server, e)
			return nil
		},
	)
	err := server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/
func (e *echoServer) GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessageResponse, error) {
	resp := &pb.GetMessageResponse{}
	ha, err := names(req.ProtoFile)
	if err != nil {
		return nil, err
	}
	for _, m := range ha.messages {
		am := &pb.AMessage{Name: m}
		resp.Messages = append(resp.Messages, am)
	}
	return resp, nil
}
func (e *echoServer) CreateDBFile(ctx context.Context, req *pb.CreateDBRequest) (*pb.CreateDBResponse, error) {
	if req.ProtoFileName == "" {
		return nil, errors.InvalidArgs(ctx, "missing proto filename", "missing proto filename")
	}
	if req.Message == "" {
		return nil, errors.InvalidArgs(ctx, "missing message name", "missing message name")
	}
	if req.ProtoFile == "" {
		return nil, errors.InvalidArgs(ctx, "missing content of protofile", "missing content of protofile")
	}
	pkg := "db"
	if req.Package != "" {
		pkg = req.Package
	}
	ids := req.IDField
	if ids == "" {
		ids = "ID"
	}
	ha, err := files(req, pkg, ids, req.TableName, req.TablePrefix)
	if err != nil {
		fmt.Printf("Failed to create files: %s\n", err)
		return nil, err
	}
	if ha == nil {
		panic("ha & err is nil")
	}
	resp := &pb.CreateDBResponse{}
	resp.GoFile = ha.gofiles[req.Message]
	if resp.GoFile == "" {
		fmt.Printf("created %d go files\n", len(ha.gofiles))
		for k, _ := range ha.gofiles {
			fmt.Printf("GOFile: %s\n", k)
		}
		return nil, errors.InvalidArgs(ctx, "failed to convert proto", "failed to convert protobuf \"%s\" to go", req.Message)
	}
	return resp, nil
}

type Handlers struct {
	pkg               string
	messages          []string
	msg               string
	err               error
	apipkg            string
	enums             []string
	idfield           string
	gofiles           map[string]string // one gofile per message
	tablename         string
	tableprefix       string
	protoCreateFields map[string]string // fieldname -> package.message (for from-rows create like "Foo: &savepb.Bar{}")
}

/************************************************************************
* create go files from protos
************************************************************************/

// create .go files
func files(req *pb.CreateDBRequest, pkg string, idfield string, tname, tprefix string) (*Handlers, error) {
	msgname := req.Message
	content := req.ProtoFile
	// find out which package the protofile lives in..
	apipkg := ""
	for _, l := range strings.Split(content, "\n") {
		if !strings.Contains(l, "package") {
			continue
		}
		apipkg = strings.TrimSuffix(strings.TrimPrefix(l, "package "), ";")
		break
	}
	if apipkg == "" {
		return nil, fmt.Errorf("Unable to find package in protofile content")
	}
	fmt.Printf("APIPackage: \"%s\" (%s)\n", apipkg, msgname)
	reader := strings.NewReader(content)
	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %s", err)
	}
	importpath := req.ImportPath
	if importpath == "" {
		// derive from protofile
		sx := strings.Split(req.ProtoFileName, "/")
		if len(sx) < 5 {
			return nil, fmt.Errorf("No importpath specified and length of protofile path is too short (%d parts) (path=%s)", len(sx), req.ProtoFileName)
		}
		importpath = fmt.Sprintf("%s/%s", sx[len(sx)-4], sx[len(sx)-3])
	}
	fmt.Printf("ImportPath: %s\n", importpath)
	ha := &Handlers{
		apipkg:            fmt.Sprintf("%s/%s", importpath, apipkg),
		pkg:               pkg,
		msg:               msgname,
		gofiles:           make(map[string]string),
		idfield:           idfield,
		tablename:         tname,
		tableprefix:       tprefix,
		protoCreateFields: make(map[string]string),
	}
	// resolve first:
	proto.Walk(definition,
		proto.WithEnum(ha.findEnum),
		proto.WithService(ha.handleService),
		proto.WithMessage(ha.handleMessageNameOnly))

	//create
	proto.Walk(definition,
		proto.WithService(ha.handleService),
		proto.WithMessage(ha.handleMessage))
	if ha.err != nil {
		return nil, ha.err
	}
	return ha, nil
}

func (h *Handlers) findEnum(s *proto.Enum) {
	fmt.Printf("Enum: %s\n", s.Name)
	h.enums = append(h.enums, s.Name)
}

func (h *Handlers) handleService(s *proto.Service) {
	//	fmt.Println(s.Name)
}
func (h *Handlers) handleMessageNameOnly(m *proto.Message) {
	fmt.Printf("Message: %s\n", m.Name)
	h.messages = append(h.messages, m.Name)
}
func (h *Handlers) handleMessage(m *proto.Message) {
	def := pb.ProtoDef{
		ImportPath: h.apipkg,
		Name:       m.Name,
	}
	//fmt.Printf("%s\n", m.Name)
	h.messages = append(h.messages, m.Name)
	if m.Name != h.msg {
		return
	}
	idf := h.idfield
	if idf == "" {
		idf = "id"
	}
	for _, e := range m.Elements {
		x, ok := e.(*proto.NormalField)
		if !ok {
			continue
		}
		if x.Name == "" {
			err := fmt.Errorf("weirdo field without name in line %d\n", x.Position.Line)
			fmt.Printf("Error: %s\n", err)
			h.err = err
			return
		}
		tn := lib.From_proto_string(x.Type)
		// find nested messages
		if tn == 0 {
			if strings.Contains(x.Type, ".") {
				fmt.Printf("Field %s Contains dot, skipping initalisation (because we cannot add imports yet)\n", x.Name)
			} else {
				if !h.IsEnum(x.Type) && !x.Repeated {
					add := fmt.Sprintf("savepb.%s", x.Type)
					fmt.Printf("Adding: %s -> \"%s\"\n", x.Name, add)
					h.protoCreateFields[x.Name] = add
				}
			}
		}
		opts := make(map[string]string)
		ignore := false
		for _, o := range x.Options {
			opts[o.Name] = o.Constant.Source
			if o.Name == "(common.sql_type)" {
				fmt.Printf("SQL_Type for \"%s\" overriden to \"%s\"\n", x.Name, o.Constant.Source)
				x.Type = o.Constant.Source
			} else if o.Name == "(common.sql_ignore)" {
				ignore = true
			} else if o.Name == "(common.sql_reference)" {
				fmt.Printf("SQL_Reference for \"%s\": \"%s\"\n", x.Name, o.Constant.Source)
			} else {
				fmt.Printf("Option for %s: %#v\n", x.Name, o)
			}
		}
		if ignore {
			continue
		}
		if strings.HasPrefix(strings.ToLower(x.Name), "obsolete_") {
			continue
		}
		if x.Repeated {
			fmt.Printf("cannot yet handle repeated fields for %s %s\n", x.Type, x.Name)
			continue
		}
		t := lib.From_proto_string(x.Type)
		pk := false
		if strings.ToLower(x.Name) == strings.ToLower(idf) {
			pk = true // primary key
		}

		if t == 0 {
			if h.IsEnum(x.Type) {
				t = lib.From_proto_string("enum")
			}
			if t == 0 {
				if x.Repeated {
					fmt.Printf("cannot yet handle repeated fields for %s %s\n", x.Type, x.Name)
					continue
				}
				/* it is a reference to another message - how do we handle it?
				* we can
				* 1) abort with error
				* 2) ignore it
				* 3) reference resolve via ID
				* services we know use it: "gitserver", "document" ...
				 */
				refaction := 2
				fmt.Printf("***** WARNING. Object references are not yet resolved (action=%d, field=\"%s\", comment=\"%s\")\n", refaction, x.Name, comment(x))
				if refaction == 0 {
					// abort with error
					h.err = fmt.Errorf("   unknown type %s for %s\n", x.Type, x.Name)
					break
				} else if refaction == 2 {
					// ignore it
					continue
				} else if refaction == 3 {
					// reference resolve:
					t = lib.From_proto_string("uint64")
				}
			}
		}
		pf := &pb.ProtoField{Options: opts, Name: x.Name, Type: t, PrimaryKey: pk}
		def.Fields = append(def.Fields, pf) // add the name to the list

		//		fmt.Printf("   %s %s (%d)\n", x.Name, x.Type, t)
	}
	if *debug {
		fmt.Printf("   %d fields\n", len(def.Fields))
	}
	creator := lib.NewCreator()
	creator.IDField = h.idfield
	creator.Pkgname = h.pkg
	creator.Structname = fmt.Sprintf("DB%s", m.Name)
	if h.tablename != "" {
		creator.TableName = h.tablename
	} else {
		creator.TableName = strings.ToLower(m.Name)
	}
	creator.TableName = h.tableprefix + creator.TableName
	creator.EnumNames = h.enums
	creator.ProtoCreateFields = h.protoCreateFields
	err := creator.CreateByDef(&def)
	if err != nil {
		h.err = fmt.Errorf("failed to create by def: %s", err)
		return
	}
	gof := creator.DBGo()
	res, err := linux.SafelyExecute([]string{"/opt/yacloud/ctools/dev/go/current/go/bin/gofmt"}, strings.NewReader(gof))
	if err != nil {
		fmt.Printf("gofmt failed:\n%s\n", gof)
		if *save_fail {
			ioutil.WriteFile("/tmp/failed.go", []byte(gof), 0660)
		}
		h.err = fmt.Errorf("Gofmt for %s failed: %s", m.Name, err)
		fmt.Println(res)
		return
	}
	h.gofiles[m.Name] = res

}

/************************************************************************
* get proto names
************************************************************************/
// get the names...
func names(content string) (*Handlers, error) {
	reader := strings.NewReader(content)
	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %s", err)
	}
	ha := &Handlers{
		pkg:     "foobarpackage",
		gofiles: make(map[string]string),
	}
	proto.Walk(definition,
		proto.WithService(ha.handleService),
		proto.WithMessage(ha.handleMessageNameOnly))
	if ha.err != nil {
		return nil, ha.err
	}
	return ha, nil
}

func (h *Handlers) IsEnum(name string) bool {
	for _, n := range h.enums {
		if n == name {
			return true
		}
	}
	return false
}

func comment(field *proto.NormalField) string {
	c := ""
	if field.Comment != nil {
		c = c + strings.Join(field.Comment.Lines, "\n")
	}
	if field.InlineComment != nil {
		c = c + strings.Join(field.InlineComment.Lines, "\n")
	}
	return ""
}
