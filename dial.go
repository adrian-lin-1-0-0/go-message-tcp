package tcp

import "net"

func Dial(address string) (Conn, error) {
	netConn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return newConn(
		connOption{
			netConn:   netConn,
			bufferLen: BufferLen,
		}), nil
}
