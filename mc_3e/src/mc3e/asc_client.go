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
	//MC_3E_ASC_SUB_HEADER uint16 = 0x5000 //副帧头
	//MC_3E_ASC_NET_NUMBER        = 0x00   //请求目标网络编号
	//MC_3E_ASC_OBJECT            = 0xFF   //请求目标站号
	//MC_3E_ASC_IO_NUMBER  uint16 = 0x03FF //请求目标模板IO编号
	//MC_3E_ASC_SLAVE             = 0x00   //请求多占站号

	MC_3E_ASC_ADU_HEADER          = 18 //副帧头~请求数据长度
	MC_3E_ASC_COMMAND_POSITION    = 22 //指令
	MC_3E_ASC_SUBCOMMAND_POSITION = 26 //子指令
	MC_3E_ASC_REGISTER_CODE       = 30 //寄存器相关信息长度
	MC_3E_ASC_REGISTER_ADDRESS    = 32
	MC_3E_ASC_REGISTER_NUMBER     = 38
	//MC_3E_MAX_ADU_LENGTH = 512 //报文最大长度
	//
	//// Default TCP timeout is not set
	//tcpTimeout     = 10 * time.Second
	//tcpIdleTimeout = 60 * time.Second
)

type ASCClientHandler struct {
	ascPackager
	ascTransporter
}

//
func NewASCClientHandler(address string) *ASCClientHandler {
	//&获取变量在计算机内存中的地址
	h := &ASCClientHandler{}
	h.Address = address
	h.Timeout = tcpTimeout
	h.IdleTimeout = tcpIdleTimeout
	return h
}

func ASCClient(address string) Client {
	handler := NewASCClientHandler(address)
	return NewClient(handler)
}

type ascPackager struct {
	// For synchronization between messages of server & client
	transactionId uint32
	// Broadcast address is 0
	SlaveId byte
}


func (mb *ascTransporter)Encode(pdu *ProtocolDataUnit) (adu []byte, err error) {
	//var ptr int = 0
	adu = make([]byte, MC_3E_ASC_ADU_HEADER + 4 + 4 + 4 + 2 + 6 + 4)

	McSetBinToChar(adu,0,Sub_Header>>8)
	McSetBinToChar(adu,2,Sub_Header)
	McSetBinToChar(adu,4,Net_Number)
	McSetBinToChar(adu,6,ObjeCt_Number)
	McSetBinToChar(adu,8,IO_Number )
	McSetBinToChar(adu,10,IO_Number >> 8)
	McSetBinToChar(adu,12,Slave_Number)

	//请求数据长度
	length := uint16(4 + 4 + 4 + 2 + 6 + 4)
	McSetBinToChar(adu,14,length>>8)
	McSetBinToChar(adu,16,length)

	//保留
	McSetBinToChar(adu,18,binary.BigEndian.Uint16(pdu.Retain)>>8)
	McSetBinToChar(adu,20,binary.BigEndian.Uint16(pdu.Retain))
	//指令
	McSetBinToChar(adu,MC_3E_ASC_COMMAND_POSITION,binary.BigEndian.Uint16(pdu.Command))
	McSetBinToChar(adu,MC_3E_ASC_COMMAND_POSITION + 2,binary.BigEndian.Uint16(pdu.Command) >> 8 )
	//copy(adu[MC_3E_BIN_COMMAND_POSITION:], pdu.Command)
	//子指令
	McSetBinToChar(adu,MC_3E_ASC_SUBCOMMAND_POSITION,binary.BigEndian.Uint16(pdu.SubCommand))
	McSetBinToChar(adu,MC_3E_ASC_SUBCOMMAND_POSITION + 2,binary.BigEndian.Uint16(pdu.SubCommand)>>8)
	//寄存器信息
	McSetBinToChar(adu, MC_3E_ASC_REGISTER_CODE,uint16(pdu.SoftComponentCode))
	McSetBinToChar(adu, MC_3E_ASC_REGISTER_ADDRESS ,uint16(pdu.SoftComponentAddress[2]))
	McSetBinToChar(adu,MC_3E_ASC_REGISTER_ADDRESS +2,uint16(pdu.SoftComponentAddress[1]))
	McSetBinToChar(adu,MC_3E_ASC_REGISTER_ADDRESS + 4,uint16(pdu.SoftComponentAddress[0]))
	McSetBinToChar(adu, MC_3E_ASC_REGISTER_NUMBER,binary.BigEndian.Uint16(pdu.SoftComponentNumber))
	McSetBinToChar(adu,MC_3E_ASC_REGISTER_NUMBER+2,binary.BigEndian.Uint16(pdu.SoftComponentNumber)>>8)
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
func (mb *ascTransporter)Decode(adu []byte) (pdu *ProtocolDataUnit, err error) {
	length := binary.BigEndian.Uint32(adu[MC_3E_ASC_ADU_HEADER-4:])
	//fmt.Printf("%d",length)
	pduLength := len(adu) - MC_3E_ASC_ADU_HEADER
	if pduLength <= 0 || pduLength != int(length) {
		err = fmt.Errorf("length in response '%v' does not match pdu data length '%v'", length, pduLength)
		return
	}
	pdu = &ProtocolDataUnit{}
	pdu.EndCode = adu[MC_3E_ASC_ADU_HEADER:]
	pdu.Data = adu[MC_3E_ASC_ADU_HEADER + 4:]
	return
}

type ascTransporter struct {
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

func (mb *ascTransporter) Connect() error {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	return mb.connect()
}

//建立连接
func (mb *ascTransporter) connect() error {
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

func (mb *ascTransporter) Send(aduRequest []byte) (aduResponse []byte, err error) {
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
	if _, err = io.ReadFull(mb.conn, data[:MC_3E_ASC_ADU_HEADER]); err != nil {
		return
	}

	//读取请求数据长度
	length := int(binary.LittleEndian.Uint16(data[MC_3E_ASC_ADU_HEADER-2:]))
	//fmt.Printf("length:%d/n",length)
	if length <= 0 {
		mb.flush(data[:])
		err = fmt.Errorf("length in response header '%v' must not be 0", length)
		return
	}
	if length > (MC_3E_MAX_ADU_LENGTH - (MC_3E_ASC_ADU_HEADER - 1)) {
		mb.flush(data[:])
		err = fmt.Errorf("length in response header '%v' must not greater than '%v'", length, MC_3E_MAX_ADU_LENGTH-MC_3E_BIN_ADU_HEADER+1)
		return
	}
	length = length + MC_3E_ASC_ADU_HEADER
	if _, err = io.ReadFull(mb.conn, data[MC_3E_ASC_ADU_HEADER:length]); err != nil {
		return
	}
	aduResponse = data[:length]
	mb.logf("received % X\n", aduResponse)
	return
}

func (mb *ascTransporter) Close() error {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	return mb.close()
}

//断开连接
func (mb *ascTransporter) close() (err error) {
	if mb.conn != nil {
		err = mb.conn.Close()
		mb.conn = nil
	}
	return
}

//清空数据流
func (mb ascTransporter) flush(b []byte) (err error) {
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

func (mb *ascTransporter) logf(format string, v ...interface{}) {
	if mb.Logger != nil {
		mb.Logger.Printf(format, v...)
	}
}

func (mb *ascTransporter) startCloseTimer() {
	if mb.IdleTimeout <= 0 {
		return
	}
	if mb.closeTimer == nil {
		mb.closeTimer = time.AfterFunc(mb.IdleTimeout, mb.closeIdle)
	} else {
		mb.closeTimer.Reset(mb.IdleTimeout)
	}
}

func (mb *ascTransporter) closeIdle() {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	if mb.IdleTimeout <= 0 {
		return
	}
	idle := time.Now().Sub(mb.lastActivity)
	if idle >= mb.IdleTimeout {
		mb.logf("MC_3E: closing connection due to idle timeout: %v", idle)
		mb.close()
	}
}
