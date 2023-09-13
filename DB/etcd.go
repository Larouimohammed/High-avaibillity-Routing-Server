package db

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

var (
	client   = SetupDb()
	ctx      = context.Background()
	Messages = make(chan string)
)

func SetupDb() *clientv3.Client {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://172.17.0.1:2379"},
		DialTimeout: 5 * time.Hour,
		DialOptions: []grpc.DialOption{},
	})
	if err != nil {
		log.Fatal(err)
	}

	// defer cli.Close()
	return cli

}
func Publish(key string) {

	for {

		_, err := client.Put(ctx, key, time.Now().Format(time.RFC1123))
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)

	}

}

func Watch(Key string) {
	ch := client.Watch(ctx, Key)
loop:
	for resp := range ch {
		for _, event := range resp.Events {
			switch event.Type {
			case clientv3.EventTypePut:
				msg, err := client.Get(ctx, Key)
				if err != nil {
					log.Fatal(err)

				}
				for _, ev := range msg.Kvs {

					t, _ := time.Parse(time.RFC1123, string(ev.Value))
					fmt.Println(string(ev.Value))

					if time.Now().After(t.Add(10 * time.Second)) {
						Messages <- "change router"
						break loop
					}

				}

			case clientv3.EventTypeDelete:
				// process with delete event

			}
		}

	}

}
