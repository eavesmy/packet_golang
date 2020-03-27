package packet_golang

import (
	"encoding/binary"
	"encoding/json"
	"reflect"
)

// create a new data pack and return bytes.
func NewPacket(id uint32, data interface{}) []byte {
	t := reflect.TypeOf(data)

	if t.Kind() != reflect.Struct {
		panic("Data must be a struct.")
	}

	b, _ := json.Marshal(data)

	buffer_id := make([]byte, 4, 4)
	buffer_len := make([]byte, 2, 2)

	length := len(b)

	binary.BigEndian.PutUint32(buffer_id, id)

	binary.BigEndian.PutUint16(buffer_len, uint16(length))

	_buffer := make([]byte, 0, 7+len(b))

	_buffer = append(_buffer, buffer_id...)
	_buffer = append(_buffer, buffer_len...)
	_buffer = append(_buffer, b...)

	_buffer = append(_buffer, '\n')

	return _buffer
}
