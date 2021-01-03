package domain

type User struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

//username, displayname,
