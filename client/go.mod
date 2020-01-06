module github.com/2HgO/grpc-gotag/client

go 1.13

replace github.com/2HgO/grpc-gotag/pb => ../pb

require (
	github.com/2HgO/grpc-gotag/pb v0.0.0-00010101000000-000000000000
	github.com/amsokol/mongo-go-driver-protobuf v1.0.0-rc5
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/smartystreets/goconvey v1.6.4 // indirect
	go.mongodb.org/mongo-driver v1.2.0
	google.golang.org/grpc v1.26.0
)
