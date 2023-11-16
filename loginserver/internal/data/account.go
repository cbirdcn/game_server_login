package data

import (
	"context"
	"errors"
	"fmt"

	"loginserver/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

// NewAccountRepo .
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}


var NICK_NAME_TO_ID_KEY = "NICK_NAME_TO_ID"

func (r *accountRepo) Create(ctx context.Context, account *biz.Account) (*biz.Account, error) {
	r.data.rdb.HSet(NICK_NAME_TO_ID_KEY, account.UserName, account.ID)
	return account, nil
}

func (r *accountRepo) Save(ctx context.Context, account *biz.Account) (*biz.Account, error) {
	return account, nil
}

func (r *accountRepo) Update(ctx context.Context, account *biz.Account) (*biz.Account, error) {
	return account, nil
}

func (r *accountRepo) FindByID(ctx context.Context, id string) (*biz.Account, error) {
	return nil, nil
}

func (r *accountRepo) FindByUsername(ctx context.Context, username string) (*biz.Account, error) {
	// NICK_NAME_TO_ID_KEY := "NICK_NAME_TO_ID"
	if r.data.rdb.HExists(NICK_NAME_TO_ID_KEY, username).Val() {
		userId := r.data.rdb.HGet(NICK_NAME_TO_ID_KEY, username).Val()
		return &biz.Account{
				ID: fmt.Sprintf("%#v",userId),
				UserName: fmt.Sprintf("%#v",username),
				// PasswordHash: fmt.Sprintf("%#v",result["password_hash"]),
				// IsBanned: fmt.Sprintf("%#v",result["is_banned"]),
				// CreatedAt: time.Time,
				// UpdatedAt: fmt.Sprintf("%#v",result["updated_at"]),
			}, nil
	}
	return nil, errors.New("NICK_NAME not exist")
	// fmt.Println(10)
	// coll := r.data.mdb.Database("global").Collection("user_account")
	// var result bson.M
	// err := coll.FindOne(ctx, bson.D{{"device_name", username}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	return nil, err
	// }

	// return &biz.Account{
	// 	ID: fmt.Sprintf("%#v",result["uid"]),
	// 	UserName: fmt.Sprintf("%#v",result["username"]),
	// 	PasswordHash: fmt.Sprintf("%#v",result["password_hash"]),
	// 	IsBanned: fmt.Sprintf("%#v",result["is_banned"]),
	// 	// CreatedAt: time.Time,
	// 	// UpdatedAt: fmt.Sprintf("%#v",result["updated_at"]),
	// }, nil
}

// func (r *accountRepo) ListByHello(context.Context, string) ([]*biz.Account, error) {
// 	return nil, nil
// }

// func (r *accountRepo) ListAll(context.Context) ([]*biz.Account, error) {
// 	return nil, nil
// }
