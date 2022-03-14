package front

type Session struct {
}

type RegisterRequest struct {
	Name         string `json:"name" xml:"name" form:"name"`
	Email        string `json:"email" xml:"email" form:"email"`
	Position     string `json:"position" xml:"position" form:"position"`
	Presentation string `json:"presentation" xml:"presentation" form:"presentation"`
	HireDate     string `json:"hire_date" xml:"hire_date" form:"hire_date"`
	Password     string `json:"password" xml:"password" form:"password"`
}

type DeleteRequest struct {
	Password string `json:"password" xml:"password" form:"password"`
}

type SearchRequest struct {
	Q string `query:"q"`
}
