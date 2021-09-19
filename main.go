package main

import (
	"fmt"
	"github.com/Kuppuch/parser/dataBase"
	"github.com/Kuppuch/parser/mongo"
	"github.com/Kuppuch/parser/readFile"
	"os"
	"time"
)

func main() {
	err := mongo.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = dataBase.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//start := time.Now()
	//err = readFile.ReadFileByLinePostgres()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//fmt.Println("Время выполнения postgres:", time.Since(start))

	start := time.Now()
	err = readFile.ReadFileByLineMongo()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Время выполнения mongoDB:", time.Since(start))
}
