package data

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"loginserver/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

type userCounterRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserCounterRepo(data *Data, logger log.Logger) biz.UserCounterRepo {
	return &userCounterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}


var USER_COUNTER = "USER_COUNTER"
var mu_user_counter sync.RWMutex

func (r *userCounterRepo) Create(ctx context.Context, account *biz.UserCounter) (*biz.UserCounter, error) {
	mu_user_counter.Lock()
	defer mu_user_counter.Unlock()
	id := r.data.rdb.Incr(USER_COUNTER).Val()
	return &biz.UserCounter{
		ID: id,
	}, nil
}

func (r *userCounterRepo) Get(ctx context.Context) (*biz.UserCounter, error) {
	mu_user_counter.RLock()
	defer mu_user_counter.RUnlock()
	id, err := strconv.ParseInt(r.data.rdb.Get(USER_COUNTER).Val(), 10, 64)
	return &biz.UserCounter{
		ID: id,
	}, err
}