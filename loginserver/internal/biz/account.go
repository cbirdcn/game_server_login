package biz

import (
	"context"
	"time"

	loginserver "loginserver/api/loginserver"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrExist = errors.NotFound(loginserver.ErrorReason_EXIST.String(), "account exists")
)

// Account is a Account model.
// 使用Mongodb，直接用主键作为账号id
type Account struct {
	ID           string    `bson:"_id"`
	UserName     string    `bson:"username"`
	PasswordHash string    `bson:"password_hash"`
	IsBanned     string    `bson:"is_banned"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

// AccountRepo is a Account repo.
type AccountRepo interface {
	Create(context.Context, *Account) (*Account, error)
	Save(context.Context, *Account) (*Account, error)
	Update(context.Context, *Account) (*Account, error)
	FindByID(context.Context, string) (*Account, error) // 通过账号id(字符串)查ObjectId主键对应的记录
	FindByUsername(context.Context, string) (*Account, error)
	// ListByHello(context.Context, string) ([]*Account, error)
	// ListAll(context.Context) ([]*Account, error)

}

// AccountUsecase is a Account usecase.
type AccountUsecase struct {
	repo AccountRepo
	log  *log.Helper
}

// NewAccountUsecase new a Account usecase.
func NewAccountUsecase(repo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateAccount creates a Account, and returns the new Account.
func (uc *AccountUsecase) CreateAccount(ctx context.Context, account *Account) (*Account, error) {
	return uc.repo.Create(ctx, account)
}

// 通过账号名获取账号信息
func (uc *AccountUsecase) GetAccountByUsername(ctx context.Context, username string) (*Account, error) {
	return uc.repo.FindByUsername(ctx, username)
}

// 更新账号信息
func (uc *AccountUsecase) Update(ctx context.Context, account *Account) (*Account, error) {
	return uc.repo.Update(ctx, account)
}
