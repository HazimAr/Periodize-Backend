package model

type Workout struct {
	Name       string `json:"workoutName"`
	Descripion string `json:"workoutDescription"`
	Type       string `json:"type"`
	Lifts      []Lift `json:"lifts"`
}
