package models

type TaskResults struct {
	Results []Task
}

type Task struct {
	IsDraft                  bool   `json:"isDraft"`
	BoxKey                   string `json:"boxKey"`
	PipelineKey              string `json:"pipelineKey"`
	CreatorKey               string `json:"creatorKey"`
	CreationDate             int64  `json:"creationDate"`
	LastStatusChangeDate     int64  `json:"lastStatusChangeDate"`
	DueDate                  int64  `json:"dueDate"`
	SortOrder                string `json:"sortOrder"`
	Text                     string `json:"text"`
	Status                   string `json:"status"`
	ReminderStatus           string `json:"reminderStatus"`
	AssignedToSharingEntries []struct {
		DisplayName string `json:"displayName"`
		FullName    string `json:"fullName"`
		Email       string `json:"email"`
		Image       string `json:"image"`
		UserKey     string `json:"userKey"`
	} `json:"assignedToSharingEntries"`
	CreatorSharingEntry struct {
		DisplayName string `json:"displayName"`
		FullName    string `json:"fullName"`
		Email       string `json:"email"`
		Image       string `json:"image"`
		UserKey     string `json:"userKey"`
	} `json:"creatorSharingEntry"`
	Key                string `json:"key"`
	LastSavedTimestamp int64  `json:"lastSavedTimestamp"`
	Actions            []struct {
		Type    string `json:"type"`
		Value   string `json:"value"`
		Contact struct {
			TeamKey           string        `json:"teamKey"`
			EmailAddresses    []string      `json:"emailAddresses"`
			PhoneNumbers      []interface{} `json:"phoneNumbers"`
			Addresses         []interface{} `json:"addresses"`
			Domains           []interface{} `json:"domains"`
			NormalizedDomains []interface{} `json:"normalizedDomains"`
			ContactLinks      []interface{} `json:"contactLinks"`
			OrgLinks          []struct {
				Key               string `json:"key"`
				IsMagicLinked     bool   `json:"isMagicLinked"`
				CreationTimestamp int64  `json:"creationTimestamp"`
			} `json:"orgLinks"`
			LastSavedUserKey   string `json:"lastSavedUserKey"`
			CreatorKey         string `json:"creatorKey"`
			CreationDate       int64  `json:"creationDate"`
			Key                string `json:"key"`
			VersionTimestamp   int64  `json:"versionTimestamp"`
			LastSavedTimestamp int64  `json:"lastSavedTimestamp"`
		} `json:"contact"`
	} `json:"actions,omitempty"`
}
