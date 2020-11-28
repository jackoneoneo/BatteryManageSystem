package communication

/**
  通讯模块，主要功能是实现A-BMS于C-BMS的实时通讯
*/
import (
	"bytes"
	"fmt"
	"github.com/labstack/gommon/log"
	"net"
	"strings"
	"sync"
	"time"
)

var CBmsSysConfMap = make(map[int]string)         //C-BMS配置缓存
var PerPackageCacheBuffMap = make(map[int][]byte) // 存放C-BMS每包的缓存
var ContentCacheBuffMap = make(map[int]bytes.Buffer)
var ConMap = make(map[string]net.Conn) //保存当前连接
func init() {
	//// 加载系统文件
	//println("加载系统文件")
	//loadSysConfig()
	//CBmsSysConfMap[1] = "127.0.0.1"
	//PerPackageCacheBuffMap[1] = make([]byte, MaxSizePerPacket)
	//ContentCacheBuffMap[1] = bytes.Buffer{}
	////CBmsSysConfMap[2] = "192.168.200.222"
	////PerPackageCacheBuffMap[2]  = make([]byte, MaxSizePerPacket)
	////ContentCacheBuffMap[2] = bytes.Buffer{}

}

//加载系统配置文件
func loadSysConfig() {
	//TODO
}

/**
  常量
*/
const (
	FrameStartFlag           = 0x68
	FrameTypeTransferDataCmd = 0x85
	FrameTypeReceiveData     = 0x82 //BC 采集数据
	FrameTypeReceiveDataCmd  = 0x85 //BC CMD数据
)

/**
 *
 */

func HeartCmd() []byte {
	var msg TcpProtocolMsg
	msg.FrameStart = FrameStartFlag
	msg.FrameLen = 0x15
	msg.FrameCtrl = 0
	msg.FrameFlag = FrameTypeTransferDataCmd
	msg.FrameCsSt = 0x1
	msg.FrameBeca = 0x1 //0x10;循环、周期
	msg.FrameAppAddr = 0
	msg.FrameInfoAddr = [3]uint8{0x1, 0, 0}
	msg.FrameInfoLen = 0x6
	t := time.Now()
	msg.FrameData = []int8{int8(t.Year() - 2000), int8(t.Month()), int8(t.Day()), int8(t.Hour()), int8(t.Minute()), int8(t.Second())}
	bindata := []byte{}
	pck := &pack{}
	bindata = pck.Unpack(&msg)
	for i := 0; i < len(bindata); i++ {
		bindata[i] = bindata[i] + 0x33
	}
	//fmt.Printf("%#x", bindata)
	return bindata
}

/**
  建立与C-BMS通讯的服务器
*/
const (
	port = ":9999"
)

//建立服务器

func CBmsServer() {
	tcpServer, _ := net.ResolveTCPAddr("tcp4", port)
	listener, _ := net.ListenTCP("tcp", tcpServer)
	//统一处理一起的连接
	go handConn()
	for {
		conn, err := listener.Accept()
		if err != nil { //监听异常处理
			log.Info(err)
			continue
		}
		ip := strings.Split(conn.RemoteAddr().String(), ":")[0]
		ConMap[ip] = conn
	}

}

/*
 使用线程来处理多个BC的连接
*/
func handConn() {
	for {
		time.Sleep(time.Duration(5) * time.Millisecond) //防止采集过快
		var wg sync.WaitGroup
		wg.Add(len(CBmsSysConfMap))
		for _, v := range CBmsSysConfMap {
			go handle(v, PerPackageCacheBuffMap[1], ContentCacheBuffMap[1], &wg)
		}
		wg.Wait()
		// TODO 根据C-BMS的实时数据统计A-BMS的实时数据
	}

}

const (
	MaxSizePerPacket     = 2048 //每帧的最大数据量
	MaxDataSizePerPacket = 1024 //数据最大量
	FrameHeadLen         = 18   //报文头长度
)

func handle(key string, buffer []byte, dataBuffer bytes.Buffer, wg *sync.WaitGroup) {
	conn, ok := ConMap[key]
	if !ok {
		//TODO 处理簇不在线的情况，数据清零等操作
		wg.Done()
		return
	}
	defer conn.Close() //发生异常把连接断开

	conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
	n, err := conn.Write(HeartCmd())
	if err != nil {
		delete(ConMap, key)
		println("111")
		println(err)
		//TODO 处理簇不在线的情况，数据清零等操作
		wg.Done()
		return
	}

	n, err = conn.Read(buffer)
	if err != nil {
		delete(ConMap, key)
		println("2222")
		println(err)
		//TODO 处理簇不在线的情况，数据清零等操作
		wg.Done()
		return
	}
	println(n)
	for i := 0; i < n; i++ {
		fmt.Printf("%#x    ", buffer[i])
		//buffer[i] -= 0x33
		//fmt.Printf("%#x    ", buffer[i])
	}
	wg.Done()

	//var packetNo uint8 = 0
	//tempData := FrameAnalysis(buffer[:n], uint16(n), packetNo)
	////dataInfoLen := len(tempData)
	////dataBuffer := bytes.Buffer{}
	//
	//dataBuffer.Write(tempData[6:]) //第一包，需要去掉时间戳

	//for dataInfoLen == MaxDataSizePerPacket { //没有到达最后一包数据
	//	packetNo += 1
	//	TempBuffer := FrameHeartAckSend(packetNo)
	//
	//	n, err = conn.Write(TempBuffer[0:(len(TempBuffer) - 2)])
	//	n, err = conn.Read(buffer)
	//	//println(packetNo ,"   ")
	//
	//	//解密
	//	for i := 0; i < n; i++ {
	//		buffer[i] -= 0x33
	//		//fmt.Printf("%#x    ", buffer[i])
	//	}
	//	tempData = FrameAnalysis(buffer[:n], uint16(n), packetNo)
	//	dataInfoLen = len(tempData)
	//	dataBuffer.Write(tempData)
	//}

	//pck := &cache.PackInfo{}
	//pck.PackCBmsInfo(dataBuffer.Bytes(), 1) //把值放到共享内存中
	//println(cache.CBmsRealTimeMap[1].BMUOnlineState)

}

const (
	FrameInformationType uint8 = 0x81 //表示发送的是C-BMS的实时数据
	FrameCmdType         uint8 = 0x93 //表示发送的为命令数据
)

/**
 * 对报文进行有效性验证
 */
func FrameAnalysis(byteArr []byte, dataLen uint16, currentPacketNo uint8) []byte {

	pck := &pack{}
	msg := pck.PackHead(byteArr)

	if (dataLen < FrameHeadLen-3) || //数据长度小于帧头长度-3（帧标志、帧实际数据长度）
		(msg.FrameStart != FrameStartFlag) ||
		((msg.FrameFlag != FrameTypeReceiveData) && (msg.FrameFlag != FrameTypeReceiveDataCmd)) ||
		(msg.FrameLen != dataLen-3) {
		log.Debug("C-BMS帧格式错误")
		return nil
	}

	infoType := msg.FrameInfoAddr[0]
	//fmt.Printf("%#x", infoType)
	switch infoType {
	case FrameInformationType: //采样数据
		packetNo := msg.FrameInfoAddr[1] //包号
		if currentPacketNo != packetNo { //发送的包号与接受的包号不一致
			log.Debug("当前包号与需要的包号不一致")
			return nil
		}
		dataLen := msg.FrameInfoLen
		if dataLen != (uint16)(len(byteArr[FrameHeadLen:])) { //检查发送的长度与接受的长度是否一致
			println("ddddd")
			return nil
		}
		//msg.FrameData = make([]int8, dataLen)
		//binary.Read(bytes.NewReader(byteArr[FrameHeadLen:]), binary.LittleEndian, &msg.FrameData)
		//fmt.Printf("%#x    ", msg.FrameData[0])
		return byteArr[FrameHeadLen:]

	case FrameCmdType: //命令数据
	default:
		log.Debug("数据类型：", infoType, "无法识别")
		return nil
	}
	return nil
}

/**
 * 发送确认帧，并请求下一包数据
 */
func FrameHeartAckSend(nextPacketNo uint8) []byte {
	var msg TcpProtocolMsg
	msg.FrameStart = FrameStartFlag
	msg.FrameLen = 0xd
	msg.FrameCtrl = 0
	msg.FrameFlag = FrameTypeTransferDataCmd
	msg.FrameCsSt = 0x1
	msg.FrameBeca = 0x1 //0x10;循环、周期
	msg.FrameAppAddr = 0
	msg.FrameInfoAddr = [3]uint8{0x12, nextPacketNo, 0}
	msg.FrameInfoLen = 0
	bindata := []byte{}
	pck := &pack{}
	bindata = pck.Unpack(&msg)
	for i := 0; i < len(bindata); i++ {
		bindata[i] = bindata[i] + 0x33
	}
	//fmt.Printf("%#x", bindata)
	return bindata
}
