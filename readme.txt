mkdb:
  create go files from a proto to update/select/insert protos in sql

Example usage:

=== list all the protobufs in a protofile:

 mkdb-client -protofile protos/golang.conradwood.net/apis/logservice/logservice.proto -messages
Name:"LogAppDef" 
Name:"LogLine" 
Name:"LogRequest" 
Name:"LogResponse" 
Name:"LogFilter" 
Name:"GetLogRequest" 
Name:"LogEntry" 
Name:"GetLogResponse" 
Name:"GetHostLogResponse" 
Name:"GetHostLogRequest" 
Name:"GetAppsResponse" 
Name:"EmptyRequest" 
Name:"CloseLogRequest" 

=== create a go file from a protobuf:
mkdb-client -protofile protos/golang.conradwood.netprotos/golang.conradwood.net//apis/logservice/logservice.proto -create -protobuf=LogLine -out=/tmp/foo.go

the contents of "/tmp/foo.go" now contain a compilable .go file with database Accessors:

grep func /tmp/foo.go:

Archive Table: (structs can be moved from main to archive using Archive() function)
func NewDBLogLine(db *sql.DB) *DBLogLine {
func (a *DBLogLine) Archive(ctx context.Context, id uint64) error {
func (a *DBLogLine) Save(ctx context.Context, p *savepb.LogLine) (uint64, error) {
func (a *DBLogLine) Update(ctx context.Context, p *savepb.LogLine) error {
func (a *DBLogLine) DeleteByid(ctx context.Context, p uint64) error {
func (a *DBLogLine) Byid(ctx context.Context, p uint64) (*savepb.LogLine, error) {
func (a *DBLogLine) ByTime(ctx context.Context, p int64) ([]*savepb.LogLine, error) {
func (a *DBLogLine) ByLine(ctx context.Context, p string) ([]*savepb.LogLine, error) {
func (a *DBLogLine) ByLevel(ctx context.Context, p int32) ([]*savepb.LogLine, error) {
func (a *DBLogLine) ByStatus(ctx context.Context, p string) ([]*savepb.LogLine, error) {
func (a *DBLogLine) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.LogLine, error) {

typical Usage:

dblogline := NewDBLogLine(sql.Open())
id := dblogline.Save(fooproto)
fooproto := dblogline.Byid(id)

========== annotations ========


// one may override the type:

message Page {
  uint64 ID=1;
  Image Image=2 [(common.sql_type)="uint64"]; // may be nul if we did not ask for it
  string TextContent=3; // text on this page
  common.Language Language=4 [(common.sql_type)="uint32"]; // enum: most commonly found language on this page
}

// one may create a foreign key:
message Foo {
 Bar BarObject=1 [(common.sql_type)="uint64",(common.sql_reference)="bar.id"];
}

// one may mark a column as unique
message Foo {
 Bar BarObject=1 [(common.sql_unique)="true"];
}

// one may mark a field as irrelevant for sql (e.g. one that is derived from other fields)
message Foo {
 Bar BarObject=1 [(common.sql_ignore)="true"];
}

// one may change the sql name of a field
message Foo {
 Bar BarObject=1 [(common.sql_name)="foocolumnname"];
}

// a field may be null (especially useful for embedded protos)
message Foo {
 Bar BarObject=1 [(common.sql_nullable)="true",(common.sql_type)="uint64"]; // may be nul if we did not ask for it
}





============= custom handlers ========
these apply filters to all query and store operations
type s struct {
}
// return map of column names and values to query for
func (s *s) FieldsToQuery(ctx context.Context) (map[string]any, error) {
	return map[string]any{"active": true}, nil // this will return only rows where active == true
}
// return map of column names and values to store in database in addition to the proto
func (s *s) FieldsToStore(ctx context.Context, i interface{}) (map[string]any, error) {
	return map[string]any{"saved": time.Now().Unix()}, nil // this will save current timestamp to column saved
}
