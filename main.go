package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/mukeshmahato17/gocache/cache"
	"github.com/mukeshmahato17/gocache/client"
)

func main() {

	var (
		listenAddr = flag.String("listenaddr", ":3000", "listen address of the server")
		leaderAddr = flag.String("leaderaddr", "", "listen address of the leader")
	)
	flag.Parse()

	opts := ServerOpts{
		ListenAddr: *listenAddr,
		IsLeader:   len(*leaderAddr) == 0,
		LeaderAddr: *leaderAddr,
	}

	go func() {
		time.Sleep(time.Second * 2)
		client, err := client.New(":3000", client.Options{})
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < 10; i++ {
			SendCommand(client)
		}
		client.Close()
		time.Sleep(time.Second * 1)
	}()

	server := NewServer(opts, cache.New())
	server.Start()
}

func SendCommand(client *client.Client) {
	_, err := client.Set(context.Background(), []byte("Foo"), []byte("Bar"), 0)
	if err != nil {
		log.Fatal(err)
	}

}
