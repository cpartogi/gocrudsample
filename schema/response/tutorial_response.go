package response

type TutorialDetail struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	TutorialType string `json:"tutorialType"`
	Keywords     string `json:"keywords"`
	Sequence     int8   `json:"sequence"`
	Description  string `json:"description"`
	LastUpdate   string `json:"lastUpdate"`
}

type TutorialTypes struct {
	Id       string `json:"id"`
	TypeName string `json:"typeName"`
}

type TutorialList struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	TutorialType string `json:"tutorialType"`
}
