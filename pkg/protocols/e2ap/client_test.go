// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2ap

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSourceIPForLoopback(t *testing.T) {
	ip, err := sourceIPFor(net.ParseIP("127.0.0.1"))
	assert.NoError(t, err)
	assert.True(t, ip.Equal(net.ParseIP("127.0.0.1")), "got %s", ip)
}

func TestSourceIPForMatchesLocalInterface(t *testing.T) {
	ip, err := sourceIPFor(net.ParseIP("192.0.2.1"))
	if err != nil {
		t.Skipf("no IPv4 route available: %v", err)
	}
	addrs, err := net.InterfaceAddrs()
	assert.NoError(t, err)
	found := false
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && ipnet.IP.Equal(ip) {
			found = true
		}
	}
	assert.True(t, found, "source IP %s is not assigned to any local interface", ip)
}
