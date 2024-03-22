package model

import "time"

type TutorialList struct {
	Id               string
	Title            string
	TutorialTypeName string
	Sequence         int8
}

type Tutorials struct {
	Id               string
	TutorialTypeId   string
	Keywords         string
	Sequence         int8
	Title            string
	Description      string
	CreatedBy        string
	CreatedAt        time.Time
	UpdatedBy        string
	UpdatedAt        time.Time
	DeletedBy        *string
	DeletedAt        *time.Time
	TutorialTypeName string
	TutorialType     TutorialTypes `pg:"rel:has-one,fk:tutorial_type_id"`
}

type TutorialTypes struct {
	Id        string
	TypeName  string
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
	DeletedBy *string
	DeletedAt *time.Time
}
