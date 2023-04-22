package app

import (
	"errors"
	"fmt"
	"io"
	"log"
	"students/domain"
)

type App struct {
	Repository domain.Repository
}

func (a *App) Run() {
	defer a.storeOutput()

	for {
		newStudent, err := a.inputData()
		if err != nil {
			if err != io.EOF {
				log.Print(err)
				continue
			}

			break
		}

		_, err = a.storeData(newStudent)
		if err != nil {
			log.Print(err)
			continue
		}
	}
}

func (a *App) inputData() (*domain.Student, error) {
	fmt.Print("Введите данные (Имя, возраст, курс) через пробел или нажмите `ctrl+d` для завершения: ")

	var age, grade int
	var name string

	_, err := fmt.Scanf("%s %d %d", &name, &age, &grade)
	if err == io.EOF {
		return nil, err
	}

	if name <= "" || age <= 0 || grade <= 0 {
		return nil, errors.New("некорректные данные")
	}

	return &domain.Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}, nil
}

func (a *App) storeData(student *domain.Student) (bool, error) {
	if _, err := a.Repository.Put(student); err != nil {
		return false, err
	}

	return true, nil
}

func (a *App) storeOutput() {
	fmt.Println("\nСтуденты из хранилища:")
	students := a.Repository.Get()

	counter := 0

	for _, student := range students {
		counter++
		fmt.Printf("\t%d) %s %d %d\n", counter, student.Name, student.Age, student.Grade)
	}
}
