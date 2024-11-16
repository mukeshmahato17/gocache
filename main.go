package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/mukeshmahato17/gocache/cache"
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
		for i := 0; i < 10; i++ {
			SendCommand()
			time.Sleep(time.Millisecond * 200)
		}
	}()

	server := NewServer(opts, cache.New())
	server.Start()
}

func SendCommand() {
	cmd := &CommandSet{
		Key:   []byte("Foo"),
		Value: []byte("Bar"),
		TTL:   0,
	}

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	conn.Write(cmd.Bytes())
}
