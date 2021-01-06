package handler

import (
	"context"
	"github.com/ashleywang1/capture-the-flag-apiserver/api"
)

func NewHandler() api.CaptureTheFlagApiServer {
	return &handler{}
}

type handler struct {}

func (k *handler) CaptureTheFlag(ctx context.Context, request *api.CaptureTheFlagRequest) (*api.CaptureTheFlagResponse, error) {
	return &api.CaptureTheFlagResponse{
		Flag:&api.Flag{
			Flag:"[the flag will be here]",
		},
	}, nil
}
