// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2ap

import (
	"context"
	"net"

	"github.com/LABORA-INF-UFG/GUARA-utils/pkg/sctp/addressing"
	"github.com/LABORA-INF-UFG/GUARA-utils/pkg/sctp/types"

	"github.com/LABORA-INF-UFG/GUARA-e2t/pkg/protocols/e2ap/procedures"

	sctp "github.com/LABORA-INF-UFG/GUARA-utils/pkg/sctp"
)

// ClientHandler is a client handler function
type ClientHandler func(conn ClientConn) ClientInterface

// ClientInterface is an interface for E2 client procedures
type ClientInterface procedures.E2NodeProcedures

// ConnectOptions are options for Connect
type ConnectOptions struct {
	localIP net.IP
}

// ConnectOption is a Connect option
type ConnectOption func(*ConnectOptions)

// WithLocalIP sets the local IP. Defaults to the source IP of the route.
func WithLocalIP(ip net.IP) ConnectOption {
	return func(options *ConnectOptions) {
		options.localIP = ip
	}
}

// Connect connects to the given address using a single local IP
func Connect(_ context.Context, address string, handler ClientHandler, opts ...ConnectOption) (ClientConn, error) {
	options := &ConnectOptions{}
	for _, opt := range opts {
		opt(options)
	}
	addr, err := addressing.ResolveAddress(types.Sctp4, address)
	if err != nil {
		return nil, err
	}
	dialOpts := []sctp.DialOption{
		sctp.WithAddressFamily(addr.AddressFamily),
		sctp.WithNonBlocking(false),
		sctp.WithMode(types.OneToOne),
		sctp.WithInitMsg(types.InitMsg{}),
	}
	localIP := options.localIP
	if localIP == nil {
		localIP, err = sourceIPFor(addr.IPAddrs[0].IP)
		if err != nil {
			log.Warnf("Could not determine source IP for %s, SCTP association will be multi-homed: %v", address, err)
		}
	}
	if localIP != nil {
		dialOpts = append(dialOpts, sctp.WithLocalAddress(localIP))
	}
	c, err := sctp.DialSCTP(addr, dialOpts...)
	if err != nil {
		return nil, err
	}
	conn := NewClientConn(c, func(conn ClientConn) ClientInterface {
		return handler(conn)
	})
	return conn, nil
}

// sourceIPFor returns the source IP used to reach dst. No packet is sent.
func sourceIPFor(dst net.IP) (net.IP, error) {
	conn, err := net.DialUDP("udp4", nil, &net.UDPAddr{IP: dst, Port: 9})
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP, nil
}
