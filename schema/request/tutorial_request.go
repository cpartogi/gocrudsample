package request

type AddTutorial struct {
	TutorialTypeId string `json:"tutorialTypeId"`
	Title          string `json:"title"`
	Sequence       int8   `json:"sequence"`
	Keywords       string `json:"keywords"`
	Description    string `json:"description"`
}

type PatchTutorial struct {
	Title string `json:"title"`
}
