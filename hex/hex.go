package hex

import (
	"fmt"
)

func displayableCharacter (ch byte) bool {
	return ch >= 32
}

func DumpAsHex(data []byte, displayWidth int) {
	
	for offset := 0; offset <= len(data); offset += displayWidth {
		str := fmt.Sprintf ("%08d:", offset)
		for rowIndex := 0; rowIndex < displayWidth; rowIndex++ {
			if offset + rowIndex < len(data) {
				str += fmt.Sprintf("%02X ", data[offset + rowIndex])
			} else {
				str += fmt.Sprintf("%02X ",0)
			}
		}
		str += "  "
		for rowIndex := 0; rowIndex < displayWidth; rowIndex++ {
			if offset + rowIndex < len(data) {
				if displayableCharacter(data[offset+rowIndex]) {
					str += fmt.Sprintf("%c", data[offset + rowIndex])
				} else {
					str += "."
				}
			} else {
				str += fmt.Sprintf("%c", 0)
			}
		}
		fmt.Println (str)		
	}
}

