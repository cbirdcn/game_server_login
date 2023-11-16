package service

import (
	"context"
	"math/rand"
	"strconv"

	pb "loginserver/api/loginserver"
	"loginserver/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AccountService struct {
	pb.UnimplementedAccountServer
	uc *biz.AccountUsecase
	//log *log.Helper
}

func NewAccountService(uc *biz.AccountUsecase, logger log.Logger) *AccountService {
	return &AccountService{uc: uc}
}

func (s *AccountService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountReply, error) {
	_, err := s.uc.GetAccountByUsername(ctx, req.GetUsername())
	if err != nil {
		// 自增ID
		id := rand.Int()
		// 创建
		s.uc.CreateAccount(ctx, &biz.Account{
			ID: strconv.Itoa(id),
			UserName: req.GetUsername(),
		})
		return &pb.CreateAccountReply{
			Id: strconv.Itoa(id),
		}, nil
	}
	return &pb.CreateAccountReply{
		Id: "0",
	}, nil
}
func (s *AccountService) LoginAccount(ctx context.Context, req *pb.LoginAccountRequest) (*pb.LoginAccountReply, error) {
	return &pb.LoginAccountReply{}, nil
}
