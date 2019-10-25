package Common

// StreamReader allows you to read data from a byte array in various formats
type StreamReader struct {
	data     []byte
	position uint64
}

// CreateStreamReader creates an instance of the stream reader
func CreateStreamReader(source []byte) *StreamReader {
	result := &StreamReader{
		data:     source,
		position: 0,
	}
	return result
}

// GetPosition returns the current stream position
func (v *StreamReader) GetPosition() uint64 {
	return v.position
}

// GetSize returns the total size of the stream in bytes
func (v *StreamReader) GetSize() uint64 {
	return uint64(len(v.data))
}

// GetByte returns a byte from the stream
func (v *StreamReader) GetByte() byte {
	result := v.data[v.position]
	v.position++
	return result
}

// GetWord returns a uint16 word from the stream
func (v *StreamReader) GetWord() uint16 {
	result := uint16(v.data[v.position])
	result += (uint16(v.data[v.position+1]) << 8)
	v.position += 2
	return result
}

// GetDword returns a uint32 dword from the stream
func (v *StreamReader) GetDword() uint32 {
	result := uint32(v.data[v.position])
	result += (uint32(v.data[v.position+1]) << 8)
	result += (uint32(v.data[v.position+2]) << 16)
	result += (uint32(v.data[v.position+3]) << 24)
	v.position += 4
	return result
}
