package models

type User struct {
	Email                    string `json:"email"`
	OrgKey                   string `json:"orgKey"`
	CreationTimestamp        int    `json:"creationTimestamp"`
	LastUpdatedTimestamp     int    `json:"lastUpdatedTimestamp"`
	LastSeenTimestamp        int    `json:"lastSeenTimestamp"`
	TimezoneID               string `json:"timezoneId"`
	GoogleProfileID          string `json:"googleProfileId"`
	GoogleProfilePhotoURL    string `json:"googleProfilePhotoUrl"`
	GoogleProfileLink        string `json:"googleProfileLink"`
	GoogleProfileFullName    string `json:"googleProfileFullName"`
	GoogleProfileFirstName   string `json:"googleProfileFirstName"`
	GoogleProfileLastName    string `json:"googleProfileLastName"`
	GoogleProfileGender      string `json:"googleProfileGender"`
	GoogleProfileLocale      string `json:"googleProfileLocale"`
	GoogleAnalyticsClientID  string `json:"googleAnalyticsClientId"`
	IsOauthComplete          bool   `json:"isOauthComplete"`
	UserKey                  string `json:"userKey"`
	DisplayName              string `json:"displayName"`
	Key                      string `json:"key"`
	TourID                   string `json:"tourId"`
	UserSettingsKey          string `json:"userSettingsKey"`
	StreakCalendarID         string `json:"streakCalendarId"`
	LastTrialLength          int    `json:"lastTrialLength"`
	OnTrialWithoutCreditCard bool   `json:"onTrialWithoutCreditCard"`
	CanceledTrial            bool   `json:"canceledTrial"`
	FirstOauthTimestamp      int    `json:"firstOauthTimestamp"`
	WantsTaskDigestEmail     bool   `json:"wantsTaskDigestEmail"`
	LastSavedTimestamp       int    `json:"lastSavedTimestamp"`
	UsedPlatforms            []struct {
		Platform       string `json:"platform"`
		FirstDateOfUse int    `json:"firstDateOfUse"`
		LastDateOfUse  int    `json:"lastDateOfUse"`
	} `json:"usedPlatforms"`
}
