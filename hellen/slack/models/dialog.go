package models

type Dialog struct {
	CallbackId 		string 			`json:"callback_id"`
	Title 			string 			`json:"title"`
	SubmitLabel 	string 			`json:"submit_label"`
	NotifyOnCancel 	bool 			`json:"notify_on_cancel"`
	State 			string  		`json:"state,omitempty"`
	Elements 		[]DialogElement `json:"elements"`
}