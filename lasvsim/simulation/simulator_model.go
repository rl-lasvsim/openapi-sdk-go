package simulation

type InitReq struct {
	ScenID      string `json:"scen_id"`
	ScenVer     string `json:"scen_ver"`
	SimRecordID string `json:"sim_record_id"`
}

type InitRes struct {
	SimulationId   string `json:"simulation_id"`
	SimulationAddr string `json:"simulation_addr"`
}

type StopReq struct {
	SimulationId string `json:"simulation_id"`
}

type StopRes struct{}

type StepReq struct {
	SimulationID string `json:"simulation_id"`
}

type StepRes struct {
	// 0:运行中;1001:正常结束;1002;未通过
	Code    StepCode `json:"code"`
	Message string   `json:"message"`
}

type StepCode int32

func (s StepCode) IsRuning() bool {
	return s >= 0 && s <= 100
}

func (s StepCode) IsStoped() bool {
	return s == 1001
}

func (s StepCode) IsUnpassed() bool {
	return s == 1002
}

type ResetReq struct {
	SimulationID     string `json:"simulation_id"`
	ResetTrafficFlow bool   `json:"reset_traffic_flow"`
	// 设定测试车辆重置时一些行为, reset_traffic_flow=true时生效。
	ResetVehicle []*ResetVehicleConfig `json:"reset_vehicle"`
	// 设定环境车辆、行人、非机动车重置时一些行为, reset_traffic_flow=true时生效。
	ResetEnvPtcs *ResetEnvPtcs `json:"reset_env_ptcs"`
}

type ResetVehicleConfig struct {
	// 测试车辆ID
	VehicleId string `json:"vehicle_id"`
	// 设置指定导航信息
	LinkPath []string `json:"link_path"`
}

type ResetEnvPtcs struct {
	// 环境车辆参数配置
	VehicleConf *VehicleDistribution `json:"vehicle_conf"`
	// 行人参数配置
	PedConf *PedestrianDistribution `json:"ped_conf"`
	// 非机动车参数配置
	NmvConf *NMVDistribution `json:"nmv_conf"`
}

type VehicleDistribution struct {
	// 密度(范围:0~1)
	Density float64 `json:"density"`
	// 期望速度下限, 单位: [m/s]
	MinTargetSpeed float64 `json:"min_target_speed"`
	// 期望速度上限, 单位: [m/s]
	MaxTargetSpeed float64 `json:"max_target_speed"`
	// 小型车、中型车、大型车的比例(举例: [0.3,0.3,0.4]参数意为30%的小型车、30%的中型车、40%的大型车)
	SizeRatio []float64 `json:"size_ratio"`
	// 保守型、适中型、激进型的比例(举例: [0.3,0.3,0.4]参数意为30%的保守型、30%的适中型、40%的激进型)
	StyleRatio []float64 `json:"style_ratio"`
	// 生疏型、一般型、熟练型的比例(举例: [0.3,0.3,0.4]参数意为30%的生疏型、30%的一般型、40%的熟练型)
	SkillRatio []float64 `json:"skill_ratio"`
}

type PedestrianDistribution struct {
	// 密度(范围:0~1)
	Density float64 `json:"density"`
	// 儿童、成人、老人的比例(举例: [0.3,0.3,0.4]参数意为30%的儿童、30%的成人、40%的老人)
	AgesRatio []float64 `json:"ages_ratio"`
	// 保守型、适中型、激进型的比例(举例: [0.3,0.3,0.4]参数意为30%的保守型、30%的适中型、40%的激进型)
	StyleRatio []float64 `json:"style_ratio"`
}

type NMVDistribution struct {
	// 密度(范围:0~1)
	Density float64 `json:"density"`
	// 子类型的比例(举例: [0.3,0.7]参数意为30%的电动自行车,70%的三轮车)
	SubtypeRatio []float64 `json:"subtype_ratio"`
	// 保守型、适中型、激进型的比例(举例: [0.3,0.3,0.4]参数意为30%的保守型、30%的适中型、40%的激进型)
	StyleRatio []float64 `json:"style_ratio"`
	// 期望速度下限, 单位: [m/s]
	MinTargetSpeed float64 `json:"min_target_speed"`
	// 期望速度上限, 单位: [m/s]
	MaxTargetSpeed float64 `json:"max_target_speed"`
}

type ResetRes struct{}

type GetCurrentStageReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// movementId
	JunctionId string `json:"junction_id"`
}
type GetCurrentStageRes struct {
	MovementIds []string `json:"movement_ids"` // 当前阶段包含的绿灯流向
	Countdown   int32    `json:"countdown"`    // 倒计时(s)
}

type GetMovementSignalReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// movementId
	MovementId string `json:"movement_id"`
}
type GetMovementSignalRes struct {
	CurrentSignal int32 `json:"current_signal"` // 当前灯色
	Countdown     int32 `json:"countdown"`      // 倒计时(s)
}

type GetSignalPlanReq struct {
	SimulationId string `json:"simulation_id"`
	JunctionId   string `json:"junction_id"`
}
type GetSignalPlanRes struct {
	JunctionId string                    `json:"junction_id"`
	Cycle      int32                     `json:"cycle"`
	Offset     int32                     `json:"offset"`
	Stages     []*GetSignalPlanRes_Stage `json:"stages"`
}
type GetSignalPlanRes_Stage struct {
	MovementIds []string `json:"movement_ids"`
	Duration    int32    `json:"duration"` // 时长(s)
}

type GetMovementListReq struct {
	SimulationId string `json:"simulation_id"`
	JunctionId   string `json:"junction_id"`
}
type GetMovementListRes struct {
	// movement 列表
	List []*Movement `json:"list"`
}

type NextStageReq struct {
	SimulationId string `json:"simulation_id"`
	JunctionId   string `json:"junction_id"`
}
type NextStageRes struct{}

type GetVehicleIdListReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
}
type GetVehicleIdListRes struct {
	// 车辆ID列表
	List []string `json:"list"`
}

type GetTestVehicleIdListReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
}
type GetTestVehicleIdListRes struct {
	// 车辆ID列表
	List []string `json:"list"`
}

type GetVehicleBaseInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID列表 - 最多支持1000个ID
	IdList []string `json:"id_list"`
}
type GetVehicleBaseInfoRes struct {
	// 车辆ID与车辆基础信息映射表
	InfoDict map[string]*GetVehicleBaseInfoRes_VehicleBaseInfo `json:"info_dict"`
}
type GetVehicleBaseInfoRes_VehicleBaseInfo struct {
	// 物体基础描述信息
	BaseInfo *ObjBaseInfo `json:"base_info"`
	// 动力学基础描述信息
	DynamicInfo *DynamicInfo `json:"dynamic_info"`
}

type GetVehiclePositionReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID列表 - 最多支持1000个ID
	IdList []string `json:"id_list"`
}
type GetVehiclePositionRes struct {
	// 车辆ID与车辆位置信息映射表
	PositionDict map[string]*Position `json:"position_dict"`
}

type GetVehicleMovingInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID列表 - 最多支持1000个ID
	IdList []string `json:"id_list"`
}
type GetVehicleMovingInfoRes struct {
	// 车辆ID与车辆移动信息映射表
	MovingInfoDict map[string]*ObjMovingInfo `json:"moving_info_dict"`
}

type GetVehicleControlInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID列表 - 最多支持1000个ID
	IdList []string `json:"id_list"`
}
type GetVehicleControlInfoRes struct {
	// 车辆ID与车辆控制参数映射表
	ControlInfoDict map[string]*ControlInfo `json:"control_info_dict"`
}

type GetVehiclePerceptionInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
}
type GetVehiclePerceptionInfoRes struct {
	List []*GetVehiclePerceptionInfoRes_PerceptionObj `json:"list"`
}

// 感知目标对象结构
type GetVehiclePerceptionInfoRes_PerceptionObj struct {
	ObjId      string         `json:"obj_id"`
	BaseInfo   *ObjBaseInfo   `json:"base_info"`
	MovingInfo *ObjMovingInfo `json:"moving_info"`
	Position   *Position      `json:"position"`
}

type GetVehicleReferenceLinesReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
}
type GetVehicleReferenceLinesRes struct {
	ReferenceLines []*ReferenceLine `json:"reference_lines"`
}

type GetVehiclePlanningInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
}
type GetVehiclePlanningInfoRes struct {
	PlanningPath []*Point `json:"planning_path"`
}

type GetVehicleNavigationInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
}
type GetVehicleNavigationInfoRes struct {
	// 暂不支持route_nav以及lane_nav
	NavigationInfo *NavigationInfo `json:"navigation_info"`
}

type GetVehicleCollisionStatusReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
}
type GetVehicleCollisionStatusRes struct {
	CollisionStatus bool `json:"collision_status"`
}

type GetVehicleTargetSpeedReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
}
type GetVehicleTargetSpeedRes struct {
	TargetSpeed float64 `json:"target_speed"`
}

type SetVehiclePlanningInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 轨迹点
	PlanningPath []*Point `json:"planning_path"`
	// 轨迹点速度
	Speed []float64 `json:"speed"`
}
type SetVehiclePlanningInfoRes struct{}

type SetVehicleControlInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 方向盘转角, 如果为空则不生效
	SteWheel *float64 `json:"ste_wheel"`
	// 纵向加速度, 如果为空则不生效
	LonAcc *float64 `json:"lon_acc"`
}
type SetVehicleControlInfoRes struct{}

type SetVehiclePositionReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 设置{x,y,z}，如果为空则不生效
	Point *Point `json:"point"`
	// 设置航向角, 如果为空则不生效
	Phi *float64 `json:"phi"`
}
type SetVehiclePositionRes struct{}

type SetVehicleMovingInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 纵向速度速度, 如果为空则不生效
	U *float64 `json:"u"`
	// 横向速度, 如果为空则不生效
	V *float64 `json:"v"`
	// 角速度, 如果为空则不生效
	W *float64 `json:"w"`
	// 纵向加速度, 如果为空则不生效
	UAcc *float64 `json:"u_acc"`
	// 横向加速度, 如果为空则不生效
	VAcc *float64 `json:"v_acc"`
	// 角加速度, 如果为空则不生效
	WAcc *float64 `json:"w_acc"`
}
type SetVehicleMovingInfoRes struct{}

type SetVehicleBaseInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 基础信息
	BaseInfo *ObjBaseInfo `json:"base_info"`
	// 动力学基础信息(暂不支持)
	DynamicInfo *DynamicInfo `json:"dynamic_info"`
}
type SetVehicleBaseInfoRes struct{}

type SetVehicleRouteNavReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 路段导航
	RouteNav []string `json:"route_nav"`
}
type SetVehicleRouteNavRes struct{}

type SetVehicleLinkNavReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 子路段导航
	LinkNav []string `json:"link_nav"`
}
type SetVehicleLinkNavRes struct{}

type SetVehicleLaneNavReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 路段导航
	LaneNav []*LaneNav `json:"lane_nav"`
}
type SetVehicleLaneNavRes struct{}

type SetVehicleDestinationReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 终点位置信息
	Destination *Point `json:"destination"`
}
type SetVehicleDestinationRes struct{}

type GetPedIdListReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
}
type GetPedIdListRes struct {
	List []string `json:"list"`
}

type GetPedBaseInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 行人ID列表 - 最多支持1000个ID
	PedIdList []string `json:"ped_id_list"`
}
type GetPedBaseInfoRes struct {
	// 行人基础信息
	BaseInfoDict map[string]*ObjBaseInfo `json:"base_info_dict"`
}

type SetPedPositionReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 行人ID
	PedId string `json:"ped_id"`
	// 坐标{x,y,z}
	Point *Point `json:"point"`
	// 航向角
	Phi *float64 `json:"phi"`
}
type SetPedPositionRes struct{}

type GetNMVIdListReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
}
type GetNMVIdListRes struct {
	List []string `json:"list"`
}

type GetNMVBaseInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 非机动车ID列表 - 最多支持1000个ID
	NmvIdList []string `json:"nmv_id_list"`
}
type GetNMVBaseInfoRes struct {
	//非机动车基础信息
	BaseInfoDict map[string]*ObjBaseInfo `json:"base_info_dict"`
}

type SetNMVPositionReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 非机动车ID
	NmvId string `json:"nmv_id"`
	// 坐标{x,y,z}
	Point *Point `json:"point"`
	// 航向角
	Phi *float64 `json:"phi"`
}
type SetNMVPositionRes struct{}

type GetStepSpawnIdListReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
}

type GetStepSpawnIdListRes struct {
	IdList []string `json:"id_list"`
}

type GetParticipantBaseInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 参与者ID列表 - 最多支持1000个ID
	ParticipantIdList []string `json:"participant_id_list"`
}

type GetParticipantBaseInfoRes struct {
	// 参与者基础信息
	BaseInfoDict map[string]*ObjBaseInfo `json:"base_info_dict"`
}

type GetParticipantMovingInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 参与者ID列表 - 最多支持1000个ID
	ParticipantIdList []string `json:"participant_id_list"`
}

type GetParticipantMovingInfoRes struct {
	// 参与者移动信息
	MovingInfoDict map[string]*ObjMovingInfo `json:"moving_info_dict"`
}

type GetParticipantPositionReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 参与者ID列表 - 最多支持1000个ID
	ParticipantIdList []string `json:"participant_id_list"`
}

type GetParticipantPositionRes struct {
	// 参与者位置信息
	PositionDict map[string]*Position `json:"position_dict"`
}

type GetVehicleSensorConfigReq struct {
	SimulationId string `json:"simulation_id"`
	VehicleId    string `json:"vehicle_id"`
}

type GetVehicleSensorConfigRes struct {
	SensorsConfig []*SensorConfig `json:"sensors_config"`
}

// 设置车辆道路感知信息
type SetVehicleRoadPerceptionInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 感知地图信息
	Noa *LocalMap `json:"noa"`
}

type SetVehicleRoadPerceptionInfoRes struct{}

// 设置车辆障碍物感知信息
type SetVehicleObstaclePerceptionInfoReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 障碍物信息
	Obstacles []*Obstacle `json:"obstacles"`
}

type SetVehicleObstaclePerceptionInfoRes struct{}

// 设置车辆指标
type SetVehicleExtraMetricsReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 额外指标
	Metrics map[string]float64 `json:"metrics"`
}

type SetVehicleExtraMetricsRes struct{}

type SetVehicleLocalPathsReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
	// 本地路径
	LocalPaths []*LocalPath `json:"local_paths"`
	// 选择的路径 - 不填默认选择概率最高的
	ChooseIdx *int32 `json:"choose_idx"`
}
type LocalPath struct {
	Points []*Point `json:"points"`
	Prob   float64  `json:"prob"`
}

type SetVehicleLocalPathsRes struct{}

type GetIdcVehicleNavReq struct {
	// 仿真ID
	SimulationId string `json:"simulation_id"`
	// 车辆ID
	VehicleId string `json:"vehicle_id"`
}

type GetIdcVehicleNavRes struct {
	// 子路段导航
	LinkPathNav []string `json:"link_path_nav"`
	// 子路段与路口拼接
	LinkJunctionNav []string `json:"link_junction_nav"`
	// 下一个路口
	NextJunctionId string `json:"next_junction_id"`
	// 距离下一个路口距离,路口内为0.0
	DisToNextJunction float64 `json:"dis_to_next_junction"`
	// 下一个movement
	NextMovementId string `json:"next_movement_id"`
}

// NOTE: ---车辆接口的细节结构---
type ObjBaseInfo struct {
	// 宽(m)
	Width float64 `json:"width"`
	// 高(m)
	Height float64 `json:"height"`
	// 长(m)
	Length float64 `json:"length"`
	// 重量(kg)
	Weight float64 `json:"weight"`

	MaxDec float64 `json:"max_dec"`
	MaxAcc float64 `json:"max_acc"`
	// m/s单位
	MaxSpeed float64 `json:"max_speed"`
}

type DynamicInfo struct {
	// 前轮转弯刚度[N/rad]
	FrontWheelStiffness float64 `json:"front_wheel_stiffness"`
	// 后轮转弯刚度[N/rad]
	RearWheelStiffness float64 `json:"rear_wheel_stiffness"`
	// 前轴到重心距离[m]
	FrontAxleToCenter float64 `json:"front_axle_to_center"`
	// 后轴到重心距离[m]
	RearAxleToCenter float64 `json:"rear_axle_to_center"`
	// 重心处的极惯性矩[kg*m^2]
	YawMomentOfInertia float64 `json:"yaw_moment_of_inertia"`
}

type Position struct {
	Point        *Point   `json:"point"`
	Phi          float64  `json:"phi"`
	LaneId       string   `json:"lane_id"`
	LinkId       string   `json:"link_id"`
	JunctionId   string   `json:"junction_id"`
	SegmentId    string   `json:"segment_id"`
	DisToLaneEnd *float64 `json:"dis_to_lane_end"`
	// 1 - 地图外 TODO 还有哪些选项
	PositionType int32 `json:"position_type"`

	Type Position_PositionType `json:"type"`
	// 朝向
	Heading *float64 `json:"heading"`
	// 横滚角
	Roll *float64 `json:"roll"`
	// 俯仰角
	Patch *float64 `json:"patch"`
	// 车道编号
	LaneIndex *int32 `json:"lane_index"`
	// 车道偏移
	LaneOffset *float64 `json:"lane_offset"`
	// s值
	S *float64 `json:"s"`
	// t值
	T *float64 `json:"t"`
}

type Position_PositionType int32

const (
	Position_POSITION_TYPE_UNKNOWN Position_PositionType = 0
	// 1. 在车道内
	Position_POSITION_TYPE_IN_LANE Position_PositionType = 1
	// 2. 在路口内
	Position_POSITION_TYPE_IN_JUNCTION Position_PositionType = 2
	// 3. 在道路外
	Position_POSITION_TYPE_OUT_ROAD Position_PositionType = 3
)

var (
	Position_PositionType_name = map[int32]string{
		0: "POSITION_TYPE_UNKNOWN",
		1: "POSITION_TYPE_IN_LANE",
		2: "POSITION_TYPE_IN_JUNCTION",
		3: "POSITION_TYPE_OUT_ROAD",
	}
	Position_PositionType_value = map[string]int32{
		"POSITION_TYPE_UNKNOWN":     0,
		"POSITION_TYPE_IN_LANE":     1,
		"POSITION_TYPE_IN_JUNCTION": 2,
		"POSITION_TYPE_OUT_ROAD":    3,
	}
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
type ObjMovingInfo struct {
	// 纵向速度[m/s]
	U float64 `protobuf:"fixed64,2,opt,name=u,proto3" json:"u"`
	// 纵向加速度[m/s^2]
	UAcc float64 `protobuf:"fixed64,3,opt,name=u_acc,json=uAcc,proto3" json:"u_acc"`
	// 横向速度[m/s]
	V float64 `protobuf:"fixed64,4,opt,name=v,proto3" json:"v"`
	// 横向加速度[m/s^2]
	VAcc float64 `protobuf:"fixed64,5,opt,name=v_acc,json=vAcc,proto3" json:"v_acc"`
	// 角速度[m/s]
	W float64 `protobuf:"fixed64,6,opt,name=w,proto3" json:"w"`
	// 角加速度[m/s^2]
	WAcc float64 `protobuf:"fixed64,7,opt,name=w_acc,json=wAcc,proto3" json:"w_acc"`

	Heading float64 `protobuf:"fixed64,1,opt,name=heading,proto3" json:"heading"` // 行驶方向(单位: 弧度)
}

type ControlInfo struct {
	// 方向盘转角[逆时针为正]
	SteWheel float64 `protobuf:"fixed64,1,opt,name=ste_wheel,json=steWheel,proto3" json:"ste_wheel"`
	// 纵向加速度[m/s^2]
	LonAcc float64 `protobuf:"fixed64,2,opt,name=lon_acc,json=lonAcc,proto3" json:"lon_acc"`
	// 左前轮扭矩[N*m]
	FlTorque float64 `protobuf:"fixed64,3,opt,name=fl_torque,json=flTorque,proto3" json:"fl_torque"`
	// 右前轮扭矩[N*m]
	FrTorque float64 `protobuf:"fixed64,4,opt,name=fr_torque,json=frTorque,proto3" json:"fr_torque"`
	// 左后轮扭矩[N*m]
	RlTorque float64 `protobuf:"fixed64,5,opt,name=rl_torque,json=rlTorque,proto3" json:"rl_torque"`
	// 右后轮扭矩[N*m]
	RrTorque float64 `protobuf:"fixed64,6,opt,name=rr_torque,json=rrTorque,proto3" json:"rr_torque"`
}

type ReferenceLine struct {
	LaneIds []string `protobuf:"bytes,1,rep,name=lane_ids,json=laneIds,proto3" json:"lane_ids"`
	// "lane"和"connection"
	LaneTypes []string `protobuf:"bytes,2,rep,name=lane_types,json=laneTypes,proto3" json:"lane_types"`
	Points    []*Point `protobuf:"bytes,3,rep,name=points,proto3" json:"points"`
	LaneIdxes []int32  `protobuf:"varint,4,rep,packed,name=lane_idxes,json=laneIdxes,proto3" json:"lane_idxes"`
	// 是否逆行
	Opposite bool `protobuf:"varint,5,opt,name=opposite,proto3" json:"opposite"`
}
type NavigationInfo struct {
	// 路段导航
	// RouteNav []string `protobuf:"bytes,1,rep,name=route_nav,json=routeNav,proto3" json:"route_nav"`
	// 子路段导航
	LinkNav []string `protobuf:"bytes,2,rep,name=link_nav,json=linkNav,proto3" json:"link_nav"`
	// 车道导航
	// LaneNav []*LaneNav `protobuf:"bytes,3,rep,name=lane_nav,json=laneNav,proto3" json:"lane_nav"`
	// 终点
	Destination *Position `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination"`
}

type SensorConfig struct {
	SensorId   string                  `json:"sensor_id"`
	SensorType SensorConfig_SensorType `json:"sensor_type"`
	// 检测范围角度（扇形圆心角）
	DetectAngle float64 `json:"detect_angle"`
	// 检测范围距离 （扇形半径）
	DetectRange float64 `json:"detect_range"`
	// 相对车辆质心纵向偏移
	InstallX float64 `json:"install_x"`
	// 相对车辆质心横向偏移
	InstallY float64 `json:"install_y"`
	// 安装位置与交通参与者朝向之间的夹角
	InstallPhi float64 `json:"install_phi"`
	// 传感器感知精度误差
	SensorError *SensorErrorConfig `json:"sensor_error"`
	// 相对车辆质心纵向偏移
	InstallLon float64 `json:"install_lon"`
	// 相对车辆质心横向偏移
	InstallLat float64 `json:"install_lat"`
}

type SensorConfig_SensorType int32

const (
	SensorConfig_UNKNOWN SensorConfig_SensorType = 0
	// 摄像头
	SensorConfig_CAMERA SensorConfig_SensorType = 1
	// 激光雷达
	SensorConfig_LIDAR SensorConfig_SensorType = 2
	// 毫米波雷达
	SensorConfig_RADAR SensorConfig_SensorType = 3
)

// Enum value maps for SensorConfig_SensorType.
var (
	SensorConfig_SensorType_name = map[int32]string{
		0: "UNKNOWN",
		1: "CAMERA",
		2: "LIDAR",
		3: "RADAR",
	}
	SensorConfig_SensorType_value = map[string]int32{
		"UNKNOWN": 0,
		"CAMERA":  1,
		"LIDAR":   2,
		"RADAR":   3,
	}
)

type SensorErrorConfig struct {
	// 位置方差
	LocationSigma float64 `json:"location_sigma"`
	// 朝向角方差
	PhiSigma float64 `json:"phi_sigma"`
	// 尺寸方差
	SizeSigma float64 `json:"size_sigma"`
	// 速度方差
	VelocitySigma float64 `json:"velocity_sigma"`
}

type LocalMap struct {
	// 车道边界线感知
	LaneBoundaries []*LaneBoundary `protobuf:"bytes,1,rep,name=lane_boundaries,json=laneBoundaries,proto3" json:"lane_boundaries"`
	// 路口感知
	Junctions []*Polygon `protobuf:"bytes,2,rep,name=junctions,proto3" json:"junctions"`
	// 人行横道感知
	Crosswalks []*Polygon `protobuf:"bytes,3,rep,name=crosswalks,proto3" json:"crosswalks"`
	// 红绿灯感知 <movement_id, TrafficLight> 0(无) | 1(红) | 2(黄) | 3(绿)
	TrafficLightColors map[string]int32 `protobuf:"bytes,4,rep,name=traffic_light_colors,json=trafficLightColors,proto3" json:"traffic_light_colors" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// 停车线
	StopLines []*StopLines `protobuf:"bytes,5,rep,name=stop_lines,json=stopLines,proto3" json:"stop_lines"`
	// 感知车道中心线
	LaneCenterLines []*LaneCenterLines `protobuf:"bytes,6,rep,name=lane_center_lines,json=laneCenterLines,proto3" json:"lane_center_lines"`
	// 虚拟区域
	VirtualPolygons []*Polygon `protobuf:"bytes,7,rep,name=virtual_polygons,json=virtualPolygons,proto3" json:"virtual_polygons"`
	// 参考线
	ReferenceLines []*ReferenceLines `protobuf:"bytes,8,rep,name=reference_lines,json=referenceLines,proto3" json:"reference_lines"`
}

// 停车线
type StopLines struct {
	// 线段
	Line *LineString `json:"line"`
	// 线段类型; "solid" : "实线" | "broken" : "虚线" | "marked": "散点标记"
	Style string `json:"style"`
	// 颜色; "yellow": "黄色" | "red": "红色" | "green": "绿色" | "blue" : "蓝色" | "purple" : "紫色" |"brown" : "棕色"
	Color string `json:"color"`
}

// 感知车道中心线
type LaneCenterLines struct {
	// 线段
	Line *LineString `json:"line"`
	// 线段类型; "solid" : "实线" | "broken" : "虚线" | "marked": "散点标记"
	Style string `json:"style"`
	// 颜色; "yellow": "黄色" | "red": "红色" | "green": "绿色" | "blue" : "蓝色" | "purple" : "紫色" |"brown" : "棕色"
	Color string `json:"color"`
}

// 车道边界线感知
type LaneBoundary struct {
	// 线段
	Line *LineString `json:"line"`
	// 线段类型; "solid" : "实线" | "broken" : "虚线" | "marked": "散点标记"
	Style string `json:"style"`
	// 颜色; "yellow": "黄色" | "red": "红色" | "green": "绿色" | "blue" : "蓝色" | "purple" : "紫色" |"brown" : "棕色"
	Color string `json:"color"`
}

// 线段
type LineString struct {
	Points []*Point `json:"points"`
}

// 多边形
type Polygon struct {
	Points []*Point `json:"points"`
	// 线段类型; "solid" : "实线" | "broken" : "虚线" | "marked": "散点标记"
	Style string `json:"style"`
	// 颜色; "yellow": "黄色" | "red": "红色" | "green": "绿色" | "blue" : "蓝色" | "purple" : "紫色" |"brown" : "棕色"
	Color string `json:"color"`
}

type ReferenceLines struct {
	// 线段
	Line *LineString `json:"line"`
	// 线段类型; "solid" : "实线" | "broken" : "虚线" | "marked": "散点标记"
	Style string `json:"style"`
	// 颜色; "yellow": "黄色" | "red": "红色" | "green": "绿色" | "blue" : "蓝色" | "purple" : "紫色" |"brown" : "棕色"
	Color string `json:"color"`
}

// 障碍物结构
type Obstacle struct {
	Id         string           `json:"id"`
	Type       Obstacle_ObjType `json:"type"`
	BaseInfo   *ObjBaseInfo     `json:"base_info"`
	MovingInfo *ObjMovingInfo   `json:"moving_info"`
	Position   *Position        `json:"position"`
}
type Obstacle_ObjType int32

const (
	// 机动车
	Obstacle_TYPE_VEHICLE Obstacle_ObjType = 0
	// 行人
	Obstacle_TYPE_PEDESTRIAN Obstacle_ObjType = 1
	// 非机动车
	Obstacle_TYPE_NMV Obstacle_ObjType = 2
)

// Enum value maps for Obstacle_ObjType.
var (
	Obstacle_ObjType_name = map[int32]string{
		0: "TYPE_VEHICLE",
		1: "TYPE_PEDESTRIAN",
		2: "TYPE_NMV",
	}
	Obstacle_ObjType_value = map[string]int32{
		"TYPE_VEHICLE":    0,
		"TYPE_PEDESTRIAN": 1,
		"TYPE_NMV":        2,
	}
)

// ---------地图movement-------
type Movement struct {
	MapId      uint64 `protobuf:"varint,1,opt,name=map_id,json=mapId,proto3" json:"map_id"`
	MovementId string `protobuf:"bytes,2,opt,name=movement_id,json=movementId,proto3" json:"movement_id"`
	// 入路口link id
	UpstreamLinkId string `protobuf:"bytes,3,opt,name=upstream_link_id,json=upstreamLinkId,proto3" json:"upstream_link_id"`
	// 出路口link id
	DownstreamLinkId string `protobuf:"bytes,4,opt,name=downstream_link_id,json=downstreamLinkId,proto3" json:"downstream_link_id"`
	// junction id
	JunctionId string `protobuf:"bytes,5,opt,name=junction_id,json=junctionId,proto3" json:"junction_id"`
	// 流向
	// "straight": "直行",
	// "left":     "左转",
	// "right":    "右转",
	// "turn":     "掉头"
	FlowDirection string `protobuf:"bytes,6,opt,name=flow_direction,json=flowDirection,proto3" json:"flow_direction"`
}

type LaneNav struct {
	Nav map[int32]string `protobuf:"bytes,1,rep,name=nav,proto3" json:"nav" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
