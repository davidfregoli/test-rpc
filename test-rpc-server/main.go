package main

import (
	"log"
	"net"
	"net/rpc"
	"sync"
)

type KV struct {
	mu   sync.Mutex
	data map[string]string
}

type GetArgs struct {
	Key string
}

type GetReply struct {
	Value string
}

type PutArgs struct {
	Key   string
	Value string
}

type PutReply struct {
}

func main() {
	kv := &KV{data: map[string]string{}}
	rpcs := rpc.NewServer()
	rpcs.Register(kv)
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		for {
			conn, err := l.Accept()
			if err == nil {
				go rpcs.ServeConn(conn)
			} else {
				break
			}
		}
		l.Close()
	}()
	for {
	}
}

func (kv *KV) Get(args *GetArgs, reply *GetReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	reply.Value = kv.data[args.Key]

	return nil
}

func (kv *KV) Put(args *PutArgs, reply *PutReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.data[args.Key] = args.Value

	return nil
}
