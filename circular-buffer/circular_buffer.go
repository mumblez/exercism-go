package circular

import "errors"

const testVersion = 4

// Buffer type holds the byte slice and fields for read and write position status
type Buffer struct {
	bit           []byte
	readposition  int
	writeposition int
}

// NewBuffer returns a new Buffer
func NewBuffer(size int) *Buffer {
	return &Buffer{bit: make([]byte, size)}
}

// ReadByte and increment readposition
func (b *Buffer) ReadByte() (res byte, err error) {
	// error if current bit = 0
	if b.bit[b.readposition] == 0 {
		err = errors.New("End of buffer")
		return
	}
	res = b.bit[b.readposition]
	b.bit[b.readposition] = 0                            // reset after read
	if b.readposition == len(b.bit)-1 && b.bit[0] != 0 { // if on last element and the first element isn't 0 then cycle round
		b.readposition = 0
		return
	}
	if b.readposition < len(b.bit)-1 {
		b.readposition++
	}
	return
}

// WriteByte and increment writeposition
func (b *Buffer) WriteByte(c byte) error {
	// stop if reach bit with value
	if b.bit[b.writeposition] != 0 {
		return errors.New("Buffer full")
	}
	b.bit[b.writeposition] = c
	// cycle if at end and beginning is 0
	if b.writeposition == len(b.bit)-1 && b.bit[0] == 0 {
		b.writeposition = 0
		return nil
	}
	if b.writeposition < len(b.bit)-1 {
		b.writeposition++
	}
	return nil
}

// Overwrite at current read position
func (b *Buffer) Overwrite(c byte) {
	// overwrite from read position
	if b.bit[b.writeposition] == 0 {
		_ = b.WriteByte(c)
		return
	}
	b.bit[b.readposition] = c
	// increment/cycle readposition after a real overwrite
	if b.readposition == len(b.bit)-1 && b.bit[0] != 0 { // if on last element and the first element isn't 0 then cycle round
		b.readposition = 0
		return
	}
	if b.readposition < len(b.bit)-1 {
		b.readposition++
	}
}

// Reset all elements to 0
func (b *Buffer) Reset() {
	for k := range b.bit {
		b.bit[k] = 0
	}
	b.writeposition, b.readposition = 0, 0
}
