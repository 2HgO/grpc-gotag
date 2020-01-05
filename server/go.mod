module github.com/2HgO/grpc-gotag/server

go 1.13

replace github.com/2HgO/grpc-gotag/pb => ../pb

require (
	github.com/2HgO/bongo v0.10.5
	github.com/2HgO/grpc-gotag/pb v0.0.0-00010101000000-000000000000
	github.com/DataDog/zstd v1.4.4 // indirect
	github.com/amsokol/mongo-go-driver-protobuf v1.0.0-rc5
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.1 // indirect
	github.com/mongodb/mongo-go-driver v1.2.0
	github.com/pkg/errors v0.8.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.2.0
	google.golang.org/grpc v1.26.0
)
