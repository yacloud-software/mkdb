package lib

import (
	"fmt"
	"golang.conradwood.net/apis/mkdb"
	"strings"
)

// get a value of an option (true if it exists)
func (c *Creator) GetOpt(field *mkdb.ProtoField, name string) (bool, string) {
	s, ok := field.Options[name]
	return ok, s
}
func (c *Creator) GetOptSQLUnique(field *mkdb.ProtoField) bool {
	found, _ := c.GetOpt(field, "(common.sql_unique)")
	return found

}
func (c *Creator) HasIgnoreOption(field *mkdb.ProtoField) bool {
	found, _ := c.GetOpt(field, "(common.sql_ignore)")
	return found
}

// of sql_reference: get key and value
func (c *Creator) GetOptSQLReference(field *mkdb.ProtoField) (bool, string, string) {
	found, s := c.GetOpt(field, "(common.sql_reference)")
	if !found {
		return false, "", ""
	}
	kv := strings.Split(s, ".")
	if len(kv) != 2 {
		return true, fmt.Sprintf("INVALID REFERENCE: \"%s\"", s), fmt.Sprintf("INVALID REFERENCE: \"%s\"", s)
	}
	return true, kv[0], kv[1]
}

// of sql_reference: get key and value
func (c *Creator) GetOptSQLNullReference(field *mkdb.ProtoField) (bool, string, string) {
	found, s := c.GetOpt(field, "(common.sql_nullreference)")
	if !found {
		return false, "", ""
	}
	kv := strings.Split(s, ".")
	if len(kv) != 2 {
		return true, fmt.Sprintf("INVALID REFERENCE: \"%s\"", s), fmt.Sprintf("INVALID REFERENCE: \"%s\"", s)
	}
	return true, kv[0], kv[1]
}


