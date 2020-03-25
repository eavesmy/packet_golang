package packet_golang

import (
	"encoding/binary"
	"io"
)

type Unpack struct {
	size       int
	bufferpool []byte
	ret        []byte
	id         uint32
}

func NewUnpack(size int) *Unpack {
	return &Unpack{
		bufferpool: make([]byte, 0, size),
		ret:        []byte{},
		size:       size,
	}
}

func (u *Unpack) Deal(data []byte) (err error) {

	// not enough data for unpack id & length
	u.bufferpool = append(u.bufferpool, data...)

	if len(u.bufferpool) < 6 {
		return io.EOF
	}

	length := binary.BigEndian.Uint16(u.bufferpool[4:6]) + 7 // 1 is '\n' length.

	if len(u.bufferpool) < int(length) {
		return io.EOF
	}

	u.id = binary.BigEndian.Uint32(u.bufferpool[:4])

	u.ret = u.bufferpool[6 : length-1]

	u.emptyPool(int(length))

	return
}

func (u *Unpack) emptyPool(limit int) {
	u.bufferpool = make([]byte, 0, u.size)
}

func (u *Unpack) Bytes() []byte {
	return u.ret
}

func (u *Unpack) Pid() uint32 {
	return u.id
}
