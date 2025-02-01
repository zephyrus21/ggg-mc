package account

import (
	"github.com/zephyrus21/ggg-mc/protos/account/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}
