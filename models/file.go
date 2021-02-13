package models

type FilesResult struct {
	Files []File
}

type File struct {
	FileType                  string `json:"fileType"`
	FileOwner                 string `json:"fileOwner"`
	CreationTimestamp         int64  `json:"creationTimestamp"`
	LastUpdatedTimestamp      int64  `json:"lastUpdatedTimestamp"`
	Size                      string `json:"size"`
	MimeType                  string `json:"mimeType"`
	FileName                  string `json:"fileName"`
	MainFileName              string `json:"mainFileName"`
	DirectURL                 string `json:"directUrl"`
	IsCurrentUserOwnerInDrive bool   `json:"isCurrentUserOwnerInDrive"`
	FileKey                   string `json:"fileKey"`
	Key                       string `json:"key"`
	BoxKey                    string `json:"boxKey"`
	GmailAPIFileID            string `json:"gmailAPIFileId"`
	GmailMessageID            string `json:"gmailMessageId"`
	LastSavedTimestamp        int64  `json:"lastSavedTimestamp"`
}
