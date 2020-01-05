package definitions

import "github.com/2HgO/grpc-gotag/server/config/db"
import js "github.com/golang/protobuf/jsonpb"

// User ...
var User = db.User

var marshaler = &js.Marshaler{OrigName: true}