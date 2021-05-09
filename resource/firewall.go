package resource

type RFirewallRequest struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	AppUUID string `json:"app_uuid"`
}
