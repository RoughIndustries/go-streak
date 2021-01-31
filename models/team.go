package models

type Teams struct {
	Results []struct {
		CreationDate int    `json:"creationDate"`
		Creator      string `json:"creator"`
		Members      []struct {
			Role         string `json:"role"`
			InviteDate   int    `json:"inviteDate"`
			InviteStatus string `json:"inviteStatus"`
			OrgKey       string `json:"orgKey"`
			DisplayName  string `json:"displayName"`
			FullName     string `json:"fullName"`
			Email        string `json:"email"`
			Image        string `json:"image"`
			UserKey      string `json:"userKey"`
		} `json:"members"`
		Name                    string `json:"name"`
		SharingRestrictedToTeam bool   `json:"sharingRestrictedToTeam"`
		Key                     string `json:"key"`
		LastSavedTimestamp      int    `json:"lastSavedTimestamp"`
	} `json:"results"`
}
