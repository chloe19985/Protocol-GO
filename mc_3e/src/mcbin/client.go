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
func (mb *client) ReadXCoils(address uint16, quantity uint16) (results []byte, err error) {
	if quantity < 0 || quantity > 4095 {
		err = fmt.Errorf("mc_3e-X:quantity '%v' must be between '%v' and '%v'", quantity, 1, 4095)
		return
	}
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_X,
		SoftComponentNumber:  quantityBlock(quantity),
		//Data:                 dataBlock(address, quantity),
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

func (mb *client) ReadDRegisters(address uint16, quantity uint16) (results []byte, err error) {
	if quantity < 0 || quantity > 7998 {
		err = fmt.Errorf("mc_3e_D:quantity '%v' must be between '%v' and '%v'", quantity, 1, 7998)
		return
	}
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand1),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_D,
		SoftComponentNumber:  quantityBlock(quantity),
		//Data:                 dataBlock(address, quantity),
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

/*
func (mb *client) WriteDRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	if quantity < 0 || quantity >7998{
		err = fmt.Errorf("mc_3e_D:quantity '%v' must be between '%v' and '%v'",quantity,1,7998)
		return
	}
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand1),
		SoftComponentCode:    MC_3E_BIN_TYPE_D,
		//SoftComponentAddress: addressBlock(address),
		//SoftComponentNumber:  Uint16ToBytes(quantity),
		Data:         dataBlockSuffix(value, address, quantity),
	}

	_, err = mb.send(&request)
	if err != nil {
		return
	}
	//// Fixed response length
	//if len(response.Data) != 4 {
	//	err = fmt.Errorf("modbus: response data size '%v' does not match expected '%v'", len(response.Data), 4)
	//	return
	//}
	//respValue := binary.BigEndian.Uint16(response.Data)
	//if address != respValue {
	//	err = fmt.Errorf("modbus: response address '%v' does not match request '%v'", respValue, address)
	//	return
	//}
	//
	//results = response.Data[2:]
	//respValue := binary.BigEndian.Uint16(results)
	//if quantity != respValue {
	//	err = fmt.Errorf("modbus: response quantity '%v' does not match request '%v'", respValue, quantity)
	//	return
	//}
	return
}*/

func (mb *client) send(request *ProtocolDataUnit) (response *ProtocolDataUnit, err error) {
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

	if uint16(response.EndCode[1])|uint16(response.EndCode[0])<<8 != 0 {
		err = responseError(response)
		return
	}
	return
}

func responseError(response *ProtocolDataUnit) error {
	mbError := &Mc3eError{EndCode: response.EndCode}
	if response.Data[MC_3E_BIN_ADU_HEADER:] != nil && len(response.Data) > 0 {
		mbError.ExceptionCode = response.EndCode
	}
	return mbError
}
