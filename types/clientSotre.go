package types

import (
	"sync"

	"github.com/Jordany_dimbiniaina/chatForFun/interfaces"
)

type TCPClientStore struct {
	sync.Mutex
	clients map[string]interfaces.ReadWriteCloser
}


func (store *TCPClientStore) Load(addr string) (interfaces.ReadWriteCloser, bool) {
	store.Lock()
	defer store.Unlock()
	conn, exists := store.clients[addr]
	return conn, exists
}

func (store *TCPClientStore) Store(addr string, conn interfaces.ReadWriteCloser) {
	store.Lock()
	defer store.Unlock()
	store.clients[addr] = conn
}

func (store *TCPClientStore) Delete(addr string) {
	store.Lock()
	defer store.Unlock()
	delete(store.clients, addr)
}

func NewTCPClientStore() *TCPClientStore {
	return &TCPClientStore{
		clients: make(map[string]interfaces.ReadWriteCloser),
	}
}