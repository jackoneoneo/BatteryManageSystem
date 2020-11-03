package communication

import (
	"bytes"
	"encoding/binary"
)

/**
  实现A-BMS与C-BMS通讯的一些结构体
*/
type pack struct {
}

/**
 *  报文转化为字符串
 */
func (pck *pack) PackHead(data []byte) *TcpProtocolMsg {
	msg := &TcpProtocolMsg{}
	dataio := bytes.NewReader(data)
	//：对于data这个字节数组，从低字节开始读，把读到的字节放进msg.len的地址为首的空间，因为这个空间就4个字节（int32）故而会读走4个字节
	//：dataio流的指向在读完之后，会指向data的倒数第5个字节
	binary.Read(dataio, binary.LittleEndian, &msg.FrameStart)
	binary.Read(dataio, binary.LittleEndian, &msg.FrameLen)
	binary.Read(dataio, binary.LittleEndian, &msg.FrameCtrl)
	binary.Read(dataio, binary.LittleEndian, &msg.FrameFlag)
	binary.Read(dataio, binary.LittleEndian, &msg.FrameCsSt)
	binary.Read(dataio, binary.LittleEndian, &msg.FrameBeca)
	binary.Read(dataio, binary.LittleEndian, &msg.FrameAppAddr)
	binary.Read(dataio, binary.LittleEndian, &msg.FrameInfoAddr)
	binary.Read(dataio, binary.LittleEndian, &msg.FrameInfoLen)
	binary.Read(dataio, binary.LittleEndian, &msg.FrameData)

	return msg
}

/**
 * 字符串转化为报文
 */
func (pck *pack) Unpack(msg *TcpProtocolMsg) []byte {
	databufio := bytes.NewBuffer([]byte{})
	binary.Write(databufio, binary.LittleEndian, msg.FrameStart)
	binary.Write(databufio, binary.LittleEndian, msg.FrameLen)
	binary.Write(databufio, binary.LittleEndian, msg.FrameCtrl)
	binary.Write(databufio, binary.LittleEndian, msg.FrameFlag)
	binary.Write(databufio, binary.LittleEndian, msg.FrameCsSt)
	binary.Write(databufio, binary.LittleEndian, msg.FrameBeca)
	binary.Write(databufio, binary.LittleEndian, msg.FrameAppAddr)
	binary.Write(databufio, binary.LittleEndian, msg.FrameInfoAddr)
	binary.Write(databufio, binary.LittleEndian, msg.FrameInfoLen)
	binary.Write(databufio, binary.LittleEndian, msg.FrameData)
	return databufio.Bytes()
}

//C-BMS 返回给A-BMS的数据结构
type TcpProtocolMsg struct {
	FrameStart    uint8    //起始符 0x68
	FrameLen      uint16   //整个数据长度(除了起始符和FrameLen)
	FrameCtrl     uint32   //控制域
	FrameFlag     uint8    //类型标识
	FrameCsSt     uint8    //可变结构限定词
	FrameBeca     uint16   //传送原因
	FrameAppAddr  uint16   //公共地址
	FrameInfoAddr [3]uint8 //信息体地址3个
	FrameInfoLen  uint16   //实际数据长度
	FrameData     []int8   //数据内容
}
