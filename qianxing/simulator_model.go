package qianxing

type InitReq struct {
	ScenID      string `json:"scen_id,omitempty"`
	ScenVer     string `json:"scen_ver,omitempty"`
	SimRecordID int    `json:"sim_record_id,omitempty"`
}

type InitRes struct {
	SimulationId   string `json:"simulation_id"`
	SimulationAddr string `json:"simulation_addr"`
}

type StopReq struct {
	SimulationId string `json:"simulation_id"`
}

type StopRes struct {
}

type StepReq struct {
	SimulationID string `json:"simulation_id"`
}

type StepRes struct {
	// Define fields according to the expected response structure
}

type ResetReq struct {
	SimulationID     string `json:"simulation_id"`
	ResetTrafficFlow bool   `json:"reset_traffic_flow"`
}

type ResetRes struct {
	// Define fields according to the expected response structure
}

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
	// 行人ID列表 - 最多支持100个ID
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
	// 非机动车ID列表 - 最多支持100个ID
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
}

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
