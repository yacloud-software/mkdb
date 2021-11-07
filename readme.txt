mkdb:
  create go files from a proto to update/select/insert protos in sql

There is also a presentation available here:
https://docs.google.com/presentation/d/1qHx1vUDHXdQ_CTQHeaJYxhDrFEoBtbegO_B_Kh_TyuU/edit#slide=id.p

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
  common.Language Language=4 [(common.sql_type)="uint32"]; // most commonly found language on this page
}

// one may create a foreign key:
message Foo {
 Bar BarObject=1 [(common.sql_type)="uint64",(common.sql_reference)="bar.id"];
}

// one may mark a column as unique
message Foo {
 Bar BarObject=1 [(common.sql_unique)="true"];
}



