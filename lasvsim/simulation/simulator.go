package simulation

import (
	"github.com/rl-lasvsim/openapi-sdk-go/lasvsim/httpclient"
)

type SimulatorConfig struct {
	ScenID      string `json:"scen_id"`
	ScenVer     string `json:"scen_ver"`
	SimRecordID string `json:"sim_record_id"`
}

type Simulator struct {
	httpClient   *httpclient.HttpClient
	SimulationId string
}

func NewSimulatorFromConfig(hCli *httpclient.HttpClient, cfg SimulatorConfig) (*Simulator, error) {
	cloneCli := hCli.Clone()

	simtor := &Simulator{
		httpClient: cloneCli,
	}

	err := simtor.initFromConfig(cfg)
	if err != nil {
		return nil, err
	}

	return simtor, nil
}

func NewSimulatorFromSim(hCli *httpclient.HttpClient, simId, simAddr string) (*Simulator, error) {
	cloneCli := hCli.Clone()
	simtor := &Simulator{
		httpClient: cloneCli,
	}

	err := simtor.initFromSim(simId, simAddr)
	if err != nil {
		return nil, err
	}

	return simtor, nil
}
func (s *Simulator) initFromConfig(simConfig SimulatorConfig) error {
	var reply struct {
		SimulationId   string `json:"simulation_id"`
		SimulationAddr string `json:"simulation_addr"`
	}

	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/init",
		&InitReq{
			ScenID:      simConfig.ScenID,
			ScenVer:     simConfig.ScenVer,
			SimRecordID: simConfig.SimRecordID,
		},
		&reply,
	)
	if err != nil {
		return err
	}

	return s.initFromSim(reply.SimulationId, reply.SimulationAddr)
}

func (s *Simulator) initFromSim(simId, simAddr string) error {
	s.httpClient.Headers["x-md-simulation_id"] = simId
	s.httpClient.Headers["x-md-rl-direct-addr"] = simAddr
	s.SimulationId = simId
	return nil
}

func (s *Simulator) Step() (*StepRes, error) {
	var reply StepRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/step",
		&StepReq{SimulationID: s.SimulationId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) Stop() error {
	var reply StopRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/stop",
		&StopReq{SimulationId: s.SimulationId},
		&reply,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Simulator) Reset(resetTrafficFlow bool) (*ResetRes, error) {
	var reply ResetRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/reset",
		&ResetReq{SimulationID: s.SimulationId, ResetTrafficFlow: resetTrafficFlow},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

// --------- 地图部分 ---------
func (s *Simulator) GetCurrentStage(junctionId string) (*GetCurrentStageRes, error) {
	var reply GetCurrentStageRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/map/traffic_light/current_stage/get",
		&GetCurrentStageReq{SimulationId: s.SimulationId, JunctionId: junctionId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
func (s *Simulator) GetMovementSignal(movementId string) (*GetMovementSignalRes, error) {
	var reply GetMovementSignalRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/map/traffic_light/phase_info/get",
		&GetMovementSignalReq{SimulationId: s.SimulationId, MovementId: movementId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
func (s *Simulator) GetSignalPlan(junctionId string) (*GetSignalPlanRes, error) {
	var reply GetSignalPlanRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/map/traffic_light/signal_plan/get",
		&GetSignalPlanReq{SimulationId: s.SimulationId, JunctionId: junctionId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
func (s *Simulator) GetMovementList(junctionId string) (*GetMovementListRes, error) {
	var reply GetMovementListRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/map/movement/list/get",
		&GetMovementListReq{SimulationId: s.SimulationId, JunctionId: junctionId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

// --------- 车辆部分 ---------
func (s *Simulator) GetVehicleIdList() (*GetVehicleIdListRes, error) {
	var reply GetVehicleIdListRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/id_list/get",
		&GetVehicleIdListReq{SimulationId: s.SimulationId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetTestVehicleIdList() (*GetTestVehicleIdListRes, error) {
	var reply GetTestVehicleIdListRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/test_vehicle/id_list/get",
		&GetTestVehicleIdListReq{SimulationId: s.SimulationId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetVehicleBaseInfo(vehicleIdList []string) (*GetVehicleBaseInfoRes, error) {
	var reply GetVehicleBaseInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/base_info/get",
		&GetVehicleBaseInfoReq{SimulationId: s.SimulationId, IdList: vehicleIdList},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetVehiclePosition(vehicleIdList []string) (*GetVehiclePositionRes, error) {
	var reply GetVehiclePositionRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/position/get",
		&GetVehiclePositionReq{SimulationId: s.SimulationId, IdList: vehicleIdList},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetVehicleMovingInfo(vehicleIdList []string) (*GetVehicleMovingInfoRes, error) {
	var reply GetVehicleMovingInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/moving_info/get",
		&GetVehicleMovingInfoReq{SimulationId: s.SimulationId, IdList: vehicleIdList},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetVehicleControlInfo(vehicleIdList []string) (*GetVehicleControlInfoRes, error) {
	var reply GetVehicleControlInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/control/get",
		&GetVehicleControlInfoReq{SimulationId: s.SimulationId, IdList: vehicleIdList},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetVehiclePerceptionInfo(vehicleId string) (*GetVehiclePerceptionInfoRes, error) {
	var reply GetVehiclePerceptionInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/perception/get",
		&GetVehiclePerceptionInfoReq{SimulationId: s.SimulationId, VehicleId: vehicleId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetVehicleReferenceLines(vehicleId string) (*GetVehicleReferenceLinesRes, error) {
	var reply GetVehicleReferenceLinesRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/reference_line/get",
		&GetVehicleReferenceLinesReq{SimulationId: s.SimulationId, VehicleId: vehicleId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetVehiclePlanningInfo(vehicleId string) (*GetVehiclePlanningInfoRes, error) {
	var reply GetVehiclePlanningInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/planning/get",
		&GetVehiclePlanningInfoReq{SimulationId: s.SimulationId, VehicleId: vehicleId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetVehicleNavigationInfo(vehicleId string) (*GetVehicleNavigationInfoRes, error) {
	var reply GetVehicleNavigationInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/navigation/get",
		&GetVehicleNavigationInfoReq{SimulationId: s.SimulationId, VehicleId: vehicleId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetVehicleCollisionStatus(vehicleId string) (*GetVehicleCollisionStatusRes, error) {
	var reply GetVehicleCollisionStatusRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/collision/get",
		&GetVehicleCollisionStatusReq{SimulationId: s.SimulationId, VehicleId: vehicleId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetVehicleTargetSpeed(vehicleId string) (*GetVehicleTargetSpeedRes, error) {
	var reply GetVehicleTargetSpeedRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/target_speed/get",
		&GetVehicleTargetSpeedReq{SimulationId: s.SimulationId, VehicleId: vehicleId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) SetVehiclePlanningInfo(vehicleId string, planningPath []*Point, speed []float64) (*SetVehiclePlanningInfoRes, error) {
	var reply SetVehiclePlanningInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/planning/set",
		&SetVehiclePlanningInfoReq{SimulationId: s.SimulationId, VehicleId: vehicleId, PlanningPath: planningPath, Speed: speed},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) SetVehicleControlInfo(vehicleId string, steWheel *float64, lonAcc *float64) (*SetVehicleControlInfoRes, error) {
	var reply SetVehicleControlInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/control/set",
		&SetVehicleControlInfoReq{SimulationId: s.SimulationId, VehicleId: vehicleId, SteWheel: steWheel, LonAcc: lonAcc},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) SetVehiclePosition(vehicleId string, point *Point, phi *float64) (*SetVehiclePositionRes, error) {
	var reply SetVehiclePositionRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/position/set",
		&SetVehiclePositionReq{SimulationId: s.SimulationId, VehicleId: vehicleId, Point: point, Phi: phi},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) SetVehicleMovingInfo(vehicleId string, u, v, w, uAcc, vAcc, wAcc *float64) (*SetVehicleMovingInfoRes, error) {
	var reply SetVehicleMovingInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/moving_info/set",
		&SetVehicleMovingInfoReq{SimulationId: s.SimulationId, VehicleId: vehicleId, U: u, V: v, W: w, UAcc: uAcc, VAcc: vAcc, WAcc: wAcc},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) SetVehicleBaseInfo(vehicleId string, baseInfo *ObjBaseInfo, dynamicInfo *DynamicInfo) (*SetVehicleBaseInfoRes, error) {
	var reply SetVehicleBaseInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/base_info/set",
		&SetVehicleBaseInfoReq{SimulationId: s.SimulationId, VehicleId: vehicleId, BaseInfo: baseInfo, DynamicInfo: dynamicInfo},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) SetVehicleLinkNav(vehicleId string, linkNav []string) (*SetVehicleLinkNavRes, error) {
	var reply SetVehicleLinkNavRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/link_nav/set",
		&SetVehicleLinkNavReq{SimulationId: s.SimulationId, VehicleId: vehicleId, LinkNav: linkNav},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) SetVehicleDestination(vehicleId string, destination *Point) (*SetVehicleDestinationRes, error) {
	var reply SetVehicleDestinationRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/destination/set",
		&SetVehicleDestinationReq{SimulationId: s.SimulationId, VehicleId: vehicleId, Destination: destination},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

// --------- 行人部分 ---------
func (s *Simulator) GetPedIdList() (*GetPedIdListRes, error) {
	var reply GetPedIdListRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/ped/id_list/get",
		&GetPedIdListReq{SimulationId: s.SimulationId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetPedBaseInfo(pedIdList []string) (*GetPedBaseInfoRes, error) {
	var reply GetPedBaseInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/ped/base_info/get",
		&GetPedBaseInfoReq{SimulationId: s.SimulationId, PedIdList: pedIdList},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) SetPedPosition(pedId string, point *Point, phi *float64) (*SetPedPositionRes, error) {
	var reply SetPedPositionRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/ped/position/set",
		&SetPedPositionReq{SimulationId: s.SimulationId, PedId: pedId, Point: point, Phi: phi},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

// --------- 非机动车部分 ---------
func (s *Simulator) GetNMVIdList() (*GetNMVIdListRes, error) {
	var reply GetNMVIdListRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/nmv/id_list/get",
		&GetNMVIdListReq{SimulationId: s.SimulationId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetNMVBaseInfo(nmvIdList []string) (*GetNMVBaseInfoRes, error) {
	var reply GetNMVBaseInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/nmv/base_info/get",
		&GetNMVBaseInfoReq{SimulationId: s.SimulationId, NmvIdList: nmvIdList},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) SetNMVPosition(nmvId string, point *Point, phi *float64) (*SetNMVPositionRes, error) {
	var reply SetNMVPositionRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/nmv/position/set",
		&SetNMVPositionReq{SimulationId: s.SimulationId, NmvId: nmvId, Point: point, Phi: phi},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetStepSpawnIdList() (*GetStepSpawnIdListRes, error) {
	var reply GetStepSpawnIdListRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/participant/step_spawn_ids/get",
		&GetStepSpawnIdListReq{SimulationId: s.SimulationId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetParticipantBaseInfo(participantIdList []string) (*GetParticipantBaseInfoRes, error) {
	var reply GetParticipantBaseInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/participant/base_info/get",
		&GetParticipantBaseInfoReq{SimulationId: s.SimulationId, ParticipantIdList: participantIdList},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetParticipantMovingInfo(participantIdList []string) (*GetParticipantMovingInfoRes, error) {
	var reply GetParticipantMovingInfoRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/participant/moving_info/get",
		&GetParticipantMovingInfoReq{SimulationId: s.SimulationId, ParticipantIdList: participantIdList},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetParticipantPosition(participantIdList []string) (*GetParticipantPositionRes, error) {
	var reply GetParticipantPositionRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/participant/position/get",
		&GetParticipantPositionReq{SimulationId: s.SimulationId, ParticipantIdList: participantIdList},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *Simulator) GetVehicleSensorConfig(vehicleId string) (*GetVehicleSensorConfigRes, error) {
	var reply GetVehicleSensorConfigRes
	err := s.httpClient.Post(
		"/openapi/cosim/v2/simulation/vehicle/sensor_config/get",
		&GetVehicleSensorConfigReq{SimulationId: s.SimulationId, VehicleId: vehicleId},
		&reply,
	)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
