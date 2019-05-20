package balancer

import (
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	consulapi "github.com/hashicorp/consul/api"
)

// ConnectClient подключение к gRPC сервису с балансировщиком
func ConnectClient(consulCli *consulapi.Client, service string) (*grpc.ClientConn, error) {
	nameResolver, servers, err := NewNameResolver(consulCli, service)
	if err != nil {
		return nil, errors.Wrap(err, "can not create name resolver")

	}

	grpcConn, err := grpc.Dial(
		servers[0],
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithBalancer(grpc.RoundRobin(nameResolver)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "can not connect to auth grpc")
	}

	nameResolver.LoadServers(servers)
	go OnlineServiceDiscovery(consulCli, nameResolver, service, servers, 15*time.Second)

	return grpcConn, nil
}
