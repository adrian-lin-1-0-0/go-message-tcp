package tcp

import (
	encode "github.com/adrian-lin-1-0-0/go-variable-length-integer-encoding"
)

type packetInfo struct {
	length    uint64
	buf       []byte
	headerLen uint8
	total     uint64
}

func (p *packetInfo) clear() {
	*p = newPacketInfo()
}

func newPacketInfo() packetInfo {
	return packetInfo{
		length:    0,
		buf:       []byte{},
		headerLen: 0,
		total:     0,
	}
}

func collectData(data []byte, c *conn) {
	if c.packetInfo.length == 0 {
		parsePacketInfo(&c.packetInfo, &data)
	}
	remaining := int(c.packetInfo.total) - len(data) - len(c.packetInfo.buf)
	if remaining == 0 {
		c.rxRequest.Push(append(c.packetInfo.buf, data...)[c.packetInfo.headerLen:])
		c.packetInfo.clear()
		return
	} else if remaining > 0 {
		c.packetInfo.buf = append(c.packetInfo.buf, data...)
		return
	} else {
		c.rxRequest.Push(append(c.packetInfo.buf, data[:len(data)+remaining]...)[c.packetInfo.headerLen:])
		c.packetInfo.clear()
		collectData(data[len(data)+remaining:], c)
	}
}

func parsePacketInfo(p *packetInfo, data *[]byte) {
	length, varlen := encode.ToUint64(*data)
	p.length = length
	p.headerLen = varlen
	p.total = p.length + uint64(p.headerLen)
}

func genPacket(data []byte) ([]byte, error) {
	varlenInt, err := encode.ToVarLenInt(uint64(len(data)))
	if err != nil {
		return nil, err
	}
	return append(varlenInt, data...), nil
}
