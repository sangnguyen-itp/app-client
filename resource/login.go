package resource

type RLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Platform string `json:"platform"`
}

type RLogoutRequest struct {
	Platform string `json:"platform"`
}
