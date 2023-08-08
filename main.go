package main

import (
	"BuilderDesignPattern/models"
	"time"
)

func main() {
	/*
		TODO:
		 - Build PDF format
		 - Implement request
		 - Save in S3 bucket
	*/
	builder := models.NewPDFFileBuilder().
		SetTitle("Test").
		SetContent("hola Mundo!").
		SetAuthor("JorgeAndrew").
		SetFontSize(10).
		SetMargings(1.1)

	err := builder.Build()
	if err != nil {
		panic(err)
	}
	time.Sleep(30 * time.Second)
}
