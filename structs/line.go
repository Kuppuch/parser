package structs

import "gorm.io/gorm"

type Line struct {
	gorm.Model
	PersonalAccount string // 900046403
	Name            string // Сорокоумов С.М.
	Address         string // Омутское с., Молодежная дом7 кв2
	AccrualPeriod   string // 0519
	Count           string // 468.47
	Number          string // 3302372357 СХВ
	Testimony       string // 65.0000
}
