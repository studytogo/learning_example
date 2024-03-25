package reader

import "github.com/armon/circbuf"

func readF() []byte {
	rb, _ := circbuf.NewBuffer(1024)
	rb.Write([]byte("我很无耐！！！！！"))
	c := rb.Bytes()
	rb.Reset()
	return c
}
