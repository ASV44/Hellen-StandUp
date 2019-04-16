package models

type OpenChannel struct {
	Status   bool 				`json:"ok"`
	Channel  map[string]string `json:"channel"`
}

func (response OpenChannel) Id() string {
	return response.Channel["id"]
}