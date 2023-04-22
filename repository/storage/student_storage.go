package storage

import (
	"errors"
	"fmt"
	"students/domain"
)

type StudentStorage struct {
	db map[string]*domain.Student
}

func NewStudentStorage() *StudentStorage {
	return &StudentStorage{
		db: make(map[string]*domain.Student, 0),
	}
}

func (ss *StudentStorage) Get() map[string]*domain.Student {
	return ss.db
}

func (ss *StudentStorage) Put(student *domain.Student) (bool, error) {
	if err := ss.contains(student); err != nil {
		return false, err
	}

	ss.db[student.Name] = student

	return true, nil
}

func (ss *StudentStorage) contains(student *domain.Student) error {
	if _, ok := ss.db[student.Name]; ok {
		return errors.New(fmt.Sprintf("Студент уже с данным именем добавлен"))
	}

	return nil
}
