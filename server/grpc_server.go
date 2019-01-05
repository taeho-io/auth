package server

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
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

func New(cfg Config) (*AuthServer, error) {
	tokenCfg := token.NewConfig(
		cfg.Settings().SigningMethod,
		cfg.Settings().SigningPEM,
		cfg.Settings().VerifyingPEM,
		cfg.Settings().TokenIssuer,
		cfg.Settings().AccessTokenExpiringDuration,
		cfg.Settings().RefreshTokenExpiringDuration,
	)
	tkn, err := token.New(tokenCfg)
	if err != nil {
		return nil, err
	}

	return &AuthServer{
		cfg: cfg,
		tkn: tkn,
	}, nil
}

func Mock() *AuthServer {
	s, _ := New(MockConfig())
	return s
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

func (s *AuthServer) Verify(ctx context.Context, req *auth.VerifyRequest) (*auth.VerifyResponse, error) {
	return handler.Verify(s.Token())(ctx, req)
}

func (s *AuthServer) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	return handler.Refresh(s.Config().(Config).Settings().AccessTokenExpiringDuration, s.Token())(ctx, req)
}

func (s *AuthServer) Parse(ctx context.Context, req *auth.ParseRequest) (*auth.ParseResponse, error) {
	return handler.Parse(s.Token())(ctx, req)
}

func (s *AuthServer) Jwks(ctx context.Context, req *auth.JwksRequest) (*auth.JwksResponse, error) {
	return handler.Jwks(s.Config().VerifyingKey())(ctx, req)
}

func NewGRPCServer(cfg Config) (*grpc.Server, error) {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	authServer, err := New(cfg)
	if err != nil {
		return nil, err
	}
	auth.RegisterAuthServer(grpcServer, authServer)
	reflection.Register(grpcServer)

	return grpcServer, nil
}

func ServeGRPC(addr string, cfg Config) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	grpcServer, err := NewGRPCServer(cfg)
	if err != nil {
		return err
	}

	return grpcServer.Serve(lis)
}
