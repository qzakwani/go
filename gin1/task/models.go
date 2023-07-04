package task

type Task struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
}
