package main

import (
	"students/app"
	"students/repository/storage"
)

func main() {
	application := &app.App{}

	application.Repository = storage.NewStudentStorage()

	application.Run()
}
