package tcp

import (
	"net"
)

type Listener interface {
	Accept() (Conn, error)
	Close() error
	Addr() net.Addr
	SetBufLen(buflen int)
	SetChecksum(enable bool)
}

type listener struct {
	netListener net.Listener
	bufferLen   int
	checksum    bool
}

func (l *listener) SetBufLen(i int) {
	l.bufferLen = i
}

func (l *listener) SetChecksum(enable bool) {
	l.checksum = enable
}

func (l *listener) Accept() (Conn, error) {
	netConn, err := l.netListener.Accept()
	if err != nil {
		return nil, err
	}

	return newConn(connOption{
		netConn:   netConn,
		bufferLen: l.bufferLen,
		checksum:  l.checksum,
	}), nil
}

func (l *listener) Close() error {
	return l.netListener.Close()
}

func (l *listener) Addr() net.Addr {
	return l.netListener.Addr()
}

func Listen(address string) (Listener, error) {
	li, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &listener{netListener: li, bufferLen: BufferLen}, nil
}
