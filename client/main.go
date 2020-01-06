package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/keepalive"
	"github.com/2HgO/grpc-gotag/pb"
	"github.com/amsokol/mongo-go-driver-protobuf/pmongo"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gin-gonic/gin"
)

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
		if err := c.ShouldBindJSON(req); err != nil {
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

	r.GET("/:id", func(c *gin.Context) {
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.Error(err)
			c.JSON(500, gin.H{"status": "not good"})
			return
		}
		req := &pb.UserID{Id: pmongo.NewObjectId(id)}
		res, err := userClient.GetUser(context.TODO(), req)
		if err != nil {
			c.Error(err)
			c.JSON(500, gin.H{"status": "not good"})
			return
		}
		c.Data(200, "application/json", res.GetData())
	})

	r.GET("/", func(c *gin.Context) {
		res, err := userClient.GetAllUsers(context.TODO(), &empty.Empty{})
		if err != nil {
			c.Error(err)
			c.JSON(500, gin.H{"status": "not good"})
			return
		}
		c.Data(200, "application/json", res.GetData())
	})

	r.Run(":55041")
}