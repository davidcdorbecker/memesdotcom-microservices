package domain

type User struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	Status      string `json:"status,omitempty"`
	DateCreated string `json:"date_created"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
