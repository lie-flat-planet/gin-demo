package service

import (
	"context"
	goRedis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"github.com/lie-flat-planet/gin-demo/internal/model"
)

type Demo struct {
	db    *gorm.DB
	redis *goRedis.Client
}

func WithMysql(db *gorm.DB) func(demo *Demo) {
	return func(demo *Demo) {
		demo.db = db
	}
}

func WithRedis(redis *goRedis.Client) func(demo *Demo) {
	return func(demo *Demo) {
		demo.redis = redis
	}

}

func NewDemo(opts ...func(demo *Demo)) *Demo {
	user := &Demo{}

	for _, opt := range opts {
		opt(user)
	}

	return user
}

func (demo *Demo) FetchFirst(ctx context.Context) (*model.User, error) {
	m := &model.User{}

	err := demo.db.First(m).Error

	return m, err
}