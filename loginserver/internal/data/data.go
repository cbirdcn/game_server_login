package data

import (
	//"context"
	"loginserver/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	// "github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAccountRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	rdb *redis.Client
	//mdb *mongo.Client
}

//// 数据库连接：NewData .
//func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
//	log := log.NewHelper(logger)
//	rdb := redis.NewClient(&redis.Options{
//		Addr:         c.Redis.Addr,
//		Password:     c.Redis.Password,
//		DB:           int(c.Redis.Db),
//		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
//		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
//		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
//	})
//	// rdb.AddHook(redisotel.TracingHook{})
//
//	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.Mongodb.Addr))
//	if err != nil {
//		panic(err)
//	}
//
//	d := &Data{
//		rdb: rdb,
//		mdb: mongoClient,
//	}
//
//	return d, func() {
//		log.Info("message", "closing the data resources")
//		if err := d.rdb.Close(); err != nil {
//			log.Error(err)
//		}
//	}, nil
//}

// 数据库连接：NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	// rdb.AddHook(redisotel.TracingHook{})

	d := &Data{
		rdb: rdb,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}