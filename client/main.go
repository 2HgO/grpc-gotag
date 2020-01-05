package main

import (
	"net/http"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"github.com/2HgO/grpc-gotag/pb"
	js "github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/gin-gonic/gin"
)

type helper struct{}

func (h helper) Name() string {
	return "helper"
}

func (h helper) Bind(r *http.Request, out interface{}) error {
	return js.Unmarshal(r.Body, out.(proto.Message))
}

var bind = helper{}

var kacp = keepalive.ClientParameters{
	Time: 10 * time.Second,
	Timeout: time.Second,
	PermitWithoutStream: true,
}

var userClient = func() pb.APIClient {
	conn, err := grpc.Dial("localhost:55051", grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatalln(err)
	}
	return pb.NewAPIClient(conn)
}()

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		req := new(pb.UserInfo)
		if err := c.ShouldBindWith(req, bind); err != nil {
			c.Error(err)
			c.JSON(500, gin.H{"status": "not good"})
			return
		}
		res, err := userClient.NewUser(context.TODO(), req)
		if err != nil {
			c.Error(err)
			c.JSON(500, gin.H{"status": "not good"})
			return
		}

		c.Data(200, "application/json", res.GetData())
	})

	r.Run(":55041")
}