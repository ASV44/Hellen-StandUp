package models

type Event struct {
	ClientMsgID string 	`json:"client_msg_id"`
	Type string 		`json:"type"`
	Text string 		`json:"text"`
	User string 		`json:"user"`
	TS string 			`json:"ts"`
	Channel string 		`json:"channel"`
	EventTS string 		`json:"event_ts"`
	ChannelType string 	`json:"channel_type"`
}