package tcp

import "net"

const BufferLen = 1024

type Conn interface {
	Read() ([]byte, error)
	Write(b []byte) (int, error)
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	SetBufLen(buflen int)
	SetChecksum(enable bool)
}

type conn struct {
	netConn    net.Conn
	packetInfo packetInfo
	rxRequest  RxQueue
	bufferLen  int
	checksum   bool
}

type connOption struct {
	netConn   net.Conn
	bufferLen int
	checksum  bool
}

func newConn(o connOption) Conn {
	return &conn{
		netConn:    o.netConn,
		packetInfo: newPacketInfo(),
		rxRequest:  NewRxQueue(),
		bufferLen:  o.bufferLen,
		checksum:   o.checksum,
	}
}

func (c *conn) SetBufLen(i int) {
	c.bufferLen = i
}

func (c *conn) SetChecksum(enable bool) {
	c.checksum = enable
}

func (c *conn) Read() ([]byte, error) {

	for {

		for len(c.rxRequest) > 0 {

			if c.checksum {
				tmp := c.rxRequest.Pop()
				if ChecksumVerify(tmp) {
					return tmp[1:], nil
				}
				continue
			}
			return c.rxRequest.Pop(), nil
		}

		b := make([]byte, c.bufferLen)
		n, err := c.netConn.Read(b)
		if err != nil {
			return nil, err
		}

		collectData(b[:n], c)
	}
}

func (c *conn) Write(b []byte) (int, error) {
	if c.checksum {
		b = AddChecksum(b)
	}
	toBeSend, err := genPacket(b)
	if err != nil {
		return 0, err
	}

	return c.netConn.Write(toBeSend)
}

func (c *conn) Close() error {
	return c.netConn.Close()
}

func (c *conn) LocalAddr() net.Addr {
	return c.netConn.LocalAddr()
}

func (c *conn) RemoteAddr() net.Addr {
	return c.netConn.RemoteAddr()
}
