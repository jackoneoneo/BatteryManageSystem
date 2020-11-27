package cache

import (
	"bytes"
	"encoding/binary"
	"entity"
)

/**
  初始化共享内存
*/
var CBmsRealTimeMap = make(map[int]entity.CBmsInfo) // C-BMS的实时数据

type PackInfo struct {
}

func (pck *PackInfo) PackCBmsInfo(data []byte, cBmsInd int) {
	msg := CBmsRealTimeMap[cBmsInd]
	dataio := bytes.NewReader(data)
	//：对于data这个字节数组，从低字节开始读，把读到的字节放进msg.len的地址为首的空间，因为这个空间就4个字节（int32）故而会读走4个字节
	//：dataio流的指向在读完之后，会指向data的倒数第5个字节
	binary.Read(dataio, binary.LittleEndian, &msg.BMUOnlineState)
	binary.Read(dataio, binary.LittleEndian, &msg.BMUNum)
	binary.Read(dataio, binary.LittleEndian, &msg.BMUDataLen)
	binary.Read(dataio, binary.LittleEndian, &msg.DMUOnlineState)
	binary.Read(dataio, binary.LittleEndian, &msg.DMUNum)
	binary.Read(dataio, binary.LittleEndian, &msg.DMUDataLen)
	binary.Read(dataio, binary.LittleEndian, &msg.HMUOnlineState)
	binary.Read(dataio, binary.LittleEndian, &msg.HMUNum)
	binary.Read(dataio, binary.LittleEndian, &msg.HMUDataLen)

	//binary.Read(dataio, binary.LittleEndian, &msg.DayInputWatt)
	//binary.Read(dataio, binary.LittleEndian, &msg.DayOutputWatt)
	//binary.Read(dataio, binary.LittleEndian, &msg.PBusResistor)
	//binary.Read(dataio, binary.LittleEndian, &msg.NBusResistor)
	//binary.Read(dataio, binary.LittleEndian, &msg.BranchResistor)
	//binary.Read(dataio, binary.LittleEndian, &msg.NBusVol)
	//binary.Read(dataio, binary.LittleEndian, &msg.AlarmState)
	//binary.Read(dataio, binary.LittleEndian, &msg.RunState)
	//binary.Read(dataio, binary.LittleEndian, &msg.InsHeartbeat)
	//binary.Read(dataio, binary.LittleEndian, &msg.Res)
	//binary.Read(dataio, binary.LittleEndian, &msg.DmuUtcu)
	//binary.Read(dataio, binary.LittleEndian, &msg.DmuRes)

}

func (pck *PackInfo) PackCDmuInfo(data []byte) *entity.CDmuInfo {
	msg := &entity.CDmuInfo{}
	dataio := bytes.NewReader(data)
	//：对于data这个字节数组，从低字节开始读，把读到的字节放进msg.len的地址为首的空间，因为这个空间就4个字节（int32）故而会读走4个字节
	//：dataio流的指向在读完之后，会指向data的倒数第5个字节
	binary.Read(dataio, binary.LittleEndian, &msg.DMUFrameFlag)
	binary.Read(dataio, binary.LittleEndian, &msg.BusVol)
	binary.Read(dataio, binary.LittleEndian, &msg.BusCur)
	binary.Read(dataio, binary.LittleEndian, &msg.SOC)
	binary.Read(dataio, binary.LittleEndian, &msg.Power)
	binary.Read(dataio, binary.LittleEndian, &msg.InputAH)
	binary.Read(dataio, binary.LittleEndian, &msg.OutputAH)
	binary.Read(dataio, binary.LittleEndian, &msg.InputWatt)
	binary.Read(dataio, binary.LittleEndian, &msg.OutputWatt)
	binary.Read(dataio, binary.LittleEndian, &msg.DayInputWatt)
	binary.Read(dataio, binary.LittleEndian, &msg.DayOutputWatt)
	binary.Read(dataio, binary.LittleEndian, &msg.PBusResistor)
	binary.Read(dataio, binary.LittleEndian, &msg.NBusResistor)
	binary.Read(dataio, binary.LittleEndian, &msg.BranchResistor)
	binary.Read(dataio, binary.LittleEndian, &msg.NBusVol)
	binary.Read(dataio, binary.LittleEndian, &msg.AlarmState)
	binary.Read(dataio, binary.LittleEndian, &msg.RunState)
	binary.Read(dataio, binary.LittleEndian, &msg.InsHeartbeat)
	binary.Read(dataio, binary.LittleEndian, &msg.Res)
	binary.Read(dataio, binary.LittleEndian, &msg.DmuUtcu)
	binary.Read(dataio, binary.LittleEndian, &msg.DmuRes)
	return msg
}

func (pck *PackInfo) PackMBmsInfo(data []byte) *entity.MBmsInfo {
	msg := &entity.MBmsInfo{}
	dataio := bytes.NewReader(data)
	//：对于data这个字节数组，从低字节开始读，把读到的字节放进msg.len的地址为首的空间，因为这个空间就4个字节（int32）故而会读走4个字节
	//：dataio流的指向在读完之后，会指向data的倒数第5个字节
	binary.Read(dataio, binary.LittleEndian, &msg.CellNum)
	binary.Read(dataio, binary.LittleEndian, &msg.TempNum)
	binary.Read(dataio, binary.LittleEndian, &msg.BalanceState)
	binary.Read(dataio, binary.LittleEndian, &msg.BalanceState2)
	binary.Read(dataio, binary.LittleEndian, &msg.TotalVol)
	binary.Read(dataio, binary.LittleEndian, &msg.AverageVol)
	binary.Read(dataio, binary.LittleEndian, &msg.LenForBMUExtr)
	binary.Read(dataio, binary.LittleEndian, &msg.BMUFrameFlag)
	binary.Read(dataio, binary.LittleEndian, &msg.CellVol)
	binary.Read(dataio, binary.LittleEndian, &msg.CellTemp)
	binary.Read(dataio, binary.LittleEndian, &msg.CellSoc)
	binary.Read(dataio, binary.LittleEndian, &msg.BMUExtr)
	binary.Read(dataio, binary.LittleEndian, &msg.BmuUtc)
	binary.Read(dataio, binary.LittleEndian, &msg.BMUProtocolFlag)
	return msg
}
func (pck *PackInfo) PackExtremumData(data []byte) *entity.ExtremumData {
	msg := &entity.ExtremumData{}
	dataio := bytes.NewReader(data)
	//：对于data这个字节数组，从低字节开始读，把读到的字节放进msg.len的地址为首的空间，因为这个空间就4个字节（int32）故而会读走4个字节
	//：dataio流的指向在读完之后，会指向data的倒数第5个字节
	binary.Read(dataio, binary.LittleEndian, &msg.TMaxInd)
	binary.Read(dataio, binary.LittleEndian, &msg.TMinInd)
	binary.Read(dataio, binary.LittleEndian, &msg.VMaxInd)
	binary.Read(dataio, binary.LittleEndian, &msg.VMinInd)
	binary.Read(dataio, binary.LittleEndian, &msg.SMaxInd)
	binary.Read(dataio, binary.LittleEndian, &msg.SMinInd)
	binary.Read(dataio, binary.LittleEndian, &msg.FillData1)
	binary.Read(dataio, binary.LittleEndian, &msg.FillData2)
	binary.Read(dataio, binary.LittleEndian, &msg.TempMax)
	binary.Read(dataio, binary.LittleEndian, &msg.TempMin)
	binary.Read(dataio, binary.LittleEndian, &msg.VolMax)
	binary.Read(dataio, binary.LittleEndian, &msg.VolMin)
	binary.Read(dataio, binary.LittleEndian, &msg.SOCMax)
	binary.Read(dataio, binary.LittleEndian, &msg.SOCMin)
	return msg
}
func (pck *PackInfo) PackHvm(data []byte) *entity.CHmuInfo {
	msg := &entity.CHmuInfo{}
	dataio := bytes.NewReader(data)
	//：对于data这个字节数组，从低字节开始读，把读到的字节放进msg.len的地址为首的空间，因为这个空间就4个字节（int32）故而会读走4个字节
	//：dataio流的指向在读完之后，会指向data的倒数第5个字节
	binary.Read(dataio, binary.LittleEndian, &msg.HMUFrameFlag)
	binary.Read(dataio, binary.LittleEndian, &msg.HMUCellTemp)
	binary.Read(dataio, binary.LittleEndian, &msg.TempState)
	binary.Read(dataio, binary.LittleEndian, &msg.FanState)
	binary.Read(dataio, binary.LittleEndian, &msg.FanRunState)
	binary.Read(dataio, binary.LittleEndian, &msg.FuseState)
	binary.Read(dataio, binary.LittleEndian, &msg.SelfState)
	binary.Read(dataio, binary.LittleEndian, &msg.HMUIOState)
	binary.Read(dataio, binary.LittleEndian, &msg.HMUDOData)
	binary.Read(dataio, binary.LittleEndian, &msg.YWKAircondOnlineState)
	binary.Read(dataio, binary.LittleEndian, &msg.SensorOnlineState)
	binary.Read(dataio, binary.LittleEndian, &msg.MeterOnlineState)
	binary.Read(dataio, binary.LittleEndian, &msg.HDAircondOnlineState)
	binary.Read(dataio, binary.LittleEndian, &msg.HmuUtc)
	binary.Read(dataio, binary.LittleEndian, &msg.HmuRes)
	return msg
}
