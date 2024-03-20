// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: v1/movie.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAccountServiceListMovies = "/v1.AccountService/ListMovies"

type AccountServiceHTTPServer interface {
	ListMovies(context.Context, *ListMovieRequest) (*ListMovieResponse, error)
}

func RegisterAccountServiceHTTPServer(s *http.Server, srv AccountServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/movies", _AccountService_ListMovies0_HTTP_Handler(srv))
}

func _AccountService_ListMovies0_HTTP_Handler(srv AccountServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMovieRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountServiceListMovies)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMovies(ctx, req.(*ListMovieRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMovieResponse)
		return ctx.Result(200, reply)
	}
}

type AccountServiceHTTPClient interface {
	ListMovies(ctx context.Context, req *ListMovieRequest, opts ...http.CallOption) (rsp *ListMovieResponse, err error)
}

type AccountServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewAccountServiceHTTPClient(client *http.Client) AccountServiceHTTPClient {
	return &AccountServiceHTTPClientImpl{client}
}

func (c *AccountServiceHTTPClientImpl) ListMovies(ctx context.Context, in *ListMovieRequest, opts ...http.CallOption) (*ListMovieResponse, error) {
	var out ListMovieResponse
	pattern := "/movies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAccountServiceListMovies))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
