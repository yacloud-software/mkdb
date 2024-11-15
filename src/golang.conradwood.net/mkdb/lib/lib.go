package lib

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"strings"
	"text/template"

	"golang.conradwood.net/apis/mkdb"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
)

var (
	debug    = flag.Bool("debug_lib", false, "debug lib")
	tmplfile = flag.String("template_file", "configs/mkdb/godb-template", "template file for godb.go")
	// list psql reserved words...
	replace_sql = map[string]string{
		"value":     "r_value",
		"table":     "r_table",
		"where":     "r_where",
		"binary":    "r_binary",
		"public":    "r_public",
		"timestamp": "r_timestamp",
		"end":       "r_end",
		"as":        "r_as",
	}
)

type Creator struct {
	Pkgname           string
	Structname        string
	Dbgo              string
	TableName         string
	Def               *mkdb.ProtoDef
	IDField           string            // name of field to use as primary key
	fieldcols         map[string]string // fieldnames -> columnames
	EnumNames         []string
	ProtoCreateFields map[string]string
}

func NewCreator() *Creator {
	res := &Creator{
		TableName:  "echotable",
		Structname: "protoSave",
		Pkgname:    "main",
		IDField:    "id",
	}
	return res
}

func T_inc2(i int) int {
	return i + 2
}

func T_inc(i int) int {
	return i + 1
}
func T_deli(i int, s string) string {
	if i != 0 {
		return s
	}
	return ""
}

/*
create the foreignkeys
*/
func (c *Creator) create_foreign_keys() []string {
	var res []string
	for _, f := range c.Def.Fields {
		found, ref, reff := c.GetOptSQLReference(f)
		if !found {
			found, ref, reff = c.GetOptSQLNullReference(f)
			if !found {
				continue
			}
		}
		col_name := c.T_col_name(f.Name)
		constrainT_name := "mkdb_fk_" + hash(c.TableName+"_"+col_name+"_"+ref+reff)
		s := fmt.Sprintf("add constraint %s FOREIGN KEY (%s) references %s (%s) on delete cascade ", constrainT_name, col_name, ref, reff)

		if c.GetOptSQLUnique(f) {
			s = s + " unique "
		}
		res = append(res, s)
	}
	return res
}

/*
create the indices - TODO
*/
func (c *Creator) create_indices() []string {
	var res []string
	for _, f := range c.Def.Fields {
		found := c.GetOptSQLUnique(f)
		if !found {
			continue
		}
		col_name := c.T_col_name(f.Name)
		constrainT_name := "uniq_" + hash(c.TableName+"_"+col_name)
		s := fmt.Sprintf("create unique index if not exists %s on %s (%s)", constrainT_name, c.TableName, col_name)
		res = append(res, s)
		s = fmt.Sprintf("alter table %s add constraint %s unique using index %s", c.TableName, constrainT_name, constrainT_name)
		res = append(res, s)
	}
	return res

}
func hash(s string) string {
	if len(s) < 64 {
		//valid sqlname
		validchars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
		res := ""
		for _, c := range s {
			valid := false
			for _, cv := range validchars {
				if cv == c {
					valid = true
					break
				}
			}
			if !valid {
				c = '_'
			}
			res = res + string(c)
		}
		return res
	}
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

/*
returns a string suitable to include in a struct initialization
for example: "Foo: &savepb.Bar{}"
*/
func (c *Creator) msgInitializers() string {
	if c.ProtoCreateFields == nil {
		return ""
	}
	if len(c.ProtoCreateFields) == 0 {
		return ""
	}
	s := ""
	deli := ""
	for k, v := range c.ProtoCreateFields {
		s = s + deli + fmt.Sprintf("%s: &%s{}", k, v)
		deli = ", "
	}
	return s
}

// return column names WITHOUT id
func (c *Creator) T_cols_no_id() []string {
	var res []string
	for _, f := range c.T_fields_no_id() {
		res = append(res, c.fieldcols[f])
	}
	return res
}

// return field names WITHOUT id
func (c *Creator) T_fields_no_id() []string {
	var res []string
	for _, f := range c.Def.Fields {
		if f.PrimaryKey {
			continue
		}
		//	res = append(res, c.GetFieldValueName(f))
		res = append(res, c.GetFieldName(f))
	}
	return res
}

// return field names WITHOUT id
func (c *Creator) T_fieldvalues_no_id() []string {
	var res []string
	for _, f := range c.Def.Fields {
		if f.PrimaryKey {
			continue
		}
		res = append(res, c.GetFieldValueName(f))
	}
	return res
}

// return names of getters for all fields except the id
func (c *Creator) T_fieldvalue_getters_no_id() []string {
	var res []string
	for _, s := range c.T_fieldvalues_no_id() {
		sg := "get_" + s
		sg = strings.ReplaceAll(sg, ".", "_")
		res = append(res, sg)
	}

	return res
}

// return names of getters for all fields except the id
// TODO: honour NULLABLE tag
func (c *Creator) T_all_getters_code() []string {
	var res []string
	for _, f := range c.Def.Fields {
		s := c.GetFieldValueName(f)
		go_sql_typ := c.T_field_go_type(c.GetFieldValueName(f))
		comt := "// getter for field \"" + f.Name + "\" (" + s + ") [" + go_sql_typ + "] \n"
		is_reference := strings.Contains(s, ".")
		s = strings.ReplaceAll(s, ".", "_")
		ret := `return ` + go_sql_typ + ` (p.` + c.GetFieldValueName(f) + ")"

		if is_reference {
			comt = "// getter for reference \"" + f.Name + "\"\n"
			sx := c.GetFieldValueName(f)
			ref := strings.SplitN(sx, ".", 2)[0]
			if c.IsNullable(f) {
				// is a reference to another proto
				go_sql_typ = "gosql.NullInt64"
				ret = "if p." + ref + "==nil { return gosql.NullInt64{Valid:false} }\n"
				ret = ret + "return gosql.NullInt64{Valid:true,Int64:int64( p." + sx + " ) }"
			} else {
				ret = "if p." + ref + `==nil { panic("field ` + ref + ` must not be nil") }
`
				ret = ret + `return p.` + c.GetFieldValueName(f)
				go_sql_typ = "uint64"
			}
		}

		sg := comt + `func (a *` + c.Structname + `) get_` + s + `(p *savepb.` + c.Def.Name + `) ` + go_sql_typ + ` {
 ` + ret + `
}
`
		res = append(res, sg)
	}

	return res
}

// return the field name
func (c *Creator) GetFieldName(f *mkdb.ProtoField) string {
	return f.Name
}

// return the field name, if field is a reference, return the REFERENCED field
func (c *Creator) GetFieldValueName(f *mkdb.ProtoField) string {
	b, _, _ := c.GetOptSQLReference(f)
	if b {
		return fmt.Sprintf("%s.ID", f.Name)
	}
	b, _, _ = c.GetOptSQLNullReference(f)
	if b {
		return fmt.Sprintf("%s.ID", f.Name)
	}
	return f.Name
}
func (c *Creator) T_id_col() string {
	return strings.ToLower(c.IDField)
}

// returns a go representation of an empty object matching the ID field's type.
// e.g. "" for string and 0 for int
func (c *Creator) T_id_null() string {
	for _, f := range c.Def.Fields {
		if !f.PrimaryKey {
			continue
		}
		if f.Type == 5 {
			return "\"\""
		} else if f.Type == 4 {
			return "0"
		} else {
			return To_go_string(f.Type) + "{}"
		}

	}
	return ""

}

func (c *Creator) T_id_field() string {
	for _, f := range c.Def.Fields {
		if f.Type == 1 {
			return f.Name
		}
	}
	return c.IDField
}

// given a fieldname, will convert to column name
func (c *Creator) T_col_name(fieldname string) string {
	s, found := c.fieldcols[fieldname]
	if !found && *debug {
		for k, v := range c.fieldcols {
			fmt.Printf("\"%s\" -> \"%s\"\n", k, v)
		}
		panic(fmt.Sprintf("no column for field \"%s\"", fieldname))
	}
	return s
}

func (c *Creator) T_field_count() int {
	return len(c.Def.Fields)
}

func (c *Creator) T_ColNameToField(colname string) string {
	pf := c.ColNameToField(colname)
	if pf == nil {
		return "NOFIELDFORCOLUMN: \"" + colname + "\""
	}
	return pf.Name
}
func (c *Creator) T_ColNameToFieldGetter(colname string) string {
	pf := c.ColNameToField(colname)
	if pf == nil {
		return "NOFIELDFORCOLUMN: \"" + colname + "\""
	}
	return pf.Name
}

func (c *Creator) ColNameToField(colname string) *mkdb.ProtoField {
	for k, v := range c.fieldcols {
		if v != colname {
			continue
		}
		for _, f := range c.Def.Fields {
			if c.GetFieldName(f) == k {
				return f
			}
		}
		return nil
	}
	return nil
}

// returns empty string or the string "not null"
func (c *Creator) T_col_notnull(name string) string {
	for _, f := range c.Def.Fields {
		if c.T_col_name(f.Name) != name {
			continue
		}
		// if it has nullreference tag, then don't add "not null"
		found, _, _ := c.GetOptSQLNullReference(f)
		if found {
			return ""
		}
		break
	}

	return "not null"

}

// given a column name will return extra create options
func (c *Creator) T_col_extraopts(name string) string {
	res := ""
	for _, f := range c.Def.Fields {
		if c.T_col_name(f.Name) != name {
			continue
		}
		found, ref, reff := c.GetOptSQLReference(f)
		if found {
			res = res + fmt.Sprintf(" references %s (%s) on delete cascade ", ref, reff)
		}
		found, ref, reff = c.GetOptSQLNullReference(f)
		if found {
			res = res + fmt.Sprintf(" references %s (%s) on delete cascade ", ref, reff)
		}
		if c.GetOptSQLUnique(f) {
			res = res + " unique "
		}
		break // out of field-loop

	}

	return res
}

// given a column name will return sql type
func (c *Creator) T_col_sqltype(name string) string {
	for _, f := range c.Def.Fields {
		if c.T_col_name(f.Name) == name {
			return to_sql_string(f.Type)
		}
		if c.T_col_name(c.GetFieldName(f)) == name {
			return to_sql_string(f.Type)
		}
	}
	if *debug {
		panic(fmt.Sprintf("No SQL for unknown field \"%s\"", name))
	}
	return fmt.Sprintf("UNKNOWN SQL (%s)", name)
}

// given a column name will return a valid sql default for this column, e.g. ” for string and 0 for int
func (c *Creator) T_col_sqldef(name string) string {
	for _, f := range c.Def.Fields {
		if c.T_col_name(f.Name) == name || c.T_col_name(c.GetFieldName(f)) == name {
			// if alter table default columns are off - insert here
			if f.Type == 5 {
				return `''`
			} else if f.Type == 10 {
				return `''`
			} else if f.Type == 6 {
				return `false`
			} else if f.Type == 3 || f.Type == 4 || f.Type == 9 || f.Type == 7 || f.Type == 8 {
				return `0`
			} else if f.Type == 2 || f.Type == 11 {
				return `0.0`
			} else {
				return fmt.Sprintf("TODO_MKDB_SQLTYPE(%d)", f.Type)
			}
		}
	}
	if *debug {
		panic(fmt.Sprintf("No SQL definition for unknown field \"%s\"", name))
	}
	return fmt.Sprintf("UNKNOWN SQL (%s)", name)
}

// given the fieldname, return the field "go type" (e.g. "int" or "uint64" or "double")
func (c *Creator) T_field_go_type(name string) string {
	if name == "" {
		panic("Attempt to resolve name with zero length")
	}
	for _, f := range c.Def.Fields {
		if strings.ToLower(f.Name) == strings.ToLower(name) {
			return To_go_string(f.Type)
		}
		// try type if it is derived
		if strings.ToLower(c.GetFieldName(f)) == strings.ToLower(name) {
			return To_go_string(f.Type)
		}
	}
	/*
		if strings.ToLower(name) == c.IDField {
			return "uint64"
		}
	*/
	if *debug {
		fmt.Printf("no type for unknown field \"%s\"", name)
	}
	return fmt.Sprintf("UNKNOWN[%s]", name)
}

func (c *Creator) CreateByDef(def *mkdb.ProtoDef) error {
	if c.IDField == "" {
		return errors.Errorf("No IDField configured")
	}
	c.fieldcols = make(map[string]string)
	for _, f := range def.Fields {
		if f.Name == "" {
			return errors.Errorf("Fields may not contain entries without a name")
		}
		fname := c.GetFieldName(f)
		l := strings.ToLower(fname)
		rpl, exists := replace_sql[l]
		if exists {
			l = rpl
		}
		gotname, newname := c.GetOpt(f, "(common.sql_name)")
		if gotname {
			fmt.Printf("Name overriden. Field \"%s\" becomes \"%s\"\n", l, newname)
			l = newname
		}
		c.fieldcols[fname] = l
	}

	c.Def = def
	t := template.New("mkdb_lib_template")
	t.Funcs(template.FuncMap{
		"inc":                      T_inc,
		"inc2":                     T_inc2,
		"deli":                     T_deli,
		"id_col":                   c.T_id_col,
		"id_field":                 c.T_id_field,
		"cols_no_id":               c.T_cols_no_id,
		"col_name":                 c.T_col_name,
		"col2fieldname":            c.T_ColNameToField,
		"col2fieldgetter":          c.T_ColNameToFieldGetter,
		"col_notnull":              c.T_col_notnull,
		"col_sqltype":              c.T_col_sqltype,
		"col_extraopts":            c.T_col_extraopts,
		"col_sqldef":               c.T_col_sqldef,
		"fields_no_id":             c.T_fields_no_id,
		"scanner":                  c.T_scanner,
		"fieldvalues_no_id":        c.T_fieldvalues_no_id,
		"all_getters_code":         c.T_all_getters_code,
		"fieldvalue_getters_no_id": c.T_fieldvalue_getters_no_id,
		"field_count":              c.T_field_count,
		"field_gotype":             c.T_field_go_type,
		"id_null":                  c.T_id_null,
		"msgInitializers":          c.msgInitializers,
		"create_foreign_keys":      c.create_foreign_keys,
		"create_indices":           c.create_indices,
	})
	foo, err := utils.ReadFile(*tmplfile)
	if err != nil {
		return err
	}
	templ, err := t.Parse(string(foo))
	if err != nil {
		return err
	}
	if c.T_id_col() == "" {
		return errors.Errorf("Unable to determine or find ID Column. Configured: \"%s\"", c.IDField)
	}
	var w bytes.Buffer
	fmt.Printf("Primary ID: %s\n", c.T_id_col())
	err = templ.Execute(&w, c)
	if err != nil {
		return err
	}
	c.Dbgo = w.String()
	return nil
}

func (c *Creator) fn(name string) string {
	return "func (a *" + c.Structname + ") " + name + "() error {\n"
}

func (c *Creator) DBGo() string {
	return c.Dbgo
}

// map contains columnname->fieldname mapping
func (c *Creator) buildinsert(vals map[string]string) string {
	s := "db.insert into " + c.TableName + " ("
	return s
}
