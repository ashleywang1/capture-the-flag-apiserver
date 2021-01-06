package server

import (
	"context"
	"fmt"
	"github.com/ashleywang1/capture-the-flag-apiserver/api"
	"github.com/rotisserie/eris"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
)

func init() {
	view.Register(ocgrpc.DefaultServerViews...)
}

type GrpcServer interface {
	Run(ctx context.Context) error
	Stop()
}

type grpcServer struct {
	server        *grpc.Server
}

func NewGrpcServer(
	ctx context.Context,
	captureTheFlagService api.CaptureTheFlagApiServer,
) GrpcServer {
	fmt.Printf("Started up GRPC Server")
	server := grpc.NewServer()
	api.RegisterCaptureTheFlagApiServer(server, captureTheFlagService)

	return &grpcServer{
		server:        server,
	}
}

func (g *grpcServer) Run(ctx context.Context) error {
	var eg errgroup.Group
	// start grpc listener
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", 1234))
	if err != nil {
		return eris.Wrapf(err, "failed to setup grpc listener")
	}
	fmt.Printf("Set up grpc listener at port %d", 1234)
	eg.Go(func() error {
		return g.server.Serve(grpcListener)
	})
	return eg.Wait()
}

func (g *grpcServer) Stop() {
	g.server.GracefulStop()
}
