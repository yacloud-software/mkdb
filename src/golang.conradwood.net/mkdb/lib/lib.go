package lib

import (
	"bytes"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/mkdb"
	"golang.conradwood.net/go-easyops/utils"
	"strings"
	"text/template"
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

func t_inc2(i int) int {
	return i + 2
}

func t_inc(i int) int {
	return i + 1
}
func t_deli(i int, s string) string {
	if i != 0 {
		return s
	}
	return ""
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
func (c *Creator) t_cols_no_id() []string {
	var res []string
	for _, f := range c.t_fields_no_id() {
		res = append(res, c.fieldcols[f])
	}
	return res
}

// return field names WITHOUT id
func (c *Creator) t_fields_no_id() []string {
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
func (c *Creator) t_fieldvalues_no_id() []string {
	var res []string
	for _, f := range c.Def.Fields {
		if f.PrimaryKey {
			continue
		}
		res = append(res, c.GetFieldValueName(f))
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
	return f.Name
}
func (c *Creator) t_id_col() string {
	return strings.ToLower(c.IDField)
}

// returns a go representation of an empty object matching the ID field's type.
// e.g. "" for string and 0 for int
func (c *Creator) t_id_null() string {
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

func (c *Creator) t_id_field() string {
	for _, f := range c.Def.Fields {
		if f.Type == 1 {
			return f.Name
		}
	}
	return c.IDField
}

// given a fieldname, will convert to column name
func (c *Creator) t_col_name(fieldname string) string {
	s, found := c.fieldcols[fieldname]
	if !found && *debug {
		for k, v := range c.fieldcols {
			fmt.Printf("\"%s\" -> \"%s\"\n", k, v)
		}
		panic(fmt.Sprintf("no column for field \"%s\"", fieldname))
	}
	return s
}

func (c *Creator) t_field_count() int {
	return len(c.Def.Fields)
}

func (c *Creator) colNameToField(colname string) *mkdb.ProtoField {
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

// given a column name will return extra create options
func (c *Creator) t_col_extraopts(name string) string {
	res := ""
	for _, f := range c.Def.Fields {
		if c.t_col_name(f.Name) != name {
			continue
		}
		found, ref, reff := c.GetOptSQLReference(f)
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
func (c *Creator) t_col_sqltype(name string) string {
	for _, f := range c.Def.Fields {
		if c.t_col_name(f.Name) == name {
			return to_sql_string(f.Type)
		}
		if c.t_col_name(c.GetFieldName(f)) == name {
			return to_sql_string(f.Type)
		}
	}
	if *debug {
		panic(fmt.Sprintf("No SQL for unknown field \"%s\"", name))
	}
	return fmt.Sprintf("UNKNOWN SQL (%s)", name)
}

// given a column name will return a valid sql default for this column, e.g. '' for string and 0 for int
func (c *Creator) t_col_sqldef(name string) string {
	for _, f := range c.Def.Fields {
		if c.t_col_name(f.Name) == name || c.t_col_name(c.GetFieldName(f)) == name {
			// if alter table default columns are off - insert here
			if f.Type == 5 {
				return `''`
			} else if f.Type == 6 {
				return `false`
			} else if f.Type == 2 || f.Type == 11 {
				return `0.0`
			} else {
				return `0`
			}
		}
	}
	if *debug {
		panic(fmt.Sprintf("No SQL definition for unknown field \"%s\"", name))
	}
	return fmt.Sprintf("UNKNOWN SQL (%s)", name)
}

// return the field "go type" (e.g. "int" or "uint64" or "double")
func (c *Creator) t_field_go_type(name string) string {
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
		panic(fmt.Sprintf("no type for unknown field \"%s\"", name))
	}
	return fmt.Sprintf("UNKNOWN[%s]", name)
}

func (c *Creator) CreateByDef(def *mkdb.ProtoDef) error {
	if c.IDField == "" {
		return fmt.Errorf("No IDField configured")
	}
	c.fieldcols = make(map[string]string)
	for _, f := range def.Fields {
		if f.Name == "" {
			return fmt.Errorf("Fields may not contain entries without a name")
		}
		fname := c.GetFieldName(f)
		l := strings.ToLower(fname)
		rpl, exists := replace_sql[l]
		if exists {
			l = rpl
		}
		c.fieldcols[fname] = l
	}

	c.Def = def
	t := template.New("foo")
	t.Funcs(template.FuncMap{
		"inc":               t_inc,
		"inc2":              t_inc2,
		"deli":              t_deli,
		"id_col":            c.t_id_col,
		"id_field":          c.t_id_field,
		"cols_no_id":        c.t_cols_no_id,
		"col_name":          c.t_col_name,
		"col_sqltype":       c.t_col_sqltype,
		"col_extraopts":     c.t_col_extraopts,
		"col_sqldef":        c.t_col_sqldef,
		"fields_no_id":      c.t_fields_no_id,
		"fieldvalues_no_id": c.t_fieldvalues_no_id,
		"field_count":       c.t_field_count,
		"field_gotype":      c.t_field_go_type,
		"id_null":           c.t_id_null,
		"msgInitializers":   c.msgInitializers,
	})
	foo, err := utils.ReadFile(*tmplfile)
	if err != nil {
		return err
	}
	templ, err := t.Parse(string(foo))
	if err != nil {
		return err
	}
	if c.t_id_col() == "" {
		return fmt.Errorf("Unable to determine or find ID Column. Configured: \"%s\"", c.IDField)
	}
	var w bytes.Buffer
	fmt.Printf("Primary ID: %s\n", c.t_id_col())
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
