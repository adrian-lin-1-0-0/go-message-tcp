package tcp

func Checksum(buf []byte) byte {
	sum := byte(0)
	for _, v := range buf {
		sum = sum ^ v
	}
	sum = sum ^ 0xff
	return sum
}

func ChecksumVerify(buf []byte) bool {
	return Checksum(buf) == 0
}

func AddChecksum(buf []byte) []byte {
	return append([]byte{Checksum(buf)}, buf...)
}
