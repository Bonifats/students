package user

type User struct {
	Id      int
	Name    string
	Age     int
	Friends []int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}
