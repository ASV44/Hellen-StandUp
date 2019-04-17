package models

type DialogElement struct {
	Type 		string 	`json:"type"`
	Label 		string 	`json:"label"`
	Name 		string 	`json:"name"`
	Optional 	bool 	`json:"optional,omitempty"`
	Hint		string	`json:"hint,omitempty"`
	Subtype		string 	`json:"subtype,omitempty"`
	Value		string  `json:"value,omitempty"`
	Placeholder	string	`json:"placeholder,omitempty"`
	MaxLength	int		`json:"max_length,omitempty"`
	MinLength	int		`json:"min_length,omitempty"`
}

const (
	Text 		= "text"
	Textarea 	= "textarea"
	Select 		= "select "
)