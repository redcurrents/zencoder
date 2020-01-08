package zencoder

type Notification struct {
	Job     *Job               `json:"job,omitempty"`
	Outputs []*OutputMediaFile `json:"outputs,omitempty"`
	Input   *InputMediaFile    `json:"input,omitempty"`
}
