package app

import (
	"errors"
	"fmt"
	"io"
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
				fmt.Print(err)
				continue
			}

			break
		}

		_, err = a.storeData(newStudent)
		if err != nil {
			fmt.Print(err)

			continue
		}
	}
}

func (a *App) inputData() (*domain.Student, error) {
	fmt.Print("Введите данные (Имя, возраст, курс) через пробел или нажмите `ctrl+d` для завершения: ")

	var age, grade int
	var name string

	_, err := fmt.Scanf("%s %d %d", &name, &age, &grade)
	if err != nil {
		return nil, err
	}

	if name <= "" || age <= 0 || grade <= 0 {
		return nil, errors.New(fmt.Sprintf("Некорректные данные %s%d%d", name, age, grade))
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

	count := 0

	for _, student := range students {
		count++
		fmt.Printf("\t%d) %s %d %d\n", count, student.Name, student.Age, student.Grade)
	}
}
