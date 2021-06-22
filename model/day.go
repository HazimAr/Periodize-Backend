package model

type Day struct {
	Name       string `json:"dayName"`
	Descripion string `json:"dayDescription"`
	// Hide		bool		`json:"hideNote"`
	Workouts []Workout `json:"workout"`
}
