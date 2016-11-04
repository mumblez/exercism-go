package circular

//package main

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
func (b *Buffer) ReadByte() (byte, error) {
	// should pop off slice as reading the oldest / earliest
	if b.readposition > len(b.bit)-1 || b.bit[b.readposition] == 0 {
		b.readposition = 0

		return 0, errors.New("End of buffer")
	}
	a := b.bit[b.readposition]
	// reset bit
	b.bit[b.readposition] = 0
	b.readposition++
	return a, nil
}

// WriteByte and increment writeposition
func (b *Buffer) WriteByte(c byte) error {
	if b.writeposition > len(b.bit)-1 {
		if b.bit[0] == 0 {
			b.writeposition = 0
		} else {
			return errors.New("Buffer full")
		}
	}
	if b.bit[b.writeposition] == 0 {
		b.bit[b.writeposition] = c
		b.writeposition++
		return nil
	}
	return errors.New("Buffer full")
}

// Overwrite at current read position
func (b *Buffer) Overwrite(c byte) {
	// overwrite from read position
	// if b.writeposition < len(b.bit) {
	// 	_ = b.WriteByte(c)
	// 	return
	// }
	b.bit[b.readposition] = c
	b.readposition++
	if b.readposition > len(b.bit)-1 {
		b.readposition = 0
	}

}

// Reset all elements to 0
func (b *Buffer) Reset() {
	for k := range b.bit {
		b.bit[k] = 0
	}
	b.writeposition, b.readposition = 0, 0
}

// func readb(b *Buffer) {
// 	a, err := b.ReadByte()
// 	if err != nil {
// 		fmt.Printf("read err = %+v\n", err)
// 		return
// 	}
// 	fmt.Printf("read = %+v\n", string(a))
// }
// func writeb(b *Buffer, c byte) {
// 	err := b.WriteByte(c)
// 	if err != nil {
// 		fmt.Printf("write err = %+v\n", err)
// 		return
// 	}
// 	fmt.Printf("wrote = %+v\n", string(c))
//
// }
//
// func main() {
//
// 	a := NewBuffer(5)
// 	writeb(a, '1')
// 	writeb(a, '2')
// 	writeb(a, '3')
// 	readb(a)
// 	readb(a)
// 	writeb(a, '4')
// 	readb(a)
// 	writeb(a, '5')
// 	writeb(a, '6')
// 	writeb(a, '7')
// 	writeb(a, '8')
// 	fmt.Println("Overwrite with A and B")
// 	a.Overwrite('A')
// 	a.Overwrite('B')
// 	fmt.Println("6,7,8,A,B ???")
// 	readb(a)
// 	readb(a)
// 	readb(a)
// 	readb(a)
// 	readb(a)
// 	fmt.Println("we should fail...")
// 	readb(a)
//
// }
