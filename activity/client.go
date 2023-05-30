package v1

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"sync/atomic"
)

// AppID .
const ServiceName = "activity"

var (
	cli   = new(atomic.Value)
	clean = new(atomic.Value)
	count int32
)

func DefaultClient() ActivityClient {
	if c := cli.Load(); c == nil {
		client, cancel, err := NewClient()
		if err != nil {
			log.Error("DefaultClient Activity NewClient err:%v", err)
			return nil
		}
		cli.Store(client)
		clean.Store(cancel)
		atomic.AddInt32(&count, 1)
		return client
	} else {
		if atomic.CompareAndSwapInt32(&count, MaxCount, 1) {
			client, cancel, err := NewClient()
			if err != nil {
				log.Error("DefaultClient Activity NewClient err:%v", err)
				return nil
			}
			clean.Load().(func())()
			cli.Store(client)
			clean.Store(cancel)
			return client
		} else {
			atomic.AddInt32(&count, 1)
		}
		return c.(ActivityClient)
	}
}

func NewClient() (ActivityClient, func(), error) {
	clientSet, err := GetClientSet()
	if err != nil {
		log.Error("NewClient GetClientSet err:%v\n", err)
		return nil, nil, err
	}

	r := NewRegistry(clientSet)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(fmt.Sprintf("%s.%s.svc.cluster.local:9000", ServiceName, CrossNamespace)), // 可用
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Error("NewClient Grpc DialInsecure err %v", err)
		return nil, nil, err
	}
	return NewActivityClient(conn), func() {
		conn.Close()
	}, nil
}
