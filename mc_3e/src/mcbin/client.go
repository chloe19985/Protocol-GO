package mcbin

import (
	"fmt"
)

type ClientHandler interface {
	Packager
	Transporter
}
//
type client struct {
	packager    Packager
	transporter Transporter
}

//// NewClient creates a new modbus client with given backend handler.
func NewClient(handler ClientHandler) Client {
	//return &client{transporter: handler}
	return &client{packager: handler, transporter: handler}
}
func (mb *client) ReadXCoils(address uint16,quantity uint16)(results []byte,err error){

	/*switch registercode {
	case MC_3E_BIN_TYPE_S:*/
		if quantity < 0 || quantity > 4095 {
			err = fmt.Errorf("mc_3e-X:quantity '%v' must be between '%v' and '%v'", quantity, 1, 4095)
			return
		}
		request := ProtocolDataUnit{
			Retain:               Int16ToBytes(Retain_Command),
			Command:              Int16ToBytes(DeviceRead),
			SubCommand:           Int16ToBytes(SubCommand1),
			SoftComponentCode:    MC_3E_BIN_TYPE_X,
			//SoftComponentAddress: addressBlock(address),
			//SoftComponentNumber:  Int16ToBytes(quantity),
			Data:                 dataBlock(address, quantity),
		}
		response, err := mb.send(&request)
		if err != nil {
			return
		}
		length := len(response.Data)
		if length <= 0 {
			err = fmt.Errorf("mc3e:response data size '%v' can not less than 0 ", length)
			return
		}
		results = response.Data[0:]
	//}
		return
}

func (mb *client) ReadDRegisters(address uint16 ,quantity uint16)(results []byte,err error){
	if quantity < 0 || quantity >7998{
		err = fmt.Errorf("mc_3e_D:quantity '%v' must be between '%v' and '%v'",quantity,1,7998)
		return
	}
	request := ProtocolDataUnit{
		Retain:               Int16ToBytes(Retain_Command),
		Command:              Int16ToBytes(DeviceRead),
		SubCommand:           Int16ToBytes(SubCommand1),
		SoftComponentCode:    MC_3E_BIN_TYPE_D,
		//SoftComponentAddress: addressBlock(address),
		//SoftComponentNumber:  Int16ToBytes(quantity),
		Data:                 dataBlock(address, quantity),
	}
	response,err := mb.send(&request)
	if err != nil{
		return
	}
	length := len(response.Data)
	if length <= 0 {
		err = fmt.Errorf("mc3e:response data size '%v' can not less than 0 ", length)
		return
	}
	results = response.Data[0:]
	//}
	return
}

func (mb *client) send(request *ProtocolDataUnit) (response *ProtocolDataUnit,err error){
	aduRequest, err := mb.packager.Encode(request)
	if err != nil {
		return
	}
	aduResponse, err := mb.transporter.Send(aduRequest)
	if err != nil {
		return
	}
	response, err = mb.packager.Decode(aduResponse)
	if err != nil {
		return
	}

	if response.Command[1] != request.Command[1] {
		err = responseError(response)
		return
	}
	if response.Data == nil || len(response.Data) == 0 {
		// Empty response
		err = fmt.Errorf("modbus: response data is empty")
		return
	}
	return
}


/*func dataBlock(value ...uint16) []byte {
	data := make([]byte, 2*len(value))
	for i, v := range value {
		binary.BigEndian.PutUint16(data[i*2:], v)
	}
	return data
}*/

func dataBlock(value ...uint16) []byte {
	b := make([]uint16,2)
	data := make([]byte, 5)
	for i, v := range value {
		b[i] = v
	}
	data[0] = byte(b[0]>>16)
	data[1] = byte(b[0]>>8)
	data[2] = byte(b[0])
	data[3] = byte(b[1]>>8)
	data[4] = byte(b[1])
	return data
}


func  Int16ToBytes( v uint16) []byte{
	b := make([]byte,2)
	//_ = b[1] // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 8)
	b[1] = byte(v)
	return b
}

func responseError(response *ProtocolDataUnit) error {
	mbError := &Mc3eError{Command: response.Command}
	if response.Data != nil && len(response.Data) > 0 {
		mbError.ExceptionCode =response.Data[0:2]
	}
	return mbError
}