package simrecord

type GetRecordIdsReq struct {
	ScenId  string `json:"scen_id"`
	ScenVer string `json:"scen_ver"`
}
type GetRecordIdsRes struct {
	Ids []string `json:"ids"`
}

type GetTrackResultsReq struct {
	Id    string `json:"id"`
	ObjId string `json:"obj_id"`
}
type GetTrackResultsRes struct {
	Data []*TrackResult `json:"data"`
}

type GetSensorResultsReq struct {
	Id    string `json:"id"`
	ObjId string `json:"obj_id"`
}
type GetSensorResultsRes struct {
	Data []*SensorResult `json:"data"`
}

type GetStepResultsReq struct {
	Id    string `json:"id"`
	ObjId string `json:"obj_id"`
}
type GetStepResultsRes struct {
	Data []*StepResult `json:"data"`
}

type GetPathResultsReq struct {
	Id    string `json:"id"`
	ObjId string `json:"obj_id"`
}
type GetPathResultsRes struct {
	Data []*PathResult `json:"data"`
}

type GetReferenceLineResultsReq struct {
	Id    string `json:"id"`
	ObjId string `json:"obj_id"`
}
type GetReferenceLineResultsRes struct {
	Data []*ReferenceLineResult `json:"data"`
}

type TrackResult struct {
	// 记录的业务ID
	RecordId string `json:"record_id"`
	// 对象ID
	ObjId string `json:"obj_id"`
	// 时间戳
	Timestamp int32 `json:"timestamp"`
	// 轨迹信息
	Result *Track `json:"result"`
}

type Track struct {
	X   float32 `json:"x"`
	Y   float32 `json:"y"`
	Z   float32 `json:"z"`
	Phi float32 `json:"phi"`
	// 车道ID，允许为空
	LaneId       string `json:"lane_id"`
	PositionType string `json:"position_type"`
}

type SensorResult struct {
	// 记录的业务ID
	RecordId string `json:"record_id"`
	// 对象ID
	ObjId string `json:"obj_id"`
	// 时间戳
	Timestamp int32 `json:"timestamp"`
	// 传感器信息
	Result []*SensorObj `json:"result"`
}

type SensorObj struct {
	Id     string  `json:"id"`
	Speed  float32 `json:"speed"`
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
	Z      float32 `json:"z"`
	Length float32 `json:"length"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Phi    float32 `json:"phi"`
	// 最低位起置1表示灯光点亮:近光灯(0) 远光灯(1) 左转向灯(2) 右转向灯(3)
	// 紧急报警灯(4) 刹车灯(5) e.g. "000001" 近光灯; "100000" 刹车灯; "111111"
	// 全亮
	ExteriorLight string `json:"exterior_light"`
	// 0(无风险); 1(低风险); 2(高风险)
	Risk_2Ego int32 `json:"risk_2_ego"`
	// 纵向加速度
	LonAcc float32 `json:"lon_acc"`
	// 横向速度
	V float32 `json:"v"`
	// 横向加速度
	LatAcc float32 `json:"lat_acc"`
	// 横摆角速度
	W float32 `json:"w"`
	// 横摆角加速度
	WAcc float32 `json:"w_acc"`
	// 车道ID，允许为空
	LaneId string `json:"lane_id"`
	// POSITION_TYPE_UNKNOWN = 0;
	// 1. 在车道内
	// POSITION_TYPE_IN_LANE = 1;
	// 2. 在路口内
	// POSITION_TYPE_IN_JUNCTION = 2;
	// 3. 在道路外
	// POSITION_TYPE_OUT_ROAD = 3;
	PositionType string `json:"position_type"`
}

type StepResult struct {
	// 记录的业务ID
	RecordId string `json:"record_id"`
	// 对象ID
	ObjId string `json:"obj_id"`
	// 时间戳
	Timestamp int32 `json:"timestamp"`
	// 瞬时结果信息
	Result *Step `json:"result"`
}

type Step struct {
	// 速度
	Speed float32 `json:"speed"`
	// 加速度
	Acc float32 `json:"acc"`
	// 里程
	Mileage float32 `json:"mileage"`
	// 方向盘转角
	SteWheel float32 `json:"ste_wheel"`
	// 转向灯
	TurnSignal string `json:"turn_signal"`
	// 横向速度
	V float32 `json:"v"`
	// 横向加速度
	LatAcc float32 `json:"lat_acc"`
	// 横摆角速度
	W float32 `json:"w"`
	// 横摆角加速度
	WAcc float32 `json:"w_acc"`
	// 期望速度
	ReferenceSpeed float32 `json:"reference_speed"`
}

type PathResult struct {
	// 记录的业务ID
	RecordId string `json:"record_id"`
	// 对象ID
	ObjId string `json:"obj_id"`
	// 时间戳
	Timestamp int32 `json:"timestamp"`
	// 瞬时结果信息
	Result []*Path `json:"result"`
}

type Path struct {
	Points []*PathPoint `json:"points"`
}

type PathPoint struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type ReferenceLineResult struct {
	// 记录的业务ID
	RecordId string `json:"record_id"`
	// 对象ID
	ObjId string `json:"obj_id"`
	// 时间戳
	Timestamp int32 `json:"timestamp"`
	// 瞬时结果信息
	Result []*ReferenceLine `json:"result"`
}

type ReferenceLine struct {
	Points    []*PathPoint `json:"points"`
	LineIds   []string     `json:"line_ids"`
	LineTypes []string     `json:"line_types"`
	LineIdxs  []int32      `json:"line_idxs"`
	// 是否逆行
	Opposite bool `json:"opposite"`
}
