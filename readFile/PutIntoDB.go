package readFile

import "github.com/Kuppuch/parser/dataBase"

func PutIntoDB() {
	for {
		select {
		case line := <-Wrt:
			dataBase.DB.Create(&line)
		case <-Done:
			return
		}
	}
}
