package traintask

type GetSceneIdListReq struct {
	TaskId uint64 `json:"task_id"`
}

type GetSceneIdListRes struct {
	SceneIdList      []string `json:"scene_id_list"`
	SceneVersionList []string `json:"scene_version_list"`
}
