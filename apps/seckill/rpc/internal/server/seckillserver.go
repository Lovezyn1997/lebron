// Code generated by goctl. DO NOT EDIT!
// Source: seckkill.proto

package server

import (
	"context"

	"github.com/zhoushuguang/lebron/apps/seckill/rpc/internal/logic"
	"github.com/zhoushuguang/lebron/apps/seckill/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/seckill/rpc/seckill"
)

type SeckillServer struct {
	svcCtx *svc.ServiceContext
	seckill.UnimplementedSeckillServer
}

func NewSeckillServer(svcCtx *svc.ServiceContext) *SeckillServer {
	return &SeckillServer{
		svcCtx: svcCtx,
	}
}

func (s *SeckillServer) SeckillProducts(ctx context.Context, in *seckill.SeckillProductsRequest) (*seckill.SeckillProductsResponse, error) {
	l := logic.NewSeckillProductsLogic(ctx, s.svcCtx)
	return l.SeckillProducts(in)
}

func (s *SeckillServer) SeckillOrder(ctx context.Context, in *seckill.SeckillOrderRequest) (*seckill.SeckillOrderResponse, error) {
	l := logic.NewSeckillOrderLogic(ctx, s.svcCtx)
	return l.SeckillOrder(in)
}
