package client

import (
	"blockchain-middleware-sdk/config"
	rpc_model "blockchain-middleware-sdk/model/rpc-model"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type MiddleWareClient struct {
	Conn rpc_model.RpcServerClient
}

func NewClient() (*MiddleWareClient, error) {
	var client MiddleWareClient
	conn, err := grpc.Dial(config.ClientConfiguration.ServerHost+":"+config.ClientConfiguration.ServerPort, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("connect rpc server failed! err: %s", err.Error())
		return nil, err
	}
	rpcServerClient := rpc_model.NewRpcServerClient(conn)
	client.Conn = rpcServerClient
	return &client, nil
}

func (c *MiddleWareClient) CallHandleInitBill(ctx context.Context, args *rpc_model.Bill) (*rpc_model.RpcResponse, error) {
	data, err := c.Conn.HandleInitBill(ctx, args)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *MiddleWareClient) CallHandleQueryBills(ctx context.Context, args *rpc_model.QueryBillWithPage) (*rpc_model.RpcResponseWithPageMeta, error) {
	data, err := c.Conn.HandleQueryBills(ctx, args)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *MiddleWareClient) CallHandleDeleteBill(ctx context.Context, args interface{}) (*rpc_model.RpcResponse, error) {
	data, err := c.CallHandleDeleteBill(ctx, args)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *MiddleWareClient) CallHandleUpdateStatus(ctx context.Context, args interface{}) (*rpc_model.RpcResponse, error) {
	data, err := c.CallHandleUpdateStatus(ctx, args)
	if err != nil {
		return nil, err
	}
	return data, nil
}
