package network

import (
	"net/http"
	"sync"
)

type Acceptor struct {
	waitGroup sync.WaitGroup
}

func NewAcceptor() Acceptor {
	return Acceptor{}
}

func (server *Acceptor) Accept() {
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()
	server.waitGroup.Add(1)
	server.waitGroup.Wait()
}