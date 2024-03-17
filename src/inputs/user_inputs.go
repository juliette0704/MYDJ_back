package inputs

type UserInitInput struct {
	Firstname  string `json:"firstname" binding:"required"`
	Lastname   string `json:"lastname" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	UID        string `json:"uid"`
	ADusername string `json:"AD_username"`
}

type UserModifyInput struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

type UserCredentialInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
