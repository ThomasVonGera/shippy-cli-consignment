module github.com/ThomasVonGera/shippy-cli-consignment

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/ThomasVonGera/shippy-service-consignment v0.0.0-20200826152532-eae5c99f2db4
	github.com/micro/go-micro/v2 v2.9.1
)
