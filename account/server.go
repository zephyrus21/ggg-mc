package account

import (
	"context"
	"fmt"
	"net"

	"github.com/zephyrus21/ggg-mc/protos/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
}

func ServeGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &grpcServer{s})
	reflection.Register(server)
	return server.Serve(lis)
}

func (s *grpcServer) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	acccount, err := s.service.CreateAccount(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountResponse{
		Account: &pb.Account{
			Id:   acccount.ID,
			Name: acccount.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	account, err := s.service.GetAccount(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:   account.ID,
			Name: account.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, req *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	accounts, err := s.service.GetAccounts(ctx, req.Skip, req.Take)
	if err != nil {
		return nil, err
	}

	resp := &pb.GetAccountsResponse{}
	for _, account := range accounts {
		resp.Accounts = append(resp.Accounts, &pb.Account{
			Id:   account.ID,
			Name: account.Name,
		})
	}

	return resp, nil
}
