package model

type BuildLog struct {
	ID        string `json:"id"`
	Repo      string `json:"repo"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Log       string `json:"log"`
}
