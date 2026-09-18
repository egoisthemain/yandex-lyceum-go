package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

func currentDayOfTheWeek() string {
	currTime := time.Now()
	DayNumber := currTime.Day()
	switch DayNumber {
	case 1:
		return "Понедельник"
	case 2:
		return "Вторник"
	case 3:
		return "Среда"
	case 4:
		return "Четверг"
	case 5:
		return "Пятница"
	case 6:
		return "Суббота"
	case 7:
		return "Воскресенье"
	default:
		return ""
	}
}

func dayOrNight() string {
	currTime := time.Now()
	hours := currTime.Hour()
	if hours >= 10 && hours <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	DayNumber := time.Now().Day()
	if DayNumber <= 5 {
		return 5 - DayNumber
	} else {
		switch DayNumber {
		case 6:
			return 6
		case 7:
			return 5
		default:
			return 0
		}
	}
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	currentDay := currentDayOfTheWeek()
	return strings.EqualFold(currentDay, answer)
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}
	return strings.EqualFold(dayOrNight(), answer), nil

}