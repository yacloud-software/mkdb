package lib

import (
	"fmt"
	"golang.conradwood.net/apis/mkdb"
)

type create_nullable_fields struct {
	varname           string // e.g. scanTarget_4
	create            string // e.g. &savepb.Foo{}
	targetname        string // e.g. foo.Label (so we can do fmt.Sprintf("%s=%s",targetname,create)
	fieldintargetname string // e.g. "ID" (refers to targetname not create), so we can fmt.Sprintf("%s.%s=%s.Value()",targetname,fieldintargetname,varname)
}

// must create code to return "foo" and "err"
func (c *Creator) T_scanner() string {
	res := `foo := &savepb.` + c.Def.Name + "{}\n"
	for k, v := range c.ProtoCreateFields {
		fmt.Printf(" \"%s\" --> \"%s\"\n", k, v)
	}
	var cnf []*create_nullable_fields
	// create the non-nullable pointers
	res = res + "// create the non-nullable pointers\n"
	for _, pf := range c.Def.GetFields() {
		if !c.IsReference(pf) {
			continue
		}
		if c.IsNullable(pf) {
			continue
		}
		protoName := c.ProtoCreateFields[pf.Name]
		res = res + fmt.Sprintf("foo.%s = &%s{} // non-nullable\n", pf.Name, protoName)
	}
	res = res + "// create variables for scan results\n"
	maxcol := 0
	for i, pf := range c.Def.GetFields() {
		maxcol++
		//		cv := GetInternalType(pf.Type)
		fmt.Printf("FIELD: %v\n", pf)
		fn := c.GetFieldValueName(pf)
		b := c.IsNullable(pf)
		if b {
			// TODO: ID might not be uint64
			// TODO: ID might not be called "ID"
			protoName := c.ProtoCreateFields[pf.Name]
			cn := &create_nullable_fields{
				varname:           fmt.Sprintf("scanTarget_%d", i),
				targetname:        pf.Name, //fmt.Sprintf("foo.%s", fn),
				create:            "&" + protoName + "{}",
				fieldintargetname: "ID",
			}
			cnf = append(cnf, cn)
			res = res + fmt.Sprintf("   %s := &gosql.NullInt64{}  ", cn.varname)
			res = res + "// " + fn + "\n"

		} else {
			res = res + fmt.Sprintf("   scanTarget_%d := &foo.%s\n", i, fn)
		}
	}

	s := ""
	deli := ""
	for i := 0; i < maxcol; i++ {
		s = s + deli + fmt.Sprintf("scanTarget_%d", i)
		deli = ", "
	}
	res = res + "err := rows.Scan(" + s + ")\n"
	for _, cn := range cnf {
		fmt.Printf("CN: %#v\n", cn)
		res = res + fmt.Sprintf(" if %s.Valid {\n", cn.varname)
		res = res + fmt.Sprintf("   if foo.%s == nil {\n", cn.targetname)
		res = res + fmt.Sprintf("      foo.%s = %s\n", cn.targetname, cn.create)
		res = res + fmt.Sprintf("   }\n")
		res = res + fmt.Sprintf(`
       _,err:=%s.Value()
       if err !=nil {
           return nil,err
       }
`, cn.varname)
		res = res + fmt.Sprintf("   foo.%s.%s = uint64(%s.Int64)\n", cn.targetname, cn.fieldintargetname, cn.varname)
		res = res + "}\n"
	}
	return "// SCANNER:\n" + res + "// END SCANNER\n"
}
func (c *Creator) IsNullable(f *mkdb.ProtoField) bool {
	b, _, _ := c.GetOptSQLNullReference(f)
	if b {
		return true
	}
	return false
}
func (c *Creator) IsReference(f *mkdb.ProtoField) bool {
	b, _, _ := c.GetOptSQLReference(f)
	if b {
		return true
	}
	b, _, _ = c.GetOptSQLNullReference(f)
	if b {
		return true
	}
	return false
}


