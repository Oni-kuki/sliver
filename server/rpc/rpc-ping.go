package rpc

import (
	"context"

	"github.com/Oni-kuki/sliver/protobuf/commonpb"
	"github.com/Oni-kuki/sliver/protobuf/sliverpb"
)

// Ping - Try to send a round trip message to the implant
func (rpc *Server) Ping(ctx context.Context, req *sliverpb.Ping) (*sliverpb.Ping, error) {
	resp := &sliverpb.Ping{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
