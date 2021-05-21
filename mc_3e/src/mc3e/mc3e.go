package mc3e

import (
	"encoding/binary"
	"fmt"
)

//FX5U各区寄存器
const(
	MC_3E_BIN_TYPE_S  = 0x98 //步进继电器 位 输入地址为10进制
	MC_3E_BIN_TYPE_X  = 0x9C //输入继电器 位 输入地址为8进制
	MC_3E_BIN_TYPE_Y  = 0x9D //输出继电器 位 输入地址为8进制
	MC_3E_BIN_TYPE_M  = 0x90 //内部继电器 位 输入地址为10进制
	MC_3E_BIN_TYPE_SM = 0x91 //特殊继电器 位 输入地址为10进制
	MC_3E_BIN_TYPE_L  = 0x92 //所存继电器 位 输入地址为10进制
	MC_3E_BIN_TYPE_F  = 0x93 //报警器 位 输入地址为10进制
	//MC_3E_BIN_TYPE_V  = 0x94 //边沿继电器 位 输入地址为10进制
	MC_3E_BIN_TYPE_B  = 0xA0 //链接继电器 位 输入地址为10进制
	MC_3E_BIN_TYPE_TS = 0xC1 //定时器触点 位 输入地址为10进制
	MC_3E_BIN_TYPE_TC = 0xC0 //定时器线圈 位 输入地址为10进制
	MC_3E_BIN_TYPE_SS = 0xC7 //累计定时器触点 位 输入地址为10进制
	MC_3E_BIN_TYPE_SC = 0xC6 //累计定时器线圈 位 输入地址为10进制
	MC_3E_BIN_TYPE_CS = 0xC4 //计时器触点 位 输入地址为10进制
	MC_3E_BIN_TYPE_CC = 0xC3 //计时器线圈 位 输入地址为10进制
	MC_3E_BIN_TYPE_SB = 0xA1 //链接特殊继电器 位 输入地址为16进制
	//MC_3E_BIN_TYPE_DX    =    //直接输入 位 输入地址为16进制
	MC_3E_BIN_TYPE_D  = 0xA8 //数据寄存器 字 输入地址为16进制
	MC_3E_BIN_TYPE_SD = 0xA9 //特殊寄存器 字 输入地址为10进制
	MC_3E_BIN_TYPE_W  = 0xB4 //链接寄存器 字 输入地址为16进制
	MC_3E_BIN_TYPE_TN = 0xC2 //定时器当前值 字 输入地址为10进制
	MC_3E_BIN_TYPE_SN = 0xC8 //累计定时器当前值 字 输入地址为10进制
	MC_3E_BIN_TYPE_CN = 0xC5 //计数器当前值 字 输入地址为10进制
	MC_3E_BIN_TYPE_SW = 0xB5 //链接特殊继电器 字 输入地址为16进制
	MC_3E_BIN_TYPE_Z  = 0xCC //变址寄存器 字 输入地址为10进制
	MC_3E_BIN_TYPE_R  = 0xAF //文件寄存器 字 输入地址为10进制
	//MC_3E_BIN_TYPE_ZR = 0xB0 //文件寄存器 字 输入地址为16进制
	//MC_3E_BIN_TYPE_END
)

//固定码
const (
	Sub_Header uint16 = 0x5000
	Net_Number = 0x00
	ObjeCt_Number = 0xFF
	IO_Number = 0xFF03
	Slave_Number = 0x00
	Retain_Command uint16 = 0x0000
	MC_3E_MAX_ADU_LENGTH = 512
)

//指令
const (
	/*
		批量读取
		位：以一点为单位读取位软元件
		字：以一点为单位读取字软元件
		   以十六点为单位读取位软元件
		批量写入：同上
	*/
	DeviceRead uint16 = 0x0104
	DeviceWrite uint16  = 0x0114
)

//子指令
const (
	SubCommand1 uint16 = 0x0000
	SubCommand2 uint16 = 0x0100
)


//错误码
const (
	ExceptionCode1 uint16  = 0xC035
	ExceptionCode2 uint16 = 0xC050
	ExceptionCode3 uint16 = 0xC051
	ExceptionCode4 uint16 = 0xC052
	ExceptionCode5 uint16 = 0xC053
	ExceptionCode6 uint16 = 0xC054
	ExceptionCode7 uint16 = 0xC056
	ExceptionCode8 uint16 = 0xC058
	ExceptionCode9 uint16 = 0xC059
	ExceptionCode10 uint16 = 0xC05B
	ExceptionCode11 uint16 = 0xC05C
	ExceptionCode12 uint16 = 0xC05F
	ExceptionCode13 uint16 = 0xC060
	ExceptionCode14 uint16 = 0xC061
	ExceptionCode15 uint16 = 0xC06F
	ExceptionCode16 uint16 = 0xC0D8
	ExceptionCode17 uint16 = 0xC200
	ExceptionCode18 uint16 = 0xC201
	ExceptionCode19 uint16 = 0xC204
	ExceptionCode20 uint16 = 0xC810
	ExceptionCode21 uint16 = 0xC815
	ExceptionCode22 uint16 = 0xC816
)


//定义一个Mc3eError结构体
type Mc3eError struct {
	EndCode       []byte
	ExceptionCode []byte
}

//格式：func(名称 类型) 方法名（参数列表）（返回值列表）
//方法：与函数区别为有接收者（从属），引用e.xx
//Error为方法名
//参数为空，返回数据类型为string
func (e *Mc3eError) Error() string {
	var name string
	switch binary.BigEndian.Uint16(e.ExceptionCode) {
	case ExceptionCode1:
		name = "响应监视定时器值以内无法执行对象设备的生存确认"
	case ExceptionCode2:
		name = "通信数据代码设置为ASCII时，接收了无法转换为二进制码的ASCII码数据"
	case ExceptionCode3:
		name = "可一次性批量读写的最大位软元件数超出容许范围"
	case ExceptionCode4:
		name = "可一次性批量读写的最大字软元件数超出容许范围"
	case ExceptionCode5:
		name = "可一次性随机读写的最大位软元件数超出容许范围"
	case ExceptionCode6:
		name = "可一次性随机读写的最大字软元件数超出容许范围"
	case ExceptionCode7:
		name = "超出最大地址的写入及读取请求"
	case ExceptionCode8:
		name = "ASCII-二进制转换后的请求数据长度与字符区（文本的一部分）的数据不符"
	case ExceptionCode9:
		name = "命令、子命令的指定有误；CPU模块中无法使用的命令、子命令"
	case ExceptionCode10:
		name = "CPU模块无法对指定软元件进行写入及读取"
	case ExceptionCode11:
		name = "请求内容有误（以位为单位对字软元件进行写入、读取等）"
	case ExceptionCode12:
		name = "是无法对对象CPU模块执行的请求"
	case ExceptionCode13:
		name = "请求内容有误（对位软元件的数据指定指定有误）"
	case ExceptionCode14:
		name = "请求数据长度与字符区(文本的一部分)的数据数不符"
	case ExceptionCode15:
		name = "通信数据代码被设置为二进制时，接收了ASCII的请求报文(本出错代码仅登录出错履历，而不返回异常响应)"
	case ExceptionCode16:
		name = "指定块数超过范围"
	case ExceptionCode17:
		name = "远程口令有误"
	case ExceptionCode18:
		name = "通信所使用的端口处于远程密码锁定状态"
	case ExceptionCode19:
		name = "与请求了远程口令解锁处理的对方设备不同"
	case ExceptionCode20:
		name = "远程密码有误(认证失败次数为9次以下)"
	case ExceptionCode21:
		name = "远程密码有误(认证失败次数为10次)"
	case ExceptionCode22:
		name = "远程密码认证闭锁中"
	default:
		name = "未知错误"
	}
	return fmt.Sprintf("mc3e:错误码'%x'(%s),function '%x'", e.ExceptionCode, name, e.EndCode)
}

/*
保留 2
指令 2
子指令 2
软元件代码 1
数据：
起始元件编号 3
软元件点数 2
*/
type ProtocolDataUnit struct {
	Retain               []byte
	Command              []byte
	SubCommand           []byte
	SoftComponentAddress []byte
	SoftComponentCode    byte
	SoftComponentNumber  []byte
	EndCode              []byte
	Data                 []byte
}

type Packager interface {
	Encode(pdu *ProtocolDataUnit) (adu []byte, err error)
	Decode(adu []byte) (pdu *ProtocolDataUnit, err error)
	//Verify(aduRequest []byte, aduResponse []byte) (err error)
}
type Transporter interface {
	Send(aduRequest []byte) (aduResponse []byte, err error)
}
