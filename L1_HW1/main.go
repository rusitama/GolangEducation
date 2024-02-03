package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	os.Args = []string{"program_name", "1.txt"} //заполняем слайс данными

	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	fileName := filepath.Base(filePth)
	fileExt := filepath.Ext(fileName)

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
