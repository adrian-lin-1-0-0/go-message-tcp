package tcp

type RxQueue [][]byte

func NewRxQueue() RxQueue {
	return [][]byte{}
}

func (q *RxQueue) Push(b []byte) {
	*q = append(*q, b)
}

func (q *RxQueue) Pop() []byte {
	b := (*q)[0]
	*q = (*q)[1:]
	return b
}
