package definitions

import "context"
import "github.com/2HgO/grpc-gotag/server/config/db"
import empty "github.com/golang/protobuf/ptypes/empty"
import "github.com/2HgO/grpc-gotag/pb"

import "encoding/json"

var user = db.User

// Server ...
type Server struct {}

// GetUser ...
func (s Server) GetUser(ctx context.Context, in *pb.UserID) (*pb.Res, error) {
	user := new(pb.UserInfo)
	if err := User.FindByID(in.GetId(), user); err != nil {
		return nil, err
	}
	data, _ := json.Marshal(user)
	return &pb.Res{
		Success: true,
		Data: data,
	}, nil
}

// NewUser ...
func (s Server) NewUser(ctx context.Context, in *pb.UserInfo) (*pb.Res, error) {
	if err := User.Insert(in); err != nil {
		return nil, err
	}
	data, _ := json.Marshal(in)
	return &pb.Res{
		Success: true,
		Data: data,
	}, nil
}

// GetAllUsers ...
func (s Server) GetAllUsers(ctx context.Context, in *empty.Empty) (*pb.Res, error) {
	users := []*pb.UserInfo{}
	iter, err := User.Find(map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	if err = iter.All(context.Background(), &users); err != nil {
		return nil, err
	}
	data, _ := json.Marshal(users)

	return &pb.Res{
		Success: true,
		Data: []byte(data),
	}, nil
}