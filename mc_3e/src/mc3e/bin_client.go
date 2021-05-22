package mc3e

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

const (

	MC_3E_BIN_ADU_HEADER          = 9  //副帧头~请求数据长度
	MC_3E_BIN_COMMAND_POSITION    = 11 //指令
	MC_3E_BIN_SUBCOMMAND_POSITION = 13 //子指令
	MC_3E_BIN_REGISTER_POSITION   = 15 //寄存器相关信息长度

	// Default TCP timeout is not set
	tcpTimeout     = 10 * time.Second
	tcpIdleTimeout = 60 * time.Second
)

type BINClientHandler struct {
	binPackager
	binTransporter
}

//
func NewBINClientHandler(address string) *BINClientHandler {
	//&获取变量在计算机内存中的地址
	h := &BINClientHandler{}
	h.Address = address
	h.Timeout = tcpTimeout
	h.IdleTimeout = tcpIdleTimeout
	return h
}

func BINClient(address string) Client {
	handler := NewBINClientHandler(address)
	return NewClient(handler)
}

type binPackager struct {
	// For synchronization between messages of server & client
	transactionId uint32
	// Broadcast address is 0
	SlaveId byte
}


func (mb *binTransporter)Encode(pdu *ProtocolDataUnit) (adu []byte, err error) {
	//var ptr int = 0
	adu = make([]byte, MC_3E_BIN_ADU_HEADER + 2 + 2 + 2 + 1 + 3 + 2)
	binary.BigEndian.PutUint16(adu,Sub_Header)
	//网络编号
	adu[2] = byte(Net_Number)
	//PC编号
	adu[3] = byte(ObjeCt_Number)
	//IO编号
	binary.BigEndian.PutUint16(adu[4:],IO_Number)
	//请求多点站号
	adu[6] = byte(Slave_Number)
	//请求数据长度
	length := uint16(2 + 2 + 2 + 1 + 5)
	binary.LittleEndian.PutUint16(adu[7:], length)
	//保留
	copy(adu[MC_3E_BIN_ADU_HEADER:],pdu.Retain)
	//指令
	copy(adu[MC_3E_BIN_COMMAND_POSITION:], pdu.Command)
	//子指令
	copy(adu[MC_3E_BIN_SUBCOMMAND_POSITION:], pdu.SubCommand)
	//寄存器信息
	copy(adu[MC_3E_BIN_REGISTER_POSITION:] ,pdu.SoftComponentAddress)
	adu[MC_3E_BIN_REGISTER_POSITION+3] = pdu.SoftComponentCode
	copy(adu[MC_3E_BIN_REGISTER_POSITION+4:], pdu.SoftComponentNumber)
	return
}

/*func (mb *binPackager) Verify(aduRequest []byte,aduResponse []byte)(err error){
	//帧头~请求多点站号
	responseVal := binary.BigEndian.Uint16(aduResponse[:7])
	requestVal := binary.BigEndian.Uint16(aduRequest[:7])
	if responseVal != requestVal {
		err = fmt.Errorf("response header '%v' does not match request '%v'", responseVal, requestVal)
		return
	}
	responseVal = binary.BigEndian.Uint16(aduResponse[])

}*/
func (mb *binTransporter)Decode(adu []byte) (pdu *ProtocolDataUnit, err error) {
	length := binary.LittleEndian.Uint16(adu[MC_3E_BIN_ADU_HEADER-2:])
	//fmt.Printf("%d",length)
	pduLength := len(adu) - MC_3E_BIN_ADU_HEADER
	if pduLength <= 0 || pduLength != int(length) {
		err = fmt.Errorf("length in response '%v' does not match pdu data length '%v'", length, pduLength)
		return
	}
	pdu = &ProtocolDataUnit{}
	pdu.EndCode = adu[MC_3E_BIN_ADU_HEADER:]
	pdu.Data = adu[MC_3E_BIN_ADU_HEADER + 2:]
	return
}

type binTransporter struct {
	// Connect string,IP address+port
	Address string
	// Connect & Read timeout
	Timeout time.Duration
	// Idle timeout to close the connection
	IdleTimeout time.Duration
	//传输日志
	Logger *log.Logger

	//	// TCP connection
	mu           sync.Mutex
	conn         net.Conn
	closeTimer   *time.Timer
	lastActivity time.Time
}

func (mb *binTransporter) Connect() error {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	return mb.connect()
}

//建立连接
func (mb *binTransporter) connect() error {
	if mb.conn == nil {
		//限制网络连接时间,Dialer结构体
		dialer := net.Dialer{Timeout: mb.Timeout}
		//Dial支持多种网络连接,返回Dial(network,address)
		conn, err := dialer.Dial("tcp", mb.Address)
		if err != nil {
			return err
		}
		mb.conn = conn
	}
	return nil
}

func (mb *binTransporter) Send(aduRequest []byte) (aduResponse []byte, err error) {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	if err = mb.connect(); err != nil {
		return
	}
	//将定时器设定为空闲时关闭
	mb.lastActivity = time.Now()
	mb.startCloseTimer()
	//设置读写超时
	var timeout time.Time
	if mb.Timeout > 0 {
		timeout = mb.lastActivity.Add(mb.Timeout)
	}
	if err = mb.conn.SetDeadline(timeout); err != nil {
		return
	}
	//发送数据
	mb.logf("sending % X", aduRequest)
	if _, err = mb.conn.Write(aduRequest); err != nil {
		return
	}
	//n,_ := mb.conn.Write(aduRequest)
	//fmt.Printf("n:%d\n",n)
	//fmt.Printf("mb.com:%q\n",mb.conn)
	//n,_ := mb.conn.Write(aduRequest)
	//fmt.Printf("%d\n",n)
	//fmt.Printf("aduRequest:%d\n",aduRequest)
	//读取副帧头~请求数据长度
	var data [MC_3E_MAX_ADU_LENGTH]byte
	//ReadFull从mb.conn中读取len(data)字节到data。即先读取副帧头~请求数据长度
	//n,_ := io.ReadFull(mb.conn, data[:MC_3E_BIN_ADU_HEADER])
	//fmt.Printf("n:%d\n",n)
	if _, err = io.ReadFull(mb.conn, data[:MC_3E_BIN_ADU_HEADER]); err != nil {
		return
	}

	//读取请求数据长度
	length := int(binary.LittleEndian.Uint16(data[MC_3E_BIN_ADU_HEADER-2:]))
	//fmt.Printf("length:%d/n",length)
	if length <= 0 {
		mb.flush(data[:])
		err = fmt.Errorf("length in response header '%v' must not be 0", length)
		return
	}
	if length > (MC_3E_MAX_ADU_LENGTH - (MC_3E_BIN_ADU_HEADER - 1)) {
		mb.flush(data[:])
		err = fmt.Errorf("length in response header '%v' must not greater than '%v'", length, MC_3E_MAX_ADU_LENGTH-MC_3E_BIN_ADU_HEADER+1)
		return
	}
	length = length + MC_3E_BIN_ADU_HEADER
	if _, err = io.ReadFull(mb.conn, data[MC_3E_BIN_ADU_HEADER:length]); err != nil {
		return
	}
	aduResponse = data[:length]
	mb.logf("received % X\n", aduResponse)
	return
}

func (mb *binTransporter) Close() error {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	return mb.close()
}

//断开连接
func (mb *binTransporter) close() (err error) {
	if mb.conn != nil {
		err = mb.conn.Close()
		mb.conn = nil
	}
	return
}

//清空数据流
func (mb binTransporter) flush(b []byte) (err error) {
	if err = mb.conn.SetReadDeadline(time.Now()); err != nil {
		return
	}
	if _, err = mb.conn.Read(b); err != nil {
		// Ignore timeout error
		//&&逻辑与
		if netError, ok := err.(net.Error); ok && netError.Timeout() {
			err = nil
		}
	}
	return
}

func (mb *binTransporter) logf(format string, v ...interface{}) {
	if mb.Logger != nil {
		mb.Logger.Printf(format, v...)
	}
}

func (mb *binTransporter) startCloseTimer() {
	if mb.IdleTimeout <= 0 {
		return
	}
	if mb.closeTimer == nil {
		mb.closeTimer = time.AfterFunc(mb.IdleTimeout, mb.closeIdle)
	} else {
		mb.closeTimer.Reset(mb.IdleTimeout)
	}
}

func (mb *binTransporter) closeIdle() {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	if mb.IdleTimeout <= 0 {
		return
	}
	idle := time.Now().Sub(mb.lastActivity)
	if idle >= mb.IdleTimeout {
		mb.logf("modbus: closing connection due to idle timeout: %v", idle)
		mb.close()
	}
}
