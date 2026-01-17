package server

import (
	"io"
)

type ReadWriteCloser interface {
	io.Reader
	io.Writer
	io.Closer
}

type ClientStore interface {
	Store(addr string, conn ReadWriteCloser)
	Load(addr string) (ReadWriteCloser, bool)
	Delete(addr string)
}
