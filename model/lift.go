package model

type Lift struct {
	Name string `json:"liftName"`
	Load string `json:"load"`
	Sets string `json:"sets"`
	Reps string `json:"reps"`
	Rest string `json:"rest"`
	Note string `json:"note"`
	Unit string `json:"unit"`
}
