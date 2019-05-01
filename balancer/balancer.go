package balancer

import (
	"fmt"
	"strconv"

	"google.golang.org/grpc/naming"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
)

// load balancer
// Copyright 2016 gRPC authors https://github.com/grpc/grpc-go/blob/master/balancer_test.go

type Watcher struct {
	// the channel to receives name resolution updates
	Update chan *naming.Update
	// the side channel to get to know how many updates in a batch
	Side chan int
	// the channel to notifiy update injector that the update reading is done
	ReadDone chan int
}

func (w *Watcher) Next() (updates []*naming.Update, err error) {
	n := <-w.Side
	if n == 0 {
		return nil, fmt.Errorf("w.side is closed")
	}
	for i := 0; i < n; i++ {
		u := <-w.Update
		if u != nil {
			updates = append(updates, u)
		}
	}
	w.ReadDone <- 0
	return
}

func (w *Watcher) Close() {
	close(w.Side)
}

// Inject naming resolution updates to the Watcher.
func (w *Watcher) Inject(updates []*naming.Update) {
	w.Side <- len(updates)
	for _, u := range updates {
		w.Update <- u
	}
	<-w.ReadDone
}

type NameResolver struct {
	W    *Watcher
	Addr string
}

func NewNameResolver(consulCli *consulapi.Client, service string) (*NameResolver, []string, error) {
	health, _, err := consulCli.Health().Service(service, "", false, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "can not get service health")
	}

	servers := []string{}
	for _, item := range health {
		address := item.Service.Address + ":" + strconv.Itoa(item.Service.Port)
		servers = append(servers, address)
	}

	if len(servers) == 0 {
		return nil, nil, errors.New("no available auth services")
	}

	nameResolver := &NameResolver{
		Addr: servers[0],
	}

	return nameResolver, servers, nil
}

func (r *NameResolver) LoadServers(servers []string) {
	if len(servers) > 1 {
		updates := []*naming.Update{}
		for i := 1; i < len(servers); i++ {
			updates = append(updates, &naming.Update{
				Op:   naming.Add,
				Addr: servers[i],
			})
		}
		r.W.Inject(updates)
	}
}

func (r *NameResolver) Resolve(target string) (naming.Watcher, error) {
	r.W = &Watcher{
		Update:   make(chan *naming.Update, 1),
		Side:     make(chan int, 1),
		ReadDone: make(chan int),
	}
	r.W.Side <- 1
	r.W.Update <- &naming.Update{
		Op:   naming.Add,
		Addr: r.Addr,
	}
	go func() {
		<-r.W.ReadDone
	}()
	return r.W, nil
}
