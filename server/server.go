package server

import (
	"net"

	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/pkg/token"
	"github.com/taeho-io/auth/server/handler"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type AuthServer struct {
	auth.AuthServer

	cfg Config
	tkn token.Token
}

func New(cfg Config) *AuthServer {
	tokenCfg := token.NewConfig(
		cfg.Settings().SigningMethod,
		cfg.Settings().SigningKey,
		cfg.Settings().TokenIssuer,
		cfg.Settings().AccessTokenExpiringDuration,
		cfg.Settings().RefreshTokenExpiringDuration,
	)
	tkn := token.New(tokenCfg)

	return &AuthServer{
		cfg: cfg,
		tkn: tkn,
	}
}

func Mock() *AuthServer {
	return New(MockConfig())
}

func (s *AuthServer) Config() Config {
	return s.cfg
}

func (s *AuthServer) RegisterServer(srv *grpc.Server) {
	auth.RegisterAuthServer(srv, s)
}

func (s *AuthServer) Token() token.Token {
	return s.tkn
}

func (s *AuthServer) Auth(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	return handler.Auth(s.Config().Settings().AccessTokenExpiringDuration, s.Token())(ctx, req)
}

func (s *AuthServer) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	return handler.Validate(s.Token())(ctx, req)
}

func (s *AuthServer) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	return handler.Refresh(s.Config().(Config).Settings().AccessTokenExpiringDuration, s.Token())(ctx, req)
}

func (s *AuthServer) Parse(ctx context.Context, req *auth.ParseRequest) (*auth.ParseResponse, error) {
	return handler.Parse(s.Token())(ctx, req)
}

func Serve() error {
	lis, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	authServer := New(NewConfig(NewSettings()))
	auth.RegisterAuthServer(grpcServer, authServer)

	reflection.Register(grpcServer)
	return grpcServer.Serve(lis)
}
