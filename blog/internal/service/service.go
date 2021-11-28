package service

import (
	"context"
	"github.com/PGshen/go-project/blog/global"
	"github.com/PGshen/go-project/blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func NewService(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.NewDao(global.DBEngine)
	return svc
}
