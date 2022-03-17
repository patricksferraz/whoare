package front

type Session struct {
}

type RegisterRequest struct {
	Name         string  `json:"name" xml:"name" form:"name"`
	Email        string  `json:"email" xml:"email" form:"email"`
	Position     string  `json:"position" xml:"position" form:"position"`
	Presentation string  `json:"presentation" xml:"presentation" form:"presentation"`
	HireDate     string  `json:"hire_date" xml:"hire_date" form:"hire_date"`
	Password     string  `json:"password" xml:"password" form:"password"`
	Skills       []Skill `json:"skills" xml:"skills" form:"skills"`
}

type Skill struct {
	Name string `json:"name" xml:"name" form:"name"`
	XP   int    `json:"xp" xml:"xp" form:"xp"`
	Note string `json:"note" xml:"note" form:"note"`
}

type DeactivateRequest struct {
	TaminationDate string `json:"termination_date" xml:"termination_date" form:"termination_date"`
	Password       string `json:"password" xml:"password" form:"password"`
}

type SearchRequest struct {
	Q string `query:"q"`
}

type ActivateRequest struct {
	HireDate string `json:"hire_date" xml:"hire_date" form:"hire_date"`
	Password string `json:"password" xml:"password" form:"password"`
}
