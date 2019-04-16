package models

type Callback struct {
	Token       string   `json:"token"`
	Challenge   string   `json:"challenge,omitempty"`
	TeamID      string   `json:"team_id,omitempty"`
	ApiAppID    string   `json:"api_app_id,omitempty"`
	Event       Event    `json:"event,omitempty"`
	Type        string   `json:"type"`
	EventID     string   `json:"event_id,omitempty"`
	EventTime   float64  `json:"event_time,omitempty"`
	AuthedUsers []string `json:"authed_users,omitempty"`
}

const (
	UrlVerification = "url_verification"
	EventCallback = "event_callback"
)