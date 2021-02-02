package models

type PipelinesResult struct {
	Pipelines []Pipeline
}

type Pipeline struct {
	TeamKey                 string `json:"teamKey"`
	CreatorKey              string `json:"creatorKey"`
	Icon                    string `json:"icon"`
	Name                    string `json:"name"`
	OrgWide                 bool   `json:"orgWide"`
	SharingRestrictedToOrg  bool   `json:"sharingRestrictedToOrg"`
	TeamWide                bool   `json:"teamWide"`
	SharingRestrictedToTeam bool   `json:"sharingRestrictedToTeam"`
	Fields                  []struct {
		Name                 string `json:"name"`
		Key                  string `json:"key"`
		Type                 string `json:"type"`
		LastUpdatedTimestamp int64  `json:"lastUpdatedTimestamp"`
		DropdownSettings     struct {
			Items []interface{} `json:"items"`
		} `json:"dropdownSettings,omitempty"`
		TagSettings struct {
			Tags []struct {
				Key string `json:"key"`
				Tag string `json:"tag"`
			} `json:"tags"`
		} `json:"tagSettings,omitempty"`
	} `json:"fields"`
	Stages               map[string]Stages `json:"stages"`
	StageOrder           []string          `json:"stageOrder"`
	StageColorTheme      string            `json:"stageColorTheme"`
	CreationTimestamp    int64             `json:"creationTimestamp"`
	LastUpdatedTimestamp int64             `json:"lastUpdatedTimestamp"`
	CustomPermissionSets struct {
	} `json:"customPermissionSets"`
	DefaultPermissionSetName  string        `json:"defaultPermissionSetName"`
	RoundRobinAssignmentUsers []interface{} `json:"roundRobinAssignmentUsers"`
	ACLEntries                []struct {
		PermissionSetName string `json:"permissionSetName"`
		DisplayName       string `json:"displayName"`
		FullName          string `json:"fullName"`
		Email             string `json:"email"`
		Image             string `json:"image"`
		UserKey           string `json:"userKey"`
	} `json:"aclEntries"`
	PipelineKey        string `json:"pipelineKey"`
	Key                string `json:"key"`
	CreationSourceType string `json:"creationSourceType"`
	BoxCountHint       int    `json:"boxCountHint"`
	BoxCount           int    `json:"boxCount"`
	LastSavedTimestamp int64  `json:"lastSavedTimestamp"`
}

type Stages struct {
	Name  string `json:"name"`
	Key   string `json:"key"`
	Color struct {
		ForegroundColor string `json:"foregroundColor"`
		BackgroundColor string `json:"backgroundColor"`
	} `json:"color"`
	BoxCount         int  `json:"boxCount"`
	IsColorFromTheme bool `json:"isColorFromTheme"`
}
