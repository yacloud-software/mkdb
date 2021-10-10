package lib

import (
	"fmt"
)

var (
	go_conv map[int32]conv
)

type conv struct {
	Go    string
	Sql   string
	Proto string
}

func init() {
	go_conv = make(map[int32]conv)
	// 1 is special type "id"
	// TODO: add_fields here
	//	go_conv[1] = conv{Go: "uint64", Sql: "integer"}
	go_conv[2] = conv{Go: "float64", Sql: "double precision"}
	go_conv[3] = conv{Go: "uint32", Sql: "integer"}
	go_conv[4] = conv{Go: "uint64", Sql: "bigint"}
	//	go_conv[5] = conv{Go: "string", Sql: "varchar(2000)"}
	go_conv[5] = conv{Go: "string", Sql: "text"}
	go_conv[6] = conv{Go: "bool", Sql: "boolean"}
	go_conv[7] = conv{Go: "int32", Sql: "integer"}
	go_conv[8] = conv{Go: "int64", Sql: "integer"}
	go_conv[9] = conv{Go: "uint32", Sql: "integer"} // enums
	go_conv[10] = conv{Go: "[]byte", Sql: "bytea", Proto: "bytes"}
	go_conv[11] = conv{Proto: "double", Go: "float64", Sql: "numeric(18,4)"}
}

func From_go_string(t string) int32 {
	if t == "enum" {
		return 9
	}
	for k, v := range go_conv {
		if v.Go == t {
			return k
		}
	}
	fmt.Printf("Weird string: \"%s\"\n", t)
	return 0
}
func From_proto_string(t string) int32 {
	if t == "enum" {
		return 9
	}
	for k, v := range go_conv {
		if v.Proto == t {
			return k
		}
	}
	for k, v := range go_conv {
		if v.Go == t {
			return k
		}
	}
	fmt.Printf("Weird string: \"%s\"\n", t)
	return 0
}

func To_go_string(i int32) string {
	f, ok := go_conv[i]
	if ok {
		return f.Go
	}

	if i == 1 {
		return "uint64"
	}

	return fmt.Sprintf("Type %d undefined", i)
}

func to_sql_string(i int32) string {
	f, ok := go_conv[i]
	if ok {
		return f.Sql
	}
	if i == 1 {
		return "integer"
	}
	return fmt.Sprintf("sql-Type %d undefined", i)

}
