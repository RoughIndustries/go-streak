package models

type ThreadsResult struct {
	Threads []Thread
}

type Thread struct {
	CreatorKey                    string        `json:"creatorKey"`
	BoxKey                        string        `json:"boxKey"`
	BoxKeys                       []string      `json:"boxKeys"`
	ThreadGmailID                 string        `json:"threadGmailId"`
	PipelineKey                   string        `json:"pipelineKey"`
	CreationTimestamp             int64         `json:"creationTimestamp"`
	LastUpdatedTimestamp          int64         `json:"lastUpdatedTimestamp"`
	GmailThreadKey                string        `json:"gmailThreadKey"`
	Key                           string        `json:"key"`
	IsRequestingUserOwner         bool          `json:"isRequestingUserOwner"`
	DoesRequestingUserHaveThread  bool          `json:"doesRequestingUserHaveThread"`
	RequestingUserThreadGmailID   string        `json:"requestingUserThreadGmailId"`
	LastEmailTimestamp            int64         `json:"lastEmailTimestamp"`
	NumberOfEmails                int           `json:"numberOfEmails"`
	Subject                       string        `json:"subject"`
	FileKeys                      []string      `json:"fileKeys"`
	AutoboxLinks                  []interface{} `json:"autoboxLinks"`
	EmailAddresses                []string      `json:"emailAddresses"`
	LastEmailFrom                 string        `json:"lastEmailFrom"`
	MessageSentTimestamps         []int64       `json:"messageSentTimestamps"`
	Names                         []string      `json:"names"`
	SenderEmailSendTimestampPairs []struct {
		Email         string `json:"email"`
		SendTimestamp int64  `json:"sendTimestamp"`
	} `json:"senderEmailSendTimestampPairs"`
	UnifiedMessageThreads []struct {
	} `json:"unifiedMessageThreads"`
	Sources []struct {
		UserKey    string `json:"userKey"`
		MessageID  string `json:"messageId"`
		ThreadID   string `json:"threadId"`
		HasOauth   bool   `json:"hasOauth"`
		Visibility string `json:"visibility"`
		Gid        string `json:"gid"`
	} `json:"sources"`
}
