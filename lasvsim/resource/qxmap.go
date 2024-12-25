package resource

// qxmap 部分
type Qxmap struct {
	// 地图id
	Id string `json:"id"`
	// 数字签名
	Digest string `json:"digest"`
	// 地图头信息
	Header *Header `json:"header"`
	// 路口组
	Junctions []*Junction `json:"junctions"`
	// 路段组
	Segments []*Segment `json:"segments"`
	// 道路组
	Roads []*Road `json:"roads"`
	// 映射关系
	Mappers []*Mapper `json:"mappers"`
	// 建筑物
	Buildings []*Building `json:"buildings"`
	// 树
	Trees []*Tree `json:"trees"`
	// 路灯
	Lamps []*Lamp `json:"lamps"`
}

type Header struct {
	North        float64           `json:"north"`
	South        float64           `json:"south"`
	East         float64           `json:"east"`
	West         float64           `json:"west"`
	CenterPoint  *Point            `json:"center_point"`
	Version      string            `json:"version"`
	Zone         float64           `json:"zone"`
	UseBias      bool              `json:"use_bias"`
	Source       Header_SourceType `json:"source"`
	CheckVersion string            `json:"check_version"`
}

type Header_SourceType int32

const (
	Header_SOURCE_QXMAP       Header_SourceType = 0
	Header_SOURCE_APOLLO      Header_SourceType = 1
	Header_SOURCE_OPENDRIVE   Header_SourceType = 2
	Header_SOURCE_ROAD_EDITOR Header_SourceType = 3
)

// Enum value maps for Header_SourceType.
var (
	Header_SourceType_name = map[int32]string{
		0: "SOURCE_QXMAP",
		1: "SOURCE_APOLLO",
		2: "SOURCE_OPENDRIVE",
		3: "SOURCE_ROAD_EDITOR",
	}
	Header_SourceType_value = map[string]int32{
		"SOURCE_QXMAP":       0,
		"SOURCE_APOLLO":      1,
		"SOURCE_OPENDRIVE":   2,
		"SOURCE_ROAD_EDITOR": 3,
	}
)

type Junction struct {
	// id
	Id string `json:"id"`
	// 地图 id
	MapId string `json:"map_id"`
	// 名称
	Name string                `json:"name"`
	Type Junction_JunctionType `json:"type"`
	// 路口形状点位
	Shape *Polygon `json:"shape"`
	// 入路口道路 id
	UpstreamSegmentIds []string `json:"upstream_segment_ids"`
	// 出路口道路 id
	DownstreamSegmentIds []string `json:"downstream_segment_ids"`
	// 流向
	Movements []*Movement `json:"movements"`
	// 连接
	Connections []*Connection `json:"connections"`
	// 人行横道
	Crosswalks []*Crosswalk `json:"crosswalks"`
	// 待行区 link组
	WaitAreas []*Link `json:"wait_areas"`
	// 环岛 link组
	Roundabout []*Link `json:"roundabout"`
	// 路口 link组
	Links []*Link `json:"links"`
	// 信号灯方案
	SignalPlan *SignalPlan `json:"signal_plan"`
}

type Junction_JunctionType int32

const (
	Junction_JUNCTION_TYPE_UNKNOWN Junction_JunctionType = 0
	// 断头路
	Junction_JUNCTION_TYPE_DEAD_END Junction_JunctionType = 1
	// 交叉口
	Junction_JUNCTION_TYPE_CROSSING Junction_JunctionType = 2
	// 环岛
	Junction_JUNCTION_TYPE_ROUNDABOUT Junction_JunctionType = 3
	// 匝道入口
	Junction_JUNCTION_TYPE_RAMP_IN Junction_JunctionType = 4
	// 匝道出口
	Junction_JUNCTION_TYPE_RAMP_OUT Junction_JunctionType = 5
	// 虚拟路口
	Junction_JUNCTION_TYPE_VIRTUAL Junction_JunctionType = 6
)

// Enum value maps for Junction_JunctionType.
var (
	Junction_JunctionType_name = map[int32]string{
		0: "JUNCTION_TYPE_UNKNOWN",
		1: "JUNCTION_TYPE_DEAD_END",
		2: "JUNCTION_TYPE_CROSSING",
		3: "JUNCTION_TYPE_ROUNDABOUT",
		4: "JUNCTION_TYPE_RAMP_IN",
		5: "JUNCTION_TYPE_RAMP_OUT",
		6: "JUNCTION_TYPE_VIRTUAL",
	}
	Junction_JunctionType_value = map[string]int32{
		"JUNCTION_TYPE_UNKNOWN":    0,
		"JUNCTION_TYPE_DEAD_END":   1,
		"JUNCTION_TYPE_CROSSING":   2,
		"JUNCTION_TYPE_ROUNDABOUT": 3,
		"JUNCTION_TYPE_RAMP_IN":    4,
		"JUNCTION_TYPE_RAMP_OUT":   5,
		"JUNCTION_TYPE_VIRTUAL":    6,
	}
)

// 多边形
type Polygon struct {
	Points []*Point `json:"points"`
}
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Movement struct {
	Id string `json:"id"`
	// 入路口link id
	UpstreamLinkId string `json:"upstream_link_id"`
	// 出路口link id
	DownstreamLinkId string `json:"downstream_link_id"`
	// junction id
	JunctionId string `json:"junction_id"`
	// 流向
	FlowDirection Direction `json:"flow_direction"`
}

type Direction int32

const (
	Direction_DIRECTION_UNKNOWN Direction = 0
	// 直行
	Direction_DIRECTION_STRAIGHT Direction = 1
	// 左转
	Direction_DIRECTION_LEFT Direction = 2
	// 右转
	Direction_DIRECTION_RIGHT Direction = 3
	// 掉头
	Direction_DIRECTION_U_TURN Direction = 4
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "DIRECTION_UNKNOWN",
		1: "DIRECTION_STRAIGHT",
		2: "DIRECTION_LEFT",
		3: "DIRECTION_RIGHT",
		4: "DIRECTION_U_TURN",
	}
	Direction_value = map[string]int32{
		"DIRECTION_UNKNOWN":  0,
		"DIRECTION_STRAIGHT": 1,
		"DIRECTION_LEFT":     2,
		"DIRECTION_RIGHT":    3,
		"DIRECTION_U_TURN":   4,
	}
)

type Connection struct {
	Id string `json:"id"`
	// 所属路口 id
	JunctionId string `json:"junction_id"`
	// 所属 lane id
	MovementId string `json:"movement_id"`
	// 入路口 lane id
	UpstreamLaneId string `json:"upstream_lane_id"`
	// 出路口 lane id
	DownstreamLaneId string `json:"downstream_lane_id"`
	// 流向
	FlowDirection Direction `json:"flow_direction"`
	// 入路口 link id
	UpstreamLinkId string `json:"upstream_link_id"`
	// 出路口 link id
	DownstreamLinkId string `json:"downstream_link_id"`
	// 车辆行驶路径
	Path *LineString `json:"path"`
}

type LineString struct {
	Points []*Point `json:"points"`
}

type Crosswalk struct {
	// 人行横道 id
	Id string `json:"id"`
	// 人行横道 方向
	Heading float64 `json:"heading"`
	// 点位 封闭多点构成
	Shape *Polygon `json:"shape"`
}

type Link struct {

	// id
	Id string `json:"id"`
	// 地图 id
	MapId string `json:"map_id"`
	// 名称
	Name string        `json:"name"`
	Type Link_LinkType `json:"type"`
	// 关联的 Link
	PairIds []string `json:"pair_ids"`
	// 宽度
	Widths []*Width `json:"widths"`
	// 车道组
	OrderedLanes []*Lane `json:"ordered_lanes"`
	// 长度
	Length float64 `json:"length"`
	// 相对于所属Segment起点的偏移；针对路段内link，路口内s_offset为0
	SOffset float64 `json:"s_offset"`
	// 排序
	LinkNum int32 `json:"link_num"`
	// 所属Segment / junction
	ParentId string `json:"parent_id"`
	// 参考线
	ReferenceLine []*ReferencePoint `json:"reference_line"`
	// 上游Link
	UpstreamLinkIds []string `json:"upstream_link_ids"`
	// 下游Link
	DownstreamLinkIds []string `json:"downstream_link_ids"`
	// 左边界
	LeftBoundary *LineString `json:"left_boundary"`
	// 右边界
	RightBoundary *LineString `json:"right_boundary"`
}
type Link_LinkType int32

const (
	Link_LINK_TYPE_UNKNOWN Link_LinkType = 0
	// 等待区
	Link_LINK_TYPE_WAIT_AREA Link_LinkType = 1
	// 环岛
	Link_LINK_TYPE_ROUNDABOUT Link_LinkType = 2
	// 路段
	Link_LINK_TYPE_SEGMENT Link_LinkType = 3
	// 路口
	Link_LINK_TYPE_JUNCTION Link_LinkType = 4
	// 提前右转
	Link_LINK_TYPE_PRE_TURN_RIGHT Link_LinkType = 5
	// 提前掉头
	Link_LINK_TYPE_PRE_U_TURN Link_LinkType = 6
)

// Enum value maps for Link_LinkType.
var (
	Link_LinkType_name = map[int32]string{
		0: "LINK_TYPE_UNKNOWN",
		1: "LINK_TYPE_WAIT_AREA",
		2: "LINK_TYPE_ROUNDABOUT",
		3: "LINK_TYPE_SEGMENT",
		4: "LINK_TYPE_JUNCTION",
		5: "LINK_TYPE_PRE_TURN_RIGHT",
		6: "LINK_TYPE_PRE_U_TURN",
	}
	Link_LinkType_value = map[string]int32{
		"LINK_TYPE_UNKNOWN":        0,
		"LINK_TYPE_WAIT_AREA":      1,
		"LINK_TYPE_ROUNDABOUT":     2,
		"LINK_TYPE_SEGMENT":        3,
		"LINK_TYPE_JUNCTION":       4,
		"LINK_TYPE_PRE_TURN_RIGHT": 5,
		"LINK_TYPE_PRE_U_TURN":     6,
	}
)

type Width struct {
	S float64 `json:"s"` // 车道距离起始点偏离距离
	A float64 `json:"a"` // 多项式参数 a, 在@s (ds=0)处偏移
	B float64 `json:"b"` // 多项式参数 b
	C float64 `json:"c"` // 多项式参数 c
	D float64 `json:"d"` // 多项式参数 d
}

type Lane struct {
	Id string `json:"id"`
	// 车道类型
	Type Lane_LaneType `json:"type"`
	// 车道偏移
	LaneNum int32 `json:"lane_num"`
	// 车道所属 link
	LinkId string `json:"link_id"`
	// 车道转向标志
	LaneTurn *LaneTurn `json:"lane_turn"`
	// 车道中的多段限速属性
	SpeedLimits []*SpeedLimit `json:"speed_limits"`
	// 停车线id
	Stopline *Stopline `json:"stopline"`
	// 车道宽度
	Widths []*Width `json:"widths"`
	// 车道中心线
	CenterLine []*CenterPoint `json:"center_line"`
	// 上游车道id
	UpstreamLaneIds []string `json:"upstream_lane_ids"`
	// 下游车道id
	DownstreamLaneIds []string `json:"downstream_lane_ids"`
	// 长度
	Length float64 `json:"length"`
	// 左车道线
	LeftLaneMarks []*LaneMark `json:"left_lane_marks"`
	// 右车道线
	RightLaneMarks []*LaneMark `json:"right_lane_marks"`
	// 左车道线点位
	LeftBoundary *LineString `json:"left_boundary"`
	// 右车道线点位
	RightBoundary *LineString `json:"right_boundary"`
	// 车道宽度（该车道宽度中）
	Width float64 `json:"width"`
}

type Lane_LaneType int32

const (
	Lane_LANE_TYPE_UNKNOWN  Lane_LaneType = 0  // 未知
	Lane_LANE_TYPE_DRIVING  Lane_LaneType = 1  // 普通机动车车道
	Lane_LANE_TYPE_BIKING   Lane_LaneType = 2  // 非机动车道
	Lane_LANE_TYPE_SIDEWALK Lane_LaneType = 3  // 人行道
	Lane_LANE_TYPE_PARKING  Lane_LaneType = 4  // 停车区
	Lane_LANE_TYPE_BORDER   Lane_LaneType = 5  // 边界线
	Lane_LANE_TYPE_MEDIAN   Lane_LaneType = 6  // 分隔带
	Lane_LANE_TYPE_BUSING   Lane_LaneType = 7  // 公交车道
	Lane_LANE_TYPE_CURB     Lane_LaneType = 8  // 路沿
	Lane_LANE_TYPE_ENTRY    Lane_LaneType = 10 // 加速车道进入段 todo 确认是否需要
	Lane_LANE_TYPE_EXIT     Lane_LaneType = 11 // 加速车道退出段
	Lane_LANE_TYPE_RAMP_IN  Lane_LaneType = 12 // 闸道车道进入段
	Lane_LANE_TYPE_RAMP_OUT Lane_LaneType = 13 // 闸道车道退出段
)

// Enum value maps for Lane_LaneType.
var (
	Lane_LaneType_name = map[int32]string{
		0:  "LANE_TYPE_UNKNOWN",
		1:  "LANE_TYPE_DRIVING",
		2:  "LANE_TYPE_BIKING",
		3:  "LANE_TYPE_SIDEWALK",
		4:  "LANE_TYPE_PARKING",
		5:  "LANE_TYPE_BORDER",
		6:  "LANE_TYPE_MEDIAN",
		7:  "LANE_TYPE_BUSING",
		8:  "LANE_TYPE_CURB",
		10: "LANE_TYPE_ENTRY",
		11: "LANE_TYPE_EXIT",
		12: "LANE_TYPE_RAMP_IN",
		13: "LANE_TYPE_RAMP_OUT",
	}
	Lane_LaneType_value = map[string]int32{
		"LANE_TYPE_UNKNOWN":  0,
		"LANE_TYPE_DRIVING":  1,
		"LANE_TYPE_BIKING":   2,
		"LANE_TYPE_SIDEWALK": 3,
		"LANE_TYPE_PARKING":  4,
		"LANE_TYPE_BORDER":   5,
		"LANE_TYPE_MEDIAN":   6,
		"LANE_TYPE_BUSING":   7,
		"LANE_TYPE_CURB":     8,
		"LANE_TYPE_ENTRY":    10,
		"LANE_TYPE_EXIT":     11,
		"LANE_TYPE_RAMP_IN":  12,
		"LANE_TYPE_RAMP_OUT": 13,
	}
)

type LaneTurn struct {
	// 带有3维朝向的点坐标
	Position *DirectedPoint `json:"position"`
	// 掉头、左转、直行和右转(TLSR)的组合(0否|1是)
	TurnCode string `json:"turn_code"`
}

type DirectedPoint struct {
	// 三维坐标点
	Point *Point `json:"point"`
	// 朝向
	Heading float64 `json:"heading"`
	// 横滚角
	Roll float64 `json:"roll"`
	// 俯仰角
	Patch float64 `json:"patch"`
}

type SpeedLimit struct {
	// 车道距离起始点偏离距离
	S float64 `json:"s"`
	// 长度
	Length float64                   `json:"length"`
	Type   SpeedLimit_SpeedLimitType `json:"type"`
	// 限速最大值
	MaxValue float64 `json:"max_value"`
	// 限速最小值
	MinValue float64 `json:"min_value"`
	// 速度单位
	Unit string `json:"unit"`
	// 限速源
	Source string `json:"source"`
}

type SpeedLimit_SpeedLimitType int32

const (
	// 无限制
	SpeedLimit_SPEED_LIMIT_UNLIMITED SpeedLimit_SpeedLimitType = 0
	// 限制
	SpeedLimit_SPEED_LIMIT_LIMITED SpeedLimit_SpeedLimitType = 1
	// 限制最大值
	SpeedLimit_SPEED_LIMIT_MAX_LIMITED SpeedLimit_SpeedLimitType = 2
	// 限制最小值
	SpeedLimit_SPEED_LIMIT_MIN_LIMITED SpeedLimit_SpeedLimitType = 3
)

// Enum value maps for SpeedLimit_SpeedLimitType.
var (
	SpeedLimit_SpeedLimitType_name = map[int32]string{
		0: "SPEED_LIMIT_UNLIMITED",
		1: "SPEED_LIMIT_LIMITED",
		2: "SPEED_LIMIT_MAX_LIMITED",
		3: "SPEED_LIMIT_MIN_LIMITED",
	}
	SpeedLimit_SpeedLimitType_value = map[string]int32{
		"SPEED_LIMIT_UNLIMITED":   0,
		"SPEED_LIMIT_LIMITED":     1,
		"SPEED_LIMIT_MAX_LIMITED": 2,
		"SPEED_LIMIT_MIN_LIMITED": 3,
	}
)

type Stopline struct {
	Id    string      `json:"id"`
	Shape *LineString `json:"shape"`
}

type LaneMark struct {
	// 距离车道起始点偏离距离
	S float64 `json:"s"`
	// 长度
	Length float64 `json:"length"`
	// 是否合并; true：该车道线被两个车道共用；false：该车道线独立存在
	IsMerge bool `json:"is_merge"`
	// 车道分界线样式
	Style LaneMark_LaneMarkStyle `json:"style"`
	// 车道线颜色
	Color LaneMark_LaneMarkColor `json:"color"`
	// 车道线宽度
	Width float64 `json:"width"`
	// 车道分界线样式数组
	Styles []LaneMark_LaneMarkStyle `json:"styles"`
	// 车道线颜色
	Colors []LaneMark_LaneMarkColor `json:"colors"`
}

type LaneMark_LaneMarkStyle int32

const (
	LaneMark_LANE_MARK_STYLE_UNKNOWN LaneMark_LaneMarkStyle = 0
	// 无
	LaneMark_LANE_MARK_STYLE_NONE LaneMark_LaneMarkStyle = 1
	// 实线
	LaneMark_LANE_MARK_STYLE_SOLID LaneMark_LaneMarkStyle = 2
	// 虚线
	LaneMark_LANE_MARK_STYLE_BROKEN LaneMark_LaneMarkStyle = 3
	// 双实线
	LaneMark_LANE_MARK_STYLE_DOUBLE_SOLID LaneMark_LaneMarkStyle = 4
	// 双虚线
	LaneMark_LANE_MARK_STYLE_DOUBLE_BROKEN LaneMark_LaneMarkStyle = 5
)

// Enum value maps for LaneMark_LaneMarkStyle.
var (
	LaneMark_LaneMarkStyle_name = map[int32]string{
		0: "LANE_MARK_STYLE_UNKNOWN",
		1: "LANE_MARK_STYLE_NONE",
		2: "LANE_MARK_STYLE_SOLID",
		3: "LANE_MARK_STYLE_BROKEN",
		4: "LANE_MARK_STYLE_DOUBLE_SOLID",
		5: "LANE_MARK_STYLE_DOUBLE_BROKEN",
	}
	LaneMark_LaneMarkStyle_value = map[string]int32{
		"LANE_MARK_STYLE_UNKNOWN":       0,
		"LANE_MARK_STYLE_NONE":          1,
		"LANE_MARK_STYLE_SOLID":         2,
		"LANE_MARK_STYLE_BROKEN":        3,
		"LANE_MARK_STYLE_DOUBLE_SOLID":  4,
		"LANE_MARK_STYLE_DOUBLE_BROKEN": 5,
	}
)

type LaneMark_LaneMarkColor int32

const (
	LaneMark_LANE_MARK_COLOR_UNKNOWN LaneMark_LaneMarkColor = 0
	// 白色
	LaneMark_LANE_MARK_COLOR_WHITE LaneMark_LaneMarkColor = 1
	// 黄色
	LaneMark_LANE_MARK_COLOR_YELLOW LaneMark_LaneMarkColor = 2
)

// Enum value maps for LaneMark_LaneMarkColor.
var (
	LaneMark_LaneMarkColor_name = map[int32]string{
		0: "LANE_MARK_COLOR_UNKNOWN",
		1: "LANE_MARK_COLOR_WHITE",
		2: "LANE_MARK_COLOR_YELLOW",
	}
	LaneMark_LaneMarkColor_value = map[string]int32{
		"LANE_MARK_COLOR_UNKNOWN": 0,
		"LANE_MARK_COLOR_WHITE":   1,
		"LANE_MARK_COLOR_YELLOW":  2,
	}
)

type CenterPoint struct {
	// 距离起始点偏离距离
	S float64 `json:"s"`
	// 朝向
	Heading float64 `json:"heading"`
	// 距离左侧边界宽度
	LeftWidth float64 `json:"left_width"`
	// 距离右侧边界宽度
	RightWidth float64 `json:"right_width"`
	//	double dis_to_end = 4;
	//
	// 三维坐标点
	Point *Point `json:"point"`
}

type SignalPlan struct {
	// 信号灯id
	Id string `json:"id"`
	// 路口 id
	JunctionId string `json:"junction_id"`
	// 信号灯周期
	Cycle int32 `json:"cycle"`
	// 信号灯偏移
	Offset int32 `json:"offset"`
	// 黄闪
	IsYellow bool `json:"is_yellow"`
	// 流向信号组
	MovementSignals map[string]*SignalPlan_MovementSignal `json:"movement_signals"`
}

type SignalPlan_MovementSignal struct {
	// movement id
	MovementId string `json:"movement_id"`
	// 交通灯杆 id 数组
	TrafficLightPoleId string `json:"traffic_light_pole_id"`
	// 信号灯杆位置
	Position *DirectedPoint `json:"position"`
	// 绿灯信号组
	SignalOfGreens []*SignalPlan_MovementSignal_SignalOfGreen `json:"signal_of_greens"`
}

// 绿灯信号
type SignalPlan_MovementSignal_SignalOfGreen struct {
	// 绿灯开始时间
	GreenStart int32 `json:"green_start"`
	// 绿灯持续时间
	GreenDuration int32 `json:"green_duration"`
	// 黄灯时间
	Yellow int32 `json:"yellow"`
	// 车辆清空红灯等待时间
	RedClean int32 `json:"red_clean"`
}

type Segment struct {
	// id
	Id string `json:"id"`
	// 地图 id
	MapId string `json:"map_id"`
	// 名称
	Name string `json:"name"`
	// 始发路口id
	StartJunctionId string `json:"start_junction_id"`
	// 终点路口id
	EndJunctionId string `json:"end_junction_id"`
	// 配对 Segment id
	PairSegmentIds []string `json:"pair_segment_ids"`
	// 相对于起点 s
	SOffset float64 `json:"s_offset"`
	// 所属 road
	RoadId string `json:"road_id"`
	// 长度
	Length float64 `json:"length"`
	// link 组
	OrderedLinks []*Link `json:"ordered_links"`
	// 交通标志
	TrafficSigns []*TrafficSign `json:"traffic_signs"`
}

type TrafficSign struct {
	Id string `json:"id"`
	// 类型
	Type TrafficSign_TrafficSignType `json:"type"`
	// 位置 三维有向点
	Position *DirectedPoint `json:"position"`
}

type TrafficSign_TrafficSignType int32

const (
	TrafficSign_SIGN_TYPE_UNKNOWN TrafficSign_TrafficSignType = 0
	// 限速
	TrafficSign_SIGN_TYPE_SPEED_LIMIT TrafficSign_TrafficSignType = 1
	// 指示牌
	TrafficSign_SIGN_TYPE_GUIDE_SIGN TrafficSign_TrafficSignType = 2
	// 信号灯杆
	TrafficSign_SIGN_TYPE_TRAFFIC_LIGHT_POLE TrafficSign_TrafficSignType = 3
)

// Enum value maps for TrafficSign_TrafficSignType.
var (
	TrafficSign_TrafficSignType_name = map[int32]string{
		0: "SIGN_TYPE_UNKNOWN",
		1: "SIGN_TYPE_SPEED_LIMIT",
		2: "SIGN_TYPE_GUIDE_SIGN",
		3: "SIGN_TYPE_TRAFFIC_LIGHT_POLE",
	}
	TrafficSign_TrafficSignType_value = map[string]int32{
		"SIGN_TYPE_UNKNOWN":            0,
		"SIGN_TYPE_SPEED_LIMIT":        1,
		"SIGN_TYPE_GUIDE_SIGN":         2,
		"SIGN_TYPE_TRAFFIC_LIGHT_POLE": 3,
	}
)

type Road struct {
	// id
	Id string `json:"id"`
	// 地图 id
	MapId string `json:"map_id"`
	// 名称
	Name string `json:"name"`
	// 道路类型
	Type Road_RoadType `json:"type"`
	// 长度
	Length float64 `json:"length"`
	// 相邻的roadId及其s值
	Neighbors []*RoadPosition `json:"neighbors"`
	// 道路控制点,控制道路中心线线性变化
	ControlPoints []*ControlPoint `json:"control_points"`
	// 道路中心线；参考线，确定道路基本形态，点位的方向等于绘制方向
	ReferenceLine []*ReferencePoint `json:"reference_line"`
	// 双向路段: 由共用某一段参考线的两组 segments 构成
	Sections []*Section `json:"sections"`
	// 路口：
	JunctionPositions []*JunctionPosition `json:"junction_positions"`
}

type Road_RoadType int32

const (
	Road_ROAD_TYPE_UNKNOWN Road_RoadType = 0
	Road_ROAD_TYPE_RAMP    Road_RoadType = 1
)

// Enum value maps for Road_RoadType.
var (
	Road_RoadType_name = map[int32]string{
		0: "ROAD_TYPE_UNKNOWN",
		1: "ROAD_TYPE_RAMP",
	}
	Road_RoadType_value = map[string]int32{
		"ROAD_TYPE_UNKNOWN": 0,
		"ROAD_TYPE_RAMP":    1,
	}
)

type RoadPosition struct {
	RoadId string  `json:"road_id"`
	RoadS  float64 `json:"road_s"`
	S      float64 `json:"s"`
}

type ControlPoint struct {
	Id    string `json:"id"`
	Point *Point `json:"point"`
}

type ReferencePoint struct {
	// s 坐标
	S float64 `json:"s"`
	// heading 朝向
	Heading float64 `json:"heading"`
	// 点坐标
	Point *Point `json:"point"`
	// 高度
	Height float64 `json:"height"`
	// 横向坡度
	CrossSlope float64 `json:"cross_slope"`
}

type Section struct {
	Id string `json:"id"`
	// 在当前道路中的s位置
	S float64 `json:"s"`
	// 当前 Section 的长度
	Length          float64 `json:"length"`
	StartJunctionId string  `json:"start_junction_id"`
	EndJunctionId   string  `json:"end_junction_id"`
	LeftSegmentId   string  `json:"left_segment_id"` // 左右根据绘制方向区分
	RightSegmentId  string  `json:"right_segment_id"`
}

// 被修改前的路口是否需要被记录
type JunctionPosition struct {
	Id string `json:"id"`
	// 对应路口的id
	JunctionId string `json:"junction_id"`
	// 在当前道路中的s位置
	S      float64 `json:"s"`
	Length float64 `json:"length"`
}

type Mapper struct {
	// 映射id
	Id string `json:"id"`
	// 地图 id
	MapId string `json:"map_id"`
	// 原始xpath
	OriginXpath *string `json:"origin_xpath"`
	// 原始elem id
	OriginElemId string `json:"origin_elem_id"`
	// 原始elem type
	OriginElemType string `json:"origin_elem_type"`
	// 目标xpath
	TargetXpath *string `json:"target_xpath"`
	// 目标elem id
	TargetElemId string `json:"target_elem_id"`
	// 目标elem type
	TargetElemType string `json:"target_elem_type"`
}

type Building struct {
	// id
	Id string `json:"id"`
	// string type = 2;
	Pos *Point `json:"pos"`
	// 区域
	Shape *Polygon `json:"shape"`
	// 朝向
	Heading float64 `json:"heading"`
	// 高度
	Height float64 `json:"height"`
	// 宽度
	Width float64 `json:"width"`
	// 长度
	Length   float64 `json:"length"`
	S        float64 `json:"s"`
	T        float64 `json:"t"`
	ModelNum int32   `json:"model_num"`
}

type Tree struct {
	// id
	Id string `json:"id"` //  string type = 2;
	// 位置
	Pos *Point `json:"pos"`
	// 朝向
	Heading float64 `json:"heading"`
	// 高度
	Height float64 `json:"height"`
	// 宽度
	Width float64 `json:"width"`
	// 长度
	Length   float64 `json:"length"`
	S        float64 `json:"s"`
	T        float64 `json:"t"`
	ModelNum int32   `json:"model_num"`
}

type Lamp struct {
	// id
	Id string `json:"id"` //  string type = 2;
	// 位置
	Pos *Point `json:"pos"`
	// 朝向
	Heading float64 `json:"heading"`
	// 高度
	Height float64 `json:"height"`
	// 宽度
	Width float64 `json:"width"`
	// 长度
	Length   float64 `json:"length"`
	S        float64 `json:"s"`
	T        float64 `json:"t"`
	ModelNum int32   `json:"model_num"`
}
