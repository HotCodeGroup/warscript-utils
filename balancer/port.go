package balancer

import (
	"encoding/json"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
)

type Bound struct {
	From int `json:"from"`
	To   int `json:"to"`
}

func findFree(consulCli *consulapi.Client, service string, b Bound) (int, error) {
	health, _, err := consulCli.Health().Service(service, "", false, nil)
	if err != nil {
		return -1, errors.Wrap(err, "can not get service health")
	}

	usedPorts := make(map[int]struct{})
	for _, item := range health {
		usedPorts[item.Service.Port] = struct{}{}
	}

	emptyPort := -1
	for port := b.From; port <= b.To; port++ {
		if _, ok := usedPorts[port]; !ok {
			emptyPort = port
			break
		}
	}
	if emptyPort == -1 {
		return -1, errors.New("no available ports")
	}

	return emptyPort, nil
}

func GetPorts(boundsKey, service string, consulCli *consulapi.Client) (int, int, error) {
	kv, _, err := consulCli.KV().Get(boundsKey, nil)
	if err != nil || kv == nil {
		return -1, -1, errors.Wrap(err, "can not get key")
	}

	bounds := make(map[string]Bound)
	err = json.Unmarshal(kv.Value, &bounds)
	if err != nil {
		return -1, -1, errors.Wrap(err, "can not unmarshal bounds")
	}

	httpPort := -1
	if _, ok := bounds["http"]; ok {
		httpPort, err = findFree(consulCli, service+"-http", bounds["http"])
		if err != nil {
			return -1, -1, errors.New("no available http ports")
		}
	}

	grpcPort := -1
	if _, ok := bounds["grpc"]; ok {
		grpcPort, err = findFree(consulCli, service+"-grpc", bounds["grpc"])
		if err != nil {
			return -1, -1, errors.New("no available grpc ports")
		}
	}

	return httpPort, grpcPort, nil
}
