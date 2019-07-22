package model

// Event struct is the structure of data
type Event struct {
	Name      string `validate:"required" form:"name"`
	Device    string `validate:"required" form:"device"`
	Referer   string `form:"referer"`
	UserAgent string `form:"user_agent"`
	CreatedAt int    `form:"created_at"`
	UserData  User   `form:"user_data"`
}

// User struct is the structure of user data in event structure
type User struct {
	ID    string
	Phone string
	Name  string
	More  string
}
