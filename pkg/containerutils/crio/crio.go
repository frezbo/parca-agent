// Copyright 2021 The Parca Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crio

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/go-kit/log"
	"google.golang.org/grpc"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

const (
	DEFAULT_SOCKET_PATH = "/run/crio/crio.sock"
	DEFAULT_TIMEOUT     = 2 * time.Second
)

type CrioClient struct {
	logger log.Logger
	conn   *grpc.ClientConn
	client pb.RuntimeServiceClient
}

func NewCrioClient(logger log.Logger, path string) (*CrioClient, error) {
	conn, err := grpc.Dial(
		path,
		grpc.WithInsecure(),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			var d net.Dialer
			ctx, cancel := context.WithTimeout(ctx, DEFAULT_TIMEOUT)
			defer cancel()

			d.LocalAddr = nil // if you have a local addr, add it here
			raddr := net.UnixAddr{Name: path, Net: "unix"}
			return d.DialContext(ctx, "unix", raddr.String())
		}),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewRuntimeServiceClient(conn)
	return &CrioClient{
		logger: logger,
		conn:   conn,
		client: client,
	}, nil
}

func (c *CrioClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}

	return nil
}

type containerInfo struct {
	PID int `json:"pid"`
}

func (c *CrioClient) PidFromContainerId(containerID string) (int, error) {
	if !strings.HasPrefix(containerID, "cri-o://") {
		return -1, fmt.Errorf("Invalid CRI %s, it should be cri-o", containerID)
	}

	containerID = strings.TrimPrefix(containerID, "cri-o://")

	request := &pb.ContainerStatusRequest{
		ContainerId: containerID,
		Verbose:     true,
	}

	status, err := c.client.ContainerStatus(context.Background(), request)
	if err != nil {
		return -1, err
	}

	infoStr, ok := status.Info["info"]
	if !ok {
		return -1, fmt.Errorf("container status reply from runtime does not contain 'info'")
	}

	var ci containerInfo
	if err := json.Unmarshal([]byte(infoStr), &ci); err != nil {
		return -1, fmt.Errorf("could not unmarshal container info")
	}

	return ci.PID, nil
}
