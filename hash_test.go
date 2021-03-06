package ring

import "testing"

func BenchmarkGenerateMultiHash(b *testing.B) {
	data := []byte{0x00, 0x12, 0x34, 0x56, 0x78, 0x00}
	buff := make([]byte, len(data))
	for i := 0; i < b.N; i++ {
		copy(buff, data)
		generateMultiHash(buff[1:5])
	}
}

func TestGenerateMultiHash(t *testing.T) {
	data := []byte{
		0x00, 0x12, 0x34, 0x56, 0x78, 0x00,
		0x00, 0x12, 0x34, 0x56, 0x78, 0x00,
		0x00, 0x12, 0x34, 0x56, 0x78, 0x00,
		0x00, 0x12, 0x34, 0x56, 0x78, 0x00,
		0x00, 0x12, 0x34, 0x56, 0x78, 0x00,
		0x00, 0x12, 0x34, 0x56, 0x78, 0x00,
		0x00, 0x12, 0x34, 0x56, 0x78, 0x00,
		0x00, 0x12, 0x34, 0x56, 0x78, 0x00,
		0x00, 0x12, 0x34, 0x56, 0x78, 0x00,
		0x00, 0x12, 0x34, 0x56, 0x78, 0x00,
		0x00, 0x12, 0x34, 0x56, 0x78, 0x00,
	}
	buff := make([]byte, len(data))
	copy(buff, data)
	generateMultiHash(buff[1:20])

	for i := range data {
		if data[i] != buff[i] {
			t.Fatalf("data not match at index: %v", i)
		}
	}
}
