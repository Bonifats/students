package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	app := &App{}

	app.Repository = NewStudentStorage()

	app.Run()
}

type Student struct {
	Name  string
	Age   int
	Grade int
}

type StudentStorage map[string]*Student

type App struct {
	Repository StudentStorage
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

func (a *App) inputData() (*Student, error) {
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

	return &Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}, nil
}

func (a *App) storeData(student *Student) (bool, error) {
	if _, err := a.Repository.Put(student); err != nil {
		return false, err
	}

	return true, nil
}

func (a *App) storeOutput() {
	fmt.Println("\nСтуденты из хранилища:")
	students := a.Repository

	counter := 0

	for _, student := range students {
		counter++
		fmt.Printf("\t%d) %s %d %d\n", counter, student.Name, student.Age, student.Grade)
	}
}

func NewStudentStorage() StudentStorage {
	return make(map[string]*Student, 0)
}

func (ss StudentStorage) Get(name string) (*Student, error) {
	if exist := ss.contains(name); !exist {
		return nil, errors.New("студент с данным именем не существует")
	}

	return ss[name], nil
}

func (ss StudentStorage) Put(student *Student) (bool, error) {
	if exist := ss.contains(student.Name); exist {
		return false, errors.New("студент с данным именем уже существует")
	}

	ss[student.Name] = student

	return true, nil
}

func (ss StudentStorage) contains(name string) bool {
	_, ok := ss[name]

	return ok
}
