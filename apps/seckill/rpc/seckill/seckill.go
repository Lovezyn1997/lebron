// Code generated by goctl. DO NOT EDIT!
// Source: seckkill.proto

package seckill

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Seckill interface {
		SeckillProducts(ctx context.Context, in *SeckillProductsRequest, opts ...grpc.CallOption) (*SeckillProductsResponse, error)
		SeckillOrder(ctx context.Context, in *SeckillOrderRequest, opts ...grpc.CallOption) (*SeckillOrderResponse, error)
	}

	defaultSeckill struct {
		cli zrpc.Client
	}
)

func NewSeckill(cli zrpc.Client) Seckill {
	return &defaultSeckill{
		cli: cli,
	}
}

func (m *defaultSeckill) SeckillProducts(ctx context.Context, in *SeckillProductsRequest, opts ...grpc.CallOption) (*SeckillProductsResponse, error) {
	client := NewSeckillClient(m.cli.Conn())
	return client.SeckillProducts(ctx, in, opts...)
}

func (m *defaultSeckill) SeckillOrder(ctx context.Context, in *SeckillOrderRequest, opts ...grpc.CallOption) (*SeckillOrderResponse, error) {
	client := NewSeckillClient(m.cli.Conn())
	return client.SeckillOrder(ctx, in, opts...)
}
