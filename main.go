package main

import (
	"fmt"
	"github.com/Kuppuch/parser/dataBase"
	"github.com/Kuppuch/parser/readFile"
	"os"
	"time"
)

func main() {
	err := dataBase.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	start := time.Now()
	err = readFile.ReadFileByLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Время выполнения:", time.Since(start))
}
