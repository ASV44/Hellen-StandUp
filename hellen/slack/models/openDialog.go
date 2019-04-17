package models

type OpenDialog struct {
	TriggerId 	string 	`json:"trigger_id"`
	Dialog 		Dialog 	`json:"dialog"`
}