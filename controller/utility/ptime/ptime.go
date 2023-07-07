package ptime

import (
	"errors"
	"strconv"
	"strings"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

func ToPersian(t time.Time) string {
	return ptime.New(t).Format("yyyy/MM/dd")
}

func ToGregorian(intime string) (time.Time, error) {
	var temp time.Time
	sp := strings.Split(intime, "/")
	if len(sp) != 3 {
		return temp, errors.New("invalid format data")
	} else if len(sp[0]) != 4 || len(sp[1]) != 2 || len(sp[2]) != 2 {
		return temp, errors.New("invalid format data")
	}
	day, err := strconv.Atoi(sp[2])
	if err != nil {
		return temp, errors.New("invalid format data")
	}
	m, err := strconv.Atoi(sp[1])
	if err != nil {
		return temp, errors.New("invalid format data")
	}
	var month ptime.Month
	switch m {
	case 1:
		month = ptime.Farvardin
	case 2:
		month = ptime.Ordibehesht
	case 3:
		month = ptime.Khordad
	case 4:
		month = ptime.Tir
	case 5:
		month = ptime.Mordad
	case 6:
		month = ptime.Shahrivar
	case 7:
		month = ptime.Mehr
	case 8:
		month = ptime.Aban
	case 9:
		month = ptime.Azar
	case 10:
		month = ptime.Dey
	case 11:
		month = ptime.Bahman
	case 12:
		month = ptime.Esfand
	default:
		return temp, errors.New("invalid number of month")

	}
	year, err := strconv.Atoi(sp[0])
	if err != nil {
		return temp, errors.New("invalid format data")
	}
	var pt ptime.Time = ptime.Date(year, month, day, 12, 59, 59, 0, ptime.Iran())
	if pt.Day() != day {
		return temp, errors.New("invalid data")
	}
	t1 := pt.Time()

	return t1, nil

}
