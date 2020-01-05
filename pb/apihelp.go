package pb

import (
	"time"
	"github.com/golang/protobuf/ptypes"
	pmongo "github.com/amsokol/mongo-go-driver-protobuf/pmongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

// SetCreated ...
func (u *UserInfo) SetCreated(_created time.Time) {
	now, err := ptypes.TimestampProto(_created)
	if err != nil {
		u.CreatedAt = ptypes.TimestampNow()
		return
	}
	u.CreatedAt = now
}

// SetModified ...
func (u *UserInfo) SetModified(_modified time.Time) {
	now, err := ptypes.TimestampProto(_modified)
	if err != nil {
		u.ModifiedAt = ptypes.TimestampNow()
		return
	}
	u.ModifiedAt = now
}

// SetID ...
func (u *UserInfo) SetID(id interface{}) {
	u.Id = pmongo.NewObjectId(id.(primitive.ObjectID))
}