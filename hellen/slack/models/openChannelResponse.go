package models

type OpenChannelResponse struct {
	Status   bool 	`json:"ok"`
	Channel  map[string]string `json:"channel"`
}

func (response OpenChannelResponse) Id() string {
	return response.Channel["id"]
}