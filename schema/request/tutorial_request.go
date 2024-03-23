package request

type AddTutorial struct {
	TutorialTypeId string `validate:"required" json:"tutorialTypeId"`
	Title          string `validate:"required" json:"title"`
	Sequence       int8   `json:"sequence"`
	Keywords       string `json:"keywords"`
	Description    string `json:"description"`
}
