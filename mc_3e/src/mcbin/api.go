package mcbin


type Client interface {
	//读取线圈
	ReadXCoils(address uint16 ,quantity uint16)(results []byte,err error)
	//读取寄存器
	ReadDRegisters(address uint16 ,quantity uint16)(results []byte,err error)
	//WriteDRegisters(address, quantity uint16, value []byte)(results []byte,err error)
	//WriteCoils(registeraddress string, quantity uint16, value []byte) (results []byte, err error)
	//WriteRegisters(registeraddress string, quantity uint16, value []byte) (results []byte, err error)
}