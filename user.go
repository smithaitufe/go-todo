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
	FindUser(id uint) (*User, error)
	CreateUser(u User) (*User, error)
	UpdateUser(u User) (*User, error)
	DeleteUser(u User) error
}
