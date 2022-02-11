package handler

import "context"

type Handler interface {
	Handler(ctx context.Context) (interface{}, error)
}

/**
API Gateway Implement
*/
type handler struct {
	req interface{}
}

func NewApiGW(req interface{}) Handler { return &handler{req} }

func (a handler) Handler(ctx context.Context) (interface{}, error) {
	panic("implement me")
}
