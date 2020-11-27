package entity

const (
	CBmsBmuNum  = 20
	CBmsHmuNum  = 5
	HmuTempNum  = 3
	MBmsCellNum = 20
	MBmsTempNum = 20
)

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
