// client create: MKDBClient
/*
  Created by /srv/home/cnw/devel/go/go-tools/src/golang.conradwood.net/gotools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.conradwood.net/apis/mkdb/mkdb.proto
   gopackage : golang.conradwood.net/apis/mkdb
   importname: ai_0
   clientfunc: GetMKDB
   serverfunc: NewMKDB
   lookupfunc: MKDBLookupID
   varname   : client_MKDBClient_0
   clientname: MKDBClient
   servername: MKDBServer
   gscvname  : mkdb.MKDB
   lockname  : lock_MKDBClient_0
   activename: active_MKDBClient_0
*/

package mkdb

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_MKDBClient_0 sync.Mutex
  client_MKDBClient_0 MKDBClient
)

func GetMKDBClient() MKDBClient { 
    if client_MKDBClient_0 != nil {
        return client_MKDBClient_0
    }

    lock_MKDBClient_0.Lock() 
    if client_MKDBClient_0 != nil {
       lock_MKDBClient_0.Unlock()
       return client_MKDBClient_0
    }

    client_MKDBClient_0 = NewMKDBClient(client.Connect(MKDBLookupID()))
    lock_MKDBClient_0.Unlock()
    return client_MKDBClient_0
}

func MKDBLookupID() string { return "mkdb.MKDB" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.
