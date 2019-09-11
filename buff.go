package gofpdf

// Buff for pdf content.
type Buff struct {
	position int
	data     []byte
}

// Write bytes to buffer.
func (b *Buff) Write(p []byte) (int, error) {
	for len(b.data) < b.position+len(p) {
		b.data = append(b.data, 0)
	}
	i := 0
	max := len(p)
	for i < max {
		b.data[i+b.position] = p[i]
		i++
	}
	b.position += i
	return 0, nil
}

// Len of buffer.
func (b *Buff) Len() int {
	return len(b.data)
}

// Bytes in Buff.
func (b *Buff) Bytes() []byte {
	return b.data
}

// Position returns current postion.
func (b *Buff) Position() int {
	return b.position
}

// SetPosition
func (b *Buff) SetPosition(pos int) {
	b.position = pos
}
