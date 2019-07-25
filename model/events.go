package model

// Event struct is the structure of data
type Event struct {
	Type      string `validate:"required" form:"type"`
	Category    string `validate:"required" form:"category"`
	Action   string `form:"action"`
	Name string `form:"action"`
	Value string `form:"value"`
	CreatedAt int    `form:"created_at"`
	UserData  User   `form:"user_data" validate:"required"`
	More EventMore `form:"more" validate:"required"`
}

// User struct is the structure of user data in event structure
type User struct {
	ID    string
	Phone string
	Name  string
	More  string
}

type EventMore struct {
	Referral string `validate:"required"`
}
