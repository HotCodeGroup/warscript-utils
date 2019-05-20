package balancer

import (
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc/naming"

	consulapi "github.com/hashicorp/consul/api"
)

// OnlineServiceDiscovery поддержание списка активнх gRPC сервисов
func OnlineServiceDiscovery(consulCli *consulapi.Client, resolver *NameResolver,
	service string, servers []string, period time.Duration) {
	currAddrs := make(map[string]struct{}, len(servers))
	for _, addr := range servers {
		currAddrs[addr] = struct{}{}
	}

	ticker := time.Tick(period)
	for _ = range ticker {
		health, _, err := consulCli.Health().Service(service, "", false, nil)
		if err != nil {
			log.Printf("cant get alive services: %s", err)
		}

		newAddrs := make(map[string]struct{}, len(health))
		for _, item := range health {
			addr := item.Service.Address +
				":" + strconv.Itoa(item.Service.Port)
			newAddrs[addr] = struct{}{}
		}

		var updates []*naming.Update
		// проверяем что удалилось
		for addr := range currAddrs {
			if _, exist := newAddrs[addr]; !exist {
				updates = append(updates, &naming.Update{
					Op:   naming.Delete,
					Addr: addr,
				})
				delete(currAddrs, addr)
			}
		}

		// проверяем что добавилось
		for addr := range newAddrs {
			if _, exist := currAddrs[addr]; !exist {
				updates = append(updates, &naming.Update{
					Op:   naming.Add,
					Addr: addr,
				})
				currAddrs[addr] = struct{}{}
			}
		}

		if len(updates) > 0 {
			resolver.W.Inject(updates)
		}
	}
}
