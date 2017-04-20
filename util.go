package ptreader

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (r *Reader) readInt() (int, error) {
	four := make([]byte, 4)
	n, err := r.r.Read(four)
	if n != 4 {
		return 0, fmt.Errorf("read int wrong")
	}
	if err != nil {
		return 0, err
	}
	var one int32
	_ = binary.Read(bytes.NewReader(four), binary.LittleEndian, &one)
	return int(one), nil
}

func (r *Reader) readUint16() (uint16, error) {
	two := make([]byte, 2)
	n, err := r.r.Read(two)
	if n != 2 {
		return 0, fmt.Errorf("read uint16 wrong")
	}
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(two), nil
}

func (r *Reader) readByte() (byte, error) {
	one := make([]byte, 1)
	n, err := r.r.Read(one)
	if n != 1 {
		return 0, fmt.Errorf("read int wrong")
	}
	if err != nil {
		return 0, err
	}
	return one[0], nil
}

func (r *Reader) readString() (string, error) {
	var buf []byte
	for {
		b, err := r.readByte()
		if err != nil {
			return "", err
		}
		if b == 0 {
			break
		}
		buf = append(buf, b)
	}
	return string(buf), nil
}

// FormatSeconds to format seconds to human readable format
func FormatSeconds(seconds int64) string {
	days := seconds / 86400

	left := seconds - days*86400

	hours := left / 3600

	left = left - hours*3600

	minutes := left / 60

	left = left - minutes*60

	return fmt.Sprintf("%2dD%2dH%2dM%2dS", days, hours, minutes, left)
}
