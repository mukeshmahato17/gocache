package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mukeshmahato17/gocache/cache"
)

type ServerOpts struct {
	ListenAddr string
	IsLeader   bool
	LeaderAddr string
}

type Server struct {
	ServerOpts

	cache cache.Cacher
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("listen error: %s", err)
	}

	log.Printf("server starting on port [%s]\n", s.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %s\n", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	// buf := make([]byte, 2048)

	fmt.Println("connection made:", conn.RemoteAddr())
	for {
		cmd, err := ParseCommand(conn)
		if err != nil {
			log.Println("invalid command error ", err)
		}
		go s.handleCommand(conn, cmd)
	}
}

func (s *Server) handleCommand(conn net.Conn, cmd any) {
	switch v := cmd.(type) {
	case *CommandSet:
		s.handleSetCommand(conn, v)
	case *CommandGet:
	}
}

func (s *Server) handleSetCommand(conn net.Conn, cmd *CommandSet) error {
	fmt.Printf("SET %s to %s\n", cmd.Key, cmd.Value)

	return s.cache.Set(cmd.Key, cmd.Value, time.Duration(cmd.TTL))
}
