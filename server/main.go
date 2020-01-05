package main

import (
	"log"
	"net"

	"github.com/2HgO/grpc-gotag/pb"
	def "github.com/2HgO/grpc-gotag/server/definitions"

	grpc "google.golang.org/grpc"

	// js "github.com/golang/protobuf/jsonpb"
)

func main() {
	lis, err := net.Listen("tcp", ":55051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	pb.RegisterAPIServer(s, new(def.Server))
	log.Println("server starting up")
	s.Serve(lis)
}
