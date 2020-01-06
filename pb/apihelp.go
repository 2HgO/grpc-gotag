package pb

import (
	"time"
	"github.com/golang/protobuf/ptypes"
	pmongo "github.com/amsokol/mongo-go-driver-protobuf/pmongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	js "github.com/golang/protobuf/jsonpb"

)

// GetID ...
func (u *UserInfo) GetID() *primitive.ObjectID {
	if u.GetId().GetValue() == "" {
		return nil
	}
	id, err := u.GetId().GetObjectID()
	if err != nil {
		panic(err)
	}
	return &id
}

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

// MarshalJSON ...
func (u *UserInfo) MarshalJSON() ([]byte, error) {
	str, err := (&js.Marshaler{OrigName: true}).MarshalToString(u)
	return []byte(str), err
}

// UnmarshalJSON ...
func (u *UserInfo) UnmarshalJSON(b []byte) error {
	return js.UnmarshalString(string(b), u)
}