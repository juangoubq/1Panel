package dto

type SettingInfo struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`

	SessionTimeout string `json:"sessionTimeout"`
	LocalTime      string `json:"localTime"`

	PanelName string `json:"panelName"`
	Theme     string `json:"theme"`
	Language  string `json:"language"`

	ServerPort             string `json:"serverPort"`
	SecurityEntrance       string `json:"securityEntrance"`
	ExpirationDays         string `json:"expirationDays"`
	ExpirationTime         string `json:"expirationTime"`
	ComplexityVerification string `json:"complexityVerification"`
	MFAStatus              string `json:"mfaStatus"`
	MFASecret              string `json:"mfaSecret"`

	MonitorStatus    string `json:"monitorStatus"`
	MonitorStoreDays string `json:"monitorStoreDays"`

	MessageType string `json:"messageType"`
	EmailVars   string `json:"emailVars"`
	WeChatVars  string `json:"weChatVars"`
	DingVars    string `json:"dingVars"`
}

type SettingUpdate struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value"`
}

type PasswordUpdate struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
