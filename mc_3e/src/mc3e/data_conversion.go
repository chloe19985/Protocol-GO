package mc3e
func dataBlock(value ...uint16) []byte {
	b := make([]uint16, 2)
	data := make([]byte, 5)
	for i, v := range value {
		b[i] = v
	}

	data[0] = byte(b[0])
	data[1] = byte(b[0] >> 8)
	data[2] = byte(b[0] >> 16)
	data[3] = byte(b[1])
	data[4] = byte(b[1] >> 8)

	return data
}

func addressBlock(value uint16) []byte {
	data := make([]byte, 3)
	data[0] = byte(value)
	data[1] = byte(value >> 8)
	data[2] = byte(value >> 16)
	return data
}
func quantityBlock(value uint16) []byte {
	data := make([]byte, 2)
	data[0] = byte(value)
	data[1] = byte(value >> 8)

	return data
}

/*
func dataBlockSuffix(suffix []byte, value ...uint16) []byte {
	b := make([]uint16,2)
	data := make([]byte, 5 + len(suffix))
	for i, v := range value {
		b[i] = v
	}
	data[0] = byte(b[0]>>16)
	data[1] = byte(b[0]>>8)
	data[2] = byte(b[0])
	data[3] = byte(b[1]>>8)
	data[4] = byte(b[1])
	copy(data[5:],suffix)
	return data
}
*/

func Uint16ToBytes(v uint16) []byte {
	b := make([]byte, 2)
	//_ = b[1] // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 8)
	b[1] = byte(v)
	return b
}


func McSetBinToChar(tab []byte , index uint8, value uint16) {
	tab[(index)]   = McBinToChar(byte(value) >> 4)
	tab[(index)+1] = McBinToChar(byte(value) & 0x0F)
}
func McBinToChar( ucByte byte) (b byte){
	if (ucByte <= 0x09) {
		b = '0' + ucByte
		return b
	}
	if ((ucByte >= 0x0A) && (ucByte <= 0x0F)) {
		b := ucByte - 0x0A + 'A'
		return b
	}
	return
}
