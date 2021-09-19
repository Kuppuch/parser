package readFile

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Kuppuch/parser/dataBase"
	"github.com/Kuppuch/parser/mongo"
	"github.com/Kuppuch/parser/structs"
	"io"
	"os"
	"regexp"
	"strings"
)

var Wrt = make(chan structs.Line)
var Done = make(chan bool)

func ReadFileByLinePostgres() error {
	file, err := os.Open("parse.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Println(file.Name())

	reader := bufio.NewReader(file)
	cnt := 0
	for {
		cnt++
		if cnt%100000 == 0 {
			fmt.Println(cnt)
		}
		lineStr, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		subLine := strings.Split(string(lineStr), ";")
		line, err := LineParser(subLine)
		if err != nil {
			return err
		}
		dataBase.DB.Create(line)
	}
	fmt.Printf("Вставлено всего записей в базу: %v \n", cnt)
	return nil
}

func ReadFileByLineMongo() error {
	file, err := os.Open("parse.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Println(file.Name())

	reader := bufio.NewReader(file)
	cnt := 0
	for {
		cnt++
		if cnt%100000 == 0 {
			fmt.Println(cnt)
		}
		lineStr, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		subLine := strings.Split(string(lineStr), ";")
		line, err := LineParser(subLine)
		if err != nil {
			return err
		}
		_, err = mongo.Collection.InsertOne(context.Background(), line)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Вставлено всего записей в базу: %v \n", cnt)
	return nil
}

func LineParser(subLine []string) (*structs.Line, error) {
	line := structs.Line{}
	for _, v := range subLine {
		// Парсим PersonalAccount
		match, _ := regexp.MatchString(`\d{9}`, v)
		if match {
			line.PersonalAccount = v
			continue
		}

		// Парсим Name
		match, _ = regexp.MatchString(`^[а-я А-Я]+\s[а-я А-Я].\s{0,1}[а-я А-Я].{0,1}\D$`, v)
		if match {
			line.Name = v
			continue
		}

		// Парсим Address
		match, _ = regexp.MatchString(`^[а-я А-Я]+.{0,2}[а-я А-Я]+.{0,2}\s[а-я А-Я]{0,1}\d{0,3}.{0,1}\s{0,}\d{1,}$`, v)
		if match {
			line.Address = v
			continue
		}
		// Парсим AccrualPeriod
		match, _ = regexp.MatchString(`^\d\d\d\d$`, v)
		if match {
			line.AccrualPeriod = v
			continue
		}
		if line.Count == "" {
			// Парсим Count
			match, _ = regexp.MatchString(`^\d{1,}.{1,}$`, v)
			if match {
				line.Count = v
				continue
			}
		}

		// Парсим Testimony
		match, _ = regexp.MatchString(`^\d{1,}.{1,}$`, v)
		if match {
			line.Testimony = v
			continue
		}
	}
	return &line, nil
}
