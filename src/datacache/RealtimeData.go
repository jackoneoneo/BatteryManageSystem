package datacache

import (
	"bytes"
	"encoding/binary"
	"sync"
)

/***
该模块主要的功能是存放设备的实时数据，相当于数据的共享内存
用来保存C-BMS、M-BMS、HVM、A-BMS的数据
*/
const (
	CBmsBmuNum  = 20
	CBmsHmuNum  = 5
	HmuTempNum  = 3
	MBmsCellNum = 20
	MBmsTempNum = 20
	MaxCBmsSize = 4
)

//BMS的数据点
type BMSCacheHandler struct {
	RwMutex          sync.RWMutex          //读写锁
	CBmsRealTimeData [MaxCBmsSize]CBmsZone //C-BMS数据点
	ABmsRealTimeData ABmsZone              // A-BMS
}

//获取C-BMS的数据的句柄
func GetCBmsRealTimeDataHandle(id int8) *CBmsZone {
	cache.BMSCache.RwMutex.RLock()
	data := &cache.BMSCache.CBmsRealTimeData[id-1]
	cache.BMSCache.RwMutex.RUnlock()
	return data
}

//获取A-BMS的数据
func GetABmsRealTimeDataHandle() *ABmsZone {
	cache.BMSCache.RwMutex.RLock()
	data := &cache.BMSCache.ABmsRealTimeData
	cache.BMSCache.RwMutex.RUnlock()
	return data
}

type CacheHandler struct {
	BMSCache BMSCacheHandler
}

type PackInfo struct {
}

/*共享内存初始化*/
var cache CacheHandler

func init() {

}

/**
  计算A-BMS的数据
*/
func CalcABmsData() {

}

func (pck *PackInfo) PackCBmsInfo(data []byte, cBmsInd int8) {

	//msg := cache.CBmsRealTimeData[cBmsInd].info
	//msg := Cache.BMSCache.CBmsRealTimeData[1].info
	//msg.BMUOnlineState = binary.LittleEndian.Uint32(data[0:4])
	//dataio := bytes.NewReader(data)
	//：对于data这个字节数组，从低字节开始读，把读到的字节放进msg.len的地址为首的空间，因为这个空间就4个字节（int32）故而会读走4个字节
	//：dataio流的指向在读完之后，会指向data的倒数第5个字节
	//binary.Read(dataio, binary.LittleEndian, &msg.BMUOnlineState)
	//binary.Read(dataio, binary.LittleEndian, &msg.BMUNum)
	//binary.Read(dataio, binary.LittleEndian, &msg.BMUDataLen)
	//binary.Read(dataio, binary.LittleEndian, &msg.DMUOnlineState)
	//binary.Read(dataio, binary.LittleEndian, &msg.DMUNum)
	//binary.Read(dataio, binary.LittleEndian, &msg.DMUDataLen)
	//binary.Read(dataio, binary.LittleEndian, &msg.HMUOnlineState)
	//binary.Read(dataio, binary.LittleEndian, &msg.HMUNum)
	//binary.Read(dataio, binary.LittleEndian, &msg.HMUDataLen)

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

func (pck *PackInfo) PackCDmuInfo(data []byte) *CDmuInfo {
	msg := &CDmuInfo{}
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

func (pck *PackInfo) PackMBmsInfo(data []byte) *MBmsInfo {
	msg := &MBmsInfo{}
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
func (pck *PackInfo) PackExtremumData(data []byte) ExtremumData {
	msg := ExtremumData{}
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
func (pck *PackInfo) PackHvm(data []byte) CHmuInfo {
	msg := CHmuInfo{}
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

/*A-BMS的数据区*/
type ABmsZone struct {
	Id             [32]int8            /* 节点标识*/
	Nbatc          uint8               /* 堆序号 */
	SysTime        [32]uint8           /* 时间标识 */
	StartIp        [32]int8            /* 节点标识*/
	StartAdr       [4]int8             /* 4字IP */
	EndIp          [32]int8            /* 节点标识*/
	BcIp           [MaxCBmsSize]string /* BCMS的IP存放数组 */
	EndAdr         [4]int8             /* 4字IP */
	SysIsolStatus  uint8               /* 系统绝缘状态 */
	BmsIsolStatus  uint8               /* BMS绝缘状态 */
	BatfIsolStatus [MaxCBmsSize]uint8  /* 电池簇绝缘状态 */
	BatfNum        uint16              /* 电池组数量 */
	RealBatfNum    uint16              /* 实际所接电池组数量 */
	BmuNum         uint16              /* 电池组下BMU数量 */
	CellsNum       uint16              /* BMU下电池数 */
	TempNum        uint16              /* BMU下温度点数 */
	DmuNum         uint16              /* BMS下DMU计数 */
	AcuNum         uint16              /* BMS下DMU计数 */
	RatedAmp       uint16              /* 单簇额定电流,倍率100 */
	RatedVolt      uint16              /* 单体额定电流,倍率1000 */
	Capacity       uint32              /* 额定容量 kwh */
	BatfMap        uint32              /* batf加载map表 */
	RebootFlag     uint8               /* 复位标志 */
	PcsVol         uint16              /* PCS 测量电压 倍率 10 */
	PcsCur         int16               /* PCS 测量电流 倍率 10 */
	PcsMvol        uint16              /* PCS 母线电压 倍率 10 */
	PcsFault       uint32              /* PCS故障码*/
	Status         uint16              /* 装置状态 */
	RFU            uint16              /* 填充预留*/
	version        string              /* 版本 */
	ProductName    string              /*产品名称*/
	BmsAAAFlag     uint32              /*紧急情况下设置成AAAA	 0x1234表示需要设置，0x0000表示不需要设置*/
	StandardStep   uint16
	StackNum       uint32
	info           ABmsInfo
}
type ABmsInfo struct {
	VMax        uint16 /* 单体最大电压, 倍率1000 V*/
	VMin        uint16 /* 单体最小电压, 倍率1000 V */
	VMaxCellInd uint16 /* 单体最大电压包内索引 */
	VMinCellInd uint16 /* 单体最大电压包内索引 */
	VMaxBmuInd  uint16 /* 单体最大电压包索引 */
	VMinBmuInd  uint16 /* 单体最小电压包索引 */
	VMaxBatfInd uint16 /* 单体最大电压簇索引 */
	VMinBatfInd uint16 /* 单体最小电压簇索引 */

	TMax        int16  /* 单体最大温度, 倍率10 */
	TMin        int16  /* 单体最小温度, 倍率10 */
	TMaxCellInd uint16 /* 单体最大电压包内索引 */
	TMinCellInd uint16 /* 单体最小电压包内索引 */
	TMaxBmuInd  uint16 /* 单体最大电压包索引 */
	TMinBmuInd  uint16 /* 单体最小电压包索引 */
	TMaxBatfInd uint16 /* 单体最大电压簇索引 */
	TMinBatfInd uint16 /* 单体最小电压簇索引 */

	SocMax    uint16 /* soc最大值,倍率10 */
	SocMin    uint16 /* soc最小值,倍率10 */
	SocMaxInd uint16 /* soc最大值, 簇索引*/
	SocMinInd uint16 /* soc最小值, 簇索引*/

	VDif   uint16 /* 单体最大压差, 倍率1000 */
	TDif   uint16 /* 单体最大温差, 倍率10 */
	SocDif uint16 /* soc最大差值,倍率10 */
	RFU1   uint16 /* 保留，对齐用 */

	SOC  uint16 /* 电池堆SOC 倍率10 */
	SOH  uint16 /* 电池堆SOH 倍率10 */
	DOD  uint16 /* 电池堆DOD 倍率10 */
	VAvg uint16 /* 电池堆平均单体电压 倍率1000 V */

	Current   int16   /* 电池堆电流 倍率 10A */
	CurAvg    uint16  /* 合闸的簇的平均电流的绝对值倍率100 A */
	Voltage   uint16  /* 电池堆电压 倍率 10  V*/
	Power     int16   /* 功率kw 倍率1 */
	Power_F32 float32 /* 功率kw 倍率1 */
	ChgState  uint16  /* 充放电状态: 0：静置; 1：充电; 2：放电 */

	ChargeCurMax    uint16 /* 最大允许充电电流 倍率 10 */
	DischargeCurMax uint16 /* 最大允许放电电流 倍率 10 */
	ChargeVolMax    uint16 /* 最大允许充电限压值 倍率 10 */
	DischargeVolMin uint16 /* 最小允许放电限压值 倍率 10 */

	LeftEnergy uint32 /* 可放电量 kwh 倍率1000 */
	UsedEnergy uint32 /* 可充电量 kwh 倍率1000 */

	InputEnergy  uint32 /* 累计充电电量 倍率100*/
	OutputEnergy uint32 /* 累计放电电量 倍率100*/

	InputAH  uint32 /* 累计充电AH 0.1AH*/
	OutputAH uint32 /* 累计放电AH 0.1AH*/

	WarnState            uint32 /* 告警状态字 */
	PrtState             uint32 /* 保护状态字 */
	FaultStatus          uint32 /* 系统故障状态字 0正常  1故障 bit0 温度采样 其他预留*/
	InsulationStatusByBa uint32 /* bams判断的绝缘故障状态字*/
	BalanceEnableState   uint32 /* 堆合算的簇均衡使能状态，一位代表一簇 */
	SysState             uint16 /* BMS系统状态，玉门项目 */
	AvgTemp              int16  /* BMS平均温度，玉门项目 */
	Reserve              uint16 /* 内存对齐，保留 */
	OnNum                uint16 /* batf合闸数量 */
	OnMap                uint32 /* batf合闸映射表 */
	OnlineMap            uint32 /* batf在线表 */
	TempCrolState        uint16 /*储能箱温控系统是否OK*/
	FireState            uint16 /*储能箱的消防情况*/
	BreakValue           uint32
	FaultMap             uint32 /* batf合闸故障映射表 */
	EnableInsuFlag       uint16
	PcsOnline            uint16
	TotalBreakState      uint16
}
type CBmsZone struct {
	Id                   [64]int8 /* 节点编号 */
	SysTime              [5]uint8 /* 时间标识 */
	Ip                   [32]int8 /* 节点Ip号 */
	BmuNum               uint16   /* BCMS下BMU数 */
	CellNum              uint16   /* BMU下单体电池数 */
	TempNum              uint16   /* BMU下温度点数 */
	RatedAmp             uint16   /* 单簇额定电流*/
	RatedVolt            uint16   /* 单体额定电流 */
	Index                uint16   /* 簇索引，0开始 */
	Channel              uint16   /* 簇所在通道 [1-8] */
	Stack                uint16
	IsOnline             uint16 /*bcms通信在线状态:( 0(离线) 1--(在线) */
	Life                 uint32 /* 数据刷新标志 */
	BmuMap               uint32 /* BMU加载map表 */
	Ver                  string /* BCMS版本号 */
	SOH                  uint16
	VAvg                 uint16 /*簇平均单体电压,1000 V*/
	AvgTemp              int16  /*簇平均单体温度*/
	TotalEnergy          uint32 /* 簇最大电量 kwh 倍率1000 */
	WarnState            uint32 /* 告警状态字 */
	PrtState             uint32 /* 保护状态字 */
	SelfWarnState        uint32 /* bams自身判定的告警状态字 */
	SelfPrtState         uint32 /* bams自身判定的保护状态字 */
	LastBreakState       uint32 /* 之前的分合闸状态*/
	LastDmuComEventState uint16 /*记录上一次DMU通信事件状态 0--恢复 1--发生*/
	info                 CBmsInfo
}

/*C-BMS的数据点*/
type CBmsInfo struct {
	BMUOnlineState  uint32               /*BMU在线状态*/
	BMUNum          uint16               /*BMU在线个数*/
	BMUDataLen      uint16               /*单个BMU数据长度*/
	DMUOnlineState  uint16               /*DMU在线状态*/
	DMUNum          uint16               /*DMU个数*/
	DMUDataLen      uint16               /*单个DMU数据长度*/
	HMUOnlineState  uint16               /*HMU在线状态*/
	HMUNum          uint16               /*HMU个数*/
	HMUDataLen      uint16               /*单个HMU数据长度*/
	BusData         CDmuInfo             /*簇DMU数据*/
	CellData        [CBmsBmuNum]MBmsInfo /*BMU数据*/
	HMUData         [CBmsHmuNum]CHmuInfo /*HMU数据*/
	BalEnable       uint32               /*均衡使能状态 */
	BalanceState    uint32               /*均衡状态字*/
	SmpFuseState    uint32               /* BMU采样保险丝状态 每个bit代表一个BMU是否有保险丝断掉 */
	BlFuseState     uint32               /* BMU均衡保险丝状态 每个bit代表一个BMU是否有保险丝断掉 */
	BcmsWarnState   uint32               /* bcms自身判定的告警状态字 */
	BcmsPrtState    uint32               /* bcms自身判定的保护状态字 */
	BcmsPrtActState uint32               /* bcms自身动作状态字 */
	BcmsSysState    uint32               /* bcms系统状态字 */
	VMax            uint16               /* 簇单体最大电压, 倍率1000 */
	VMin            uint16               /* 簇单体最小电压, 倍率1000*/
	TMax            int16                /* 簇单体最大温度, 倍率10 */
	TMin            int16                /* 簇单体最小温度, 倍率10 */
	SocMax          uint16               /* 簇内单体最大SOC */
	SocMin          uint16               /* 簇内单体最小SOC*/
	VMaxBmuInd      uint8                /* 簇单体最大电压包索引 */
	VMaxCellInd     uint8                /* 簇单体最大电压包内索引 */
	VMinBmuInd      uint8                /* 簇单体最小电压包索引 */
	VMinCellInd     uint8                /* 簇单体最大电压包内索引 */
	TMaxBmuInd      uint8                /* 簇单体最大电压包索引 */
	TMaxCellInd     uint8                /* 簇单体最大电压包内索引 */
	TMinBmuInd      uint8                /* 簇单体最小电压包索引 */
	TMinCellInd     uint8                /* 簇单体最小电压包内索引 */
	BmuVAvg         uint16               /*簇内BMU平均电压(各BMU总压(单体电压之和)求平均),1000 V*/
	BreakControl    uint8                /* 电池簇断路器控制状态 */
	BreakFault      uint8                /* 电池簇断路器故障 */
	LeftEnergy      uint32               /* 可放电量 kwh 倍率100 */
	UsedEnergy      uint32               /* 可充电量 kwh 倍率100 */
	SetRealAH       uint16               /*设置AH*/
	BcVersion       [8]uint8             /*BC版本号*/
	BcmsFaultStatus uint32               /*故障状态字 0正常  1故障 bit0 温度采样 其他预留 */
	BcRes           [60]uint8            /*BC预留*/
}

/*HVM数据点*/
type CDmuInfo struct {
	DMUFrameFlag   uint16    /* DMU主动上报的帧 */
	BusVol         uint16    /* 母线电压,mv */
	BusCur         int16     /* 母线电流,ma */
	SOC            uint16    /* SOC 0.1% */
	Power          int32     /* 瞬时功率 0.1KW*/
	InputAH        uint32    /* 累计充电AH 0.1AH */
	OutputAH       uint32    /* 累计放电AH 0.1AH */
	InputWatt      uint32    /* 累计充电电量 0.1KWh*/
	OutputWatt     uint32    /* 累计放电电量0.1KWh */
	DayInputWatt   uint32    /* 日累计充电电量 0.1KWh*/
	DayOutputWatt  uint32    /* 日累计放电电量 0.1KWh */
	PBusResistor   uint16    /* 正对地电阻   */
	NBusResistor   uint16    /* 负对地电阻 */
	BranchResistor uint16    /* 支路绝缘电阻 */
	NBusVol        uint16    /* 负母线电压*/
	AlarmState     uint16    /* 告警状态字 */
	RunState       uint16    /* 运行状态字 */
	InsHeartbeat   uint8     //绝缘检测心跳位
	Res            [3]uint8  //保留，对齐
	DmuUtcu        int32     /*时间*/
	DmuRes         [16]uint8 //预留
}

/*M-BMS的数据点*/
type MBmsInfo struct {
	CellNum         uint8               /*电池单体个数*/
	TempNum         uint8               /*温度个数*/
	BalanceState    uint16              /*均衡状态 0%*/
	BalanceState2   uint16              /*均衡状态 0%*/
	TotalVol        uint16              /*BMU的总电压 mV(单体电压之和)*/
	AverageVol      uint16              /*单体平均电压mV*/
	LenForBMUExtr   uint16              /*BMU的极值长度*/
	BMUFrameFlag    uint32              /*主动上报的帧标志*/
	CellVol         [MBmsCellNum]uint16 /* 单体电压 mv*/
	CellTemp        [MBmsTempNum]int16  /* 单体温度 0.1度*/
	CellSoc         [MBmsCellNum]uint16 /* 单体SOC */
	BMUExtr         ExtremumData        /*BMU的极值*/
	BmuUtc          uint32              /*时间*/
	BMUProtocolFlag uint32              /*协议帧类型*/
}

type ExtremumData struct {
	TMaxInd   uint8 /* 最高温度编号 */
	TMinInd   uint8 /* 最低温度编号 */
	VMaxInd   uint8 /* 最高电压编号 */
	VMinInd   uint8 /* 最低电压编号 */
	SMaxInd   uint8 /* 最高SOC编号 */
	SMinInd   uint8 /* 最低SOC编号 */
	FillData1 uint8 /*填充数据*/
	FillData2 uint8
	TempMax   uint8 /* 最高温度 变比10 单位 度*/
	TempMin   uint8 /* 最低温度 变比10 单位 度*/
	VolMax    uint8 /* 最高电压 mv*/
	VolMin    uint8 /* 最低电压 mv*/
	SOCMax    uint8 /* 最高SOC */
	SOCMin    uint8 /* 最低SOC */
}

/*HMU的数据点*/
type CHmuInfo struct {
	HMUFrameFlag          uint16             /*HMU主动上报的帧 	*/
	HMUCellTemp           [HmuTempNum]uint16 /*温度摄氏度		*/
	TempState             uint8              /*温度状态			*/
	FanState              uint8              //风扇异常状态
	FanRunState           uint8              /*风扇运行状态		*/
	FuseState             uint8              //熔断器状态
	SelfState             uint16             //自检状态
	HMUIOState            uint32             /*HMU的DI状态		*/
	HMUDOData             uint16             //HMU的DO数据
	YWKAircondOnlineState uint8              //英维克空调在线状态
	SensorOnlineState     uint8              //传感器在线状态
	MeterOnlineState      uint8              //表计在线状态
	HDAircondOnlineState  uint8              //黑盾空调在线状态
	HmuUtc                uint32             //HMU时间
	HmuRes                [16]uint8          //预留
}
