package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	gatewayUserAgent = "grpcgateway-user-agent"
	grpcUserAgent = "user-agent"
	xForwardedFor = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtd := &Metadata{}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(gatewayUserAgent); len(userAgents) > 0 {
			mtd.UserAgent = userAgents[0]
		}

		if userAgents := md.Get(grpcUserAgent); len(userAgents) > 0 {
			mtd.UserAgent = userAgents[0]
		}

		if clientIPs := md.Get(xForwardedFor); len(clientIPs) > 0 {
			mtd.ClientIP = clientIPs[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtd.ClientIP = p.Addr.String()
	}

	return mtd
}