package todo

type User struct {
	ID        uint `gorm:"primary_key"`
	FirstName string
	LastName  string
	Email     string
	Todos     []*Todo
}

type UserRepository interface {
	FindUsers() ([]User, error)
	FindUser(id int64) (*User, error)
	CreateUser(t User) (*User, error)
	UpdateUser(id int64, t User) (*User, error)
	DeleteUser(t User) (*User, error)
}
