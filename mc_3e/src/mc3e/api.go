package mc3e


type Client interface {
	//读取线圈
	ReadSCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadXCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadYCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadMCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadSMCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadLCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadFCoils(address uint16 ,quantity uint16)(results []byte,err error)
	//ReadVCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadBCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadTSCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadTCCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadSSCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadSCCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadCSCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadCCCoils(address uint16 ,quantity uint16)(results []byte,err error)
	ReadSBCoils(address uint16 ,quantity uint16)(results []byte,err error)

	//读取寄存器
	ReadDRegisters(address uint16 ,quantity uint16)(results []byte,err error)
	ReadSDRegisters(address uint16 ,quantity uint16)(results []byte,err error)
	ReadWRegisters(address uint16 ,quantity uint16)(results []byte,err error)
	ReadTNRegisters(address uint16 ,quantity uint16)(results []byte,err error)
	ReadSNRegisters(address uint16 ,quantity uint16)(results []byte,err error)
	ReadCNRegisters(address uint16 ,quantity uint16)(results []byte,err error)
	ReadSWRegisters(address uint16 ,quantity uint16)(results []byte,err error)
	ReadZRegisters(address uint16 ,quantity uint16)(results []byte,err error)
	ReadRRegisters(address uint16 ,quantity uint16)(results []byte,err error)
	//ReadZRRegisters(address uint16 ,quantity uint16)(results []byte,err error)

	//WriteDRegisters(address, quantity uint16, value []byte)(results []byte,err error)
	//WriteCoils(registeraddress string, quantity uint16, value []byte) (results []byte, err error)
	//WriteRegisters(registeraddress string, quantity uint16, value []byte) (results []byte, err error)

}