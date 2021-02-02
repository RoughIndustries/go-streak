package models

type BoxesResult struct {
	Boxes []Box
}

type Box struct {
	LastSavedTimestamp              int64  `json:"lastSavedTimestamp"`
	BoxKey                          string `json:"boxKey"`
	PipelineKey                     string `json:"pipelineKey"`
	CreatorKey                      string `json:"creatorKey"`
	CreationTimestamp               int64  `json:"creationTimestamp"`
	LastUpdatedTimestamp            int64  `json:"lastUpdatedTimestamp"`
	LastEmailUpdatedTimestamp       int64  `json:"lastEmailUpdatedTimestamp,omitempty"`
	LastStageChangeTimestamp        int64  `json:"lastStageChangeTimestamp"`
	LastEmailFrom                   string `json:"lastEmailFrom,omitempty"`
	FirstEmailFrom                  string `json:"firstEmailFrom,omitempty"`
	TotalNumberOfEmails             int    `json:"totalNumberOfEmails"`
	TotalNumberOfSentEmails         int    `json:"totalNumberOfSentEmails"`
	TotalNumberOfLastSentEmailViews int    `json:"totalNumberOfLastSentEmailViews"`
	TotalNumberOfSentEmailsViews    int    `json:"totalNumberOfSentEmailsViews"`
	TotalNumberOfReceivedEmails     int    `json:"totalNumberOfReceivedEmails"`
	FirstEmailTimestamp             int64  `json:"firstEmailTimestamp,omitempty"`
	FirstEmailReceivedTimestamp     int64  `json:"firstEmailReceivedTimestamp,omitempty"`
	LastEmailReceivedTimestamp      int64  `json:"lastEmailReceivedTimestamp,omitempty"`
	Name                            string `json:"name"`
	Notes                           string `json:"notes"`
	AssignedToSharingEntries        []struct {
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
	FollowerSharingEntries []struct {
		DisplayName string `json:"displayName"`
		FullName    string `json:"fullName"`
		Email       string `json:"email"`
		Image       string `json:"image"`
		UserKey     string `json:"userKey"`
	} `json:"followerSharingEntries"`
	StageKey      string        `json:"stageKey"`
	FollowerKeys  []string      `json:"followerKeys"`
	LinkedBoxKeys []interface{} `json:"linkedBoxKeys"`
	Contacts      []struct {
		IsStarred   bool   `json:"isStarred"`
		IsAutoboxed bool   `json:"isAutoboxed"`
		Key         string `json:"key"`
	} `json:"contacts,omitempty"`
	Organizations []struct {
		IsStarred   bool   `json:"isStarred"`
		IsAutoboxed bool   `json:"isAutoboxed"`
		Key         string `json:"key"`
	} `json:"organizations,omitempty"`
	EmailAddressesAutoExtracted  []string      `json:"emailAddressesAutoExtracted"`
	EmailAddressesBlacklist      []interface{} `json:"emailAddressesBlacklist"`
	EmailAddresses               []string      `json:"emailAddresses"`
	TaskCompleteCount            int           `json:"taskCompleteCount"`
	TaskIncompleteCount          int           `json:"taskIncompleteCount"`
	TaskOverdueCount             int           `json:"taskOverdueCount"`
	TaskAssigneeKeySet           []string      `json:"taskAssigneeKeySet"`
	OverdueTaskAssigneeKeySet    []interface{} `json:"overdueTaskAssigneeKeySet"`
	IncompleteTaskAssigneeKeySet []interface{} `json:"incompleteTaskAssigneeKeySet"`
	TaskAssigneeSharingEntrySet  []struct {
		DisplayName string `json:"displayName"`
		FullName    string `json:"fullName"`
		Email       string `json:"email"`
		Image       string `json:"image"`
		UserKey     string `json:"userKey"`
	} `json:"taskAssigneeSharingEntrySet"`
	OverdueTaskAssigneeSharingEntrySet    []interface{}     `json:"overdueTaskAssigneeSharingEntrySet"`
	IncompleteTaskAssigneeSharingEntrySet []interface{}     `json:"incompleteTaskAssigneeSharingEntrySet"`
	TaskPercentageComplete                float64           `json:"taskPercentageComplete,omitempty"`
	TaskTotal                             int               `json:"taskTotal"`
	CallLogCount                          int               `json:"callLogCount"`
	MeetingNotesCount                     int               `json:"meetingNotesCount"`
	MostRecentCallLogTimestamp            int64             `json:"mostRecentCallLogTimestamp,omitempty"`
	FirstCallLogTimestamp                 int64             `json:"firstCallLogTimestamp,omitempty"`
	TotalCallLogDuration                  int               `json:"totalCallLogDuration"`
	TotalMeetingNotesDuration             int               `json:"totalMeetingNotesDuration"`
	LastTouchpointTimestamp               int64             `json:"lastTouchpointTimestamp,omitempty"`
	LastTouchpointType                    string            `json:"lastTouchpointType,omitempty"`
	FollowerCount                         int               `json:"followerCount"`
	CommentCount                          int               `json:"commentCount"`
	GmailThreadCount                      int               `json:"gmailThreadCount"`
	FileCount                             int               `json:"fileCount"`
	Key                                   string            `json:"key"`
	Freshness                             float64           `json:"freshness"`
	FirstEmailSentTimestamp               int64             `json:"firstEmailSentTimestamp,omitempty"`
	LastEmailSentTimestamp                int64             `json:"lastEmailSentTimestamp,omitempty"`
	EarliestOverdueTaskDueDate            int64             `json:"earliestOverdueTaskDueDate,omitempty"`
	MostRecentLastSentEmailViewsTimestamp int64             `json:"mostRecentLastSentEmailViewsTimestamp,omitempty"`
	MostRecentSentEmailsViewsTimestamp    int64             `json:"mostRecentSentEmailsViewsTimestamp,omitempty"`
	LastCommentTimestamp                  int64             `json:"lastCommentTimestamp,omitempty"`
	MostRecentMeetingNotesTimestamp       int64             `json:"mostRecentMeetingNotesTimestamp,omitempty"`
	FirstMeetingNotesTimestamp            int64             `json:"firstMeetingNotesTimestamp,omitempty"`
	Fields                                map[string]Fields `json:"fields"`
	SoonestTaskDueDate                    int64             `json:"soonestTaskDueDate,omitempty"`
}

type Fields struct {
	FieldItems map[string]string
}
