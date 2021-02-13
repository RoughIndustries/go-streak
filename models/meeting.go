package models

type MeetingResults struct {
	Results []Meeting `json:"results"`
}

type Meeting struct {
	BoxKey             string `json:"boxKey"`
	PipelineKey        string `json:"pipelineKey"`
	OrgKey             string `json:"orgKey"`
	CreatorKey         string `json:"creatorKey"`
	IsDraft            bool   `json:"isDraft"`
	CreationDate       int64  `json:"creationDate"`
	Key                string `json:"key"`
	MeetingType        string `json:"meetingType"`
	StartTimestamp     int64  `json:"startTimestamp"`
	Duration           int    `json:"duration"`
	Notes              string `json:"notes"`
	LastSavedTimestamp int64  `json:"lastSavedTimestamp"`
}
