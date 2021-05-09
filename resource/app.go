package resource

type RAppRequest struct {
	Name          string `json:"name"`
	LatestVersion string `json:"latest_version"`
	RunningStatus string `json:"running_status"`
	Type          string `json:"type"`
}
