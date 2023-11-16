package biz

import (
	"context"
	"time"

	loginserver "loginserver/api/loginserver"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)


// Account is a Account model.
// 使用Mongodb，直接用主键作为账号id
type UserCounter struct {
	ID           int64    `bson:"_id"`
}

type UserCounterRepo interface {
	Create(context.Context, *UserCounter) (*UserCounter, error)
	Get(context.Context) (*UserCounter, error)
}

type UserCounterUsecase struct {
	repo UserCounterRepo
	log  *log.Helper
}

func NewUserCounterUsecase(repo UserCounterRepo, logger log.Logger) *UserCounterUsecase {
	return &UserCounterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserCounterUsecase) Generate(ctx context.Context, user_counter *UserCounter) (*UserCounter, error) {
	return uc.repo.Create(ctx, user_counter)
}

func (uc *UserCounterUsecase) Get(ctx context.Context) (*UserCounter, error) {
	return uc.repo.Get(ctx)
}

