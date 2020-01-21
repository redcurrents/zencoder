package zencoder

type Notification struct {
	Job     *Job               `json:"job,omitempty"`
	Outputs []*OutputMediaFile `json:"outputs,omitempty"`
	Input   *InputMediaFile    `json:"input,omitempty"`
}

func (n *Notification) Errors() (mediaFileErrors []*MediaFileError) {
	for _, output := range n.Outputs {
		if outputErrs := output.Errors(); len(outputErrs) > 0 {
			mediaFileErrors = append(mediaFileErrors, outputErrs...)
		}
	}
	return
}
