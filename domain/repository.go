package domain

type Repository interface {
	Get() map[string]*Student
	Put(student *Student) (bool, error)
}
