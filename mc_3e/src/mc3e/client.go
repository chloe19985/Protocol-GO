package mc3e

import (
	"encoding/binary"
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

func (mb *client) ReadSCoils(address uint16, quantity uint16) (results []byte, err error) {
	if quantity < 0 || quantity > 4095 {
		err = fmt.Errorf("MC_3E_S:quantity '%v' must be between '%v' and '%v'", quantity, 0, 4095)
		return
	}
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_S,
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

	return
}

func (mb *client) ReadXCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 01777 {
		err = fmt.Errorf("MC_3E_X:quantity '%v' must be between '%v' and '%v'", quantity, 0, 1777)
		return
	}*/
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

	return
}

func (mb *client) ReadYCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 1777 {
		err = fmt.Errorf("MC_3E_Y:quantity '%v' must be between '%v' and '%v'", quantity, 0, 1777)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_Y,
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

	return
}

func (mb *client) ReadMCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity >7679 {
		err = fmt.Errorf("MC_3E_M:quantity '%v' must be between '%v' and '%v'", quantity, 0, 7679)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_M,
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

	return
}

func (mb *client) ReadSMCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 9999 {
		err = fmt.Errorf("MC_3E_SM:quantity '%v' must be between '%v' and '%v'", quantity, 0, 9999)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_SM,
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

	return
}

func (mb *client) ReadLCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 7679 {
		err = fmt.Errorf("MC_3E_L:quantity '%v' must be between '%v' and '%v'", quantity, 0, 7679)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_L,
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

	return
}

func (mb *client) ReadFCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 127 {
		err = fmt.Errorf("MC_3E_F:quantity '%v' must be between '%v' and '%v'", quantity, 0, 127)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_F,
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

	return
}

func (mb *client) ReadBCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 0XFF {
		err = fmt.Errorf("MC_3E_B:quantity '%v' must be between '%v' and '%v'", quantity, 0, 0XFF)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_B,
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

	return
}

func (mb *client) ReadTSCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 511 {
		err = fmt.Errorf("MC_3E_TS:quantity '%v' must be between '%v' and '%v'", quantity, 0, 511)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_TS,
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

	return
}

func (mb *client) ReadTCCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 511 {
		err = fmt.Errorf("MC_3E_TC:quantity '%v' must be between '%v' and '%v'", quantity, 0, 511)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_TC,
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

	return
}

func (mb *client) ReadSSCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 15 {
		err = fmt.Errorf("MC_3E_SS:quantity '%v' must be between '%v' and '%v'", quantity, 0, 15)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_SS,
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

	return
}

func (mb *client) ReadSCCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 15 {
		err = fmt.Errorf("MC_3E_SC:quantity '%v' must be between '%v' and '%v'", quantity, 0, 15)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_SC,
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

	return
}

func (mb *client) ReadCSCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 255 {
		err = fmt.Errorf("MC_3E_CS:quantity '%v' must be between '%v' and '%v'", quantity, 0, 255)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_CS,
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

	return
}

func (mb *client) ReadCCCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 255 {
		err = fmt.Errorf("MC_3E_CC:quantity '%v' must be between '%v' and '%v'", quantity, 0, 255)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_CC,
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

	return
}

func (mb *client) ReadSBCoils(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 0X01FF {
		err = fmt.Errorf("MC_3E_SB:quantity '%v' must be between '%v' and '%v'", quantity, 0, 0X01FF)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand2),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_SB,
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

	return
}

//读寄存器
func (mb *client) ReadDRegisters(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 7999 {
		err = fmt.Errorf("MC_3E_D:quantity '%v' must be between '%v' and '%v'", quantity, 0, 7999)
		return
	}*/
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

	return
}

func (mb *client) ReadSDRegisters(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 11999 {
		err = fmt.Errorf("MC_3E_SD:quantity '%v' must be between '%v' and '%v'", quantity, 0, 11999)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand1),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_SD,
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

	return
}

func (mb *client) ReadWRegisters(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 0X01FF {
		err = fmt.Errorf("MC_3E_SD:quantity '%v' must be between '%v' and '%v'", quantity, 0, 0X01FF)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand1),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_W,
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

	return
}

func (mb *client) ReadTNRegisters(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 511 {
		err = fmt.Errorf("MC_3E_TN:quantity '%v' must be between '%v' and '%v'", quantity, 0, 511)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand1),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_TN,
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

	return
}

func (mb *client) ReadSNRegisters(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 15 {
		err = fmt.Errorf("MC_3E_SN:quantity '%v' must be between '%v' and '%v'", quantity, 0, 15)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand1),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_SN,
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

	return
}

func (mb *client) ReadCNRegisters(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 255 {
		err = fmt.Errorf("MC_3E_CN:quantity '%v' must be between '%v' and '%v'", quantity, 0, 255)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand1),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_CN,
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

	return
}

func (mb *client) ReadSWRegisters(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 0X01FF {
		err = fmt.Errorf("MC_3E_SW:quantity '%v' must be between '%v' and '%v'", quantity, 0, 0X01FF)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand1),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_SW,
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

	return
}

func (mb *client) ReadZRegisters(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 19 {
		err = fmt.Errorf("MC_3E_Z:quantity '%v' must be between '%v' and '%v'", quantity, 0, 19)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand1),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_Z,
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

	return
}

func (mb *client) ReadRRegisters(address uint16, quantity uint16) (results []byte, err error) {
	/*if quantity < 0 || quantity > 32767 {
		err = fmt.Errorf("MC_3E_R:quantity '%v' must be between '%v' and '%v'", quantity, 0, 32767)
		return
	}*/
	request := ProtocolDataUnit{
		Retain:               Uint16ToBytes(Retain_Command),
		Command:              Uint16ToBytes(DeviceRead),
		SubCommand:           Uint16ToBytes(SubCommand1),
		SoftComponentAddress: addressBlock(address),
		SoftComponentCode:    MC_3E_BIN_TYPE_R,
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
	if response.Data[0:] != nil && len(response.Data) > 0 {
		mbError.ExceptionCode = response.EndCode[0:2]
		b := binary.LittleEndian.Uint16(mbError.ExceptionCode)
		binary.BigEndian.PutUint16(mbError.ExceptionCode,b)

	}
	return mbError
}
