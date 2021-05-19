package mcbin


type Client interface {
	//读取线圈
	//ReadSCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadXCoils(address uint16 ,quantity uint16)(results []byte,err error)
	//ReadYCoils(address uint16 ,quantity uint16)(results []byte,err error)
	//ReadMCoils(address uint16 ,quantity uint16)(results []byte,err error)
	//ReadSMCoils(address uint16 ,quantity uint16)(results []byte,err error)
	//ReadLCoils(address uint16 ,quantity uint16)(results []byte,err error)
	//读取寄存器
	ReadDRegisters(address uint16 ,quantity uint16)(results []byte,err error)
	//WriteDRegisters(address, quantity uint16, value []byte)(results []byte,err error)
	//WriteCoils(registeraddress string, quantity uint16, value []byte) (results []byte, err error)
	//WriteRegisters(registeraddress string, quantity uint16, value []byte) (results []byte, err error)
}