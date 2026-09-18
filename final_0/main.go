package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	var date string
	var name, surname, thirdname string
	var pay1, pay2, pay3 float64
	const dateForm = "10.04.2005"

	fmt.Scanln(&date)
	fmt.Scanln(&name)
	fmt.Scanln(&surname)
	fmt.Scanln(&thirdname)
	fmt.Scanln(&pay1)
	fmt.Scanln(&pay2)
	fmt.Scanln(&pay3)

	dateParts := strings.Split(date, ".")
	// fmt.Println(dateParts)
	// fmt.Println(dateInts)
	dateTime, _ := time.Parse(time.DateOnly, fmt.Sprintf("%s-%s-%s", dateParts[2], dateParts[1], dateParts[0]))
	// fmt.Println(dateTime)
	dateNew := dateTime.Add(time.Hour * 24 * 15)
	date = dateNew.Format(time.DateOnly)
	dateParts = strings.Split(date, "-")
	date = fmt.Sprintf("%s.%s.%s", dateParts[2], dateParts[1], dateParts[0])
	// fmt.Println(dateNew)
	// fmt.Println(date)

	summ := pay1 + pay2 + pay3
	rubles, kopees := math.Modf(summ)
	kopees = math.Floor(kopees * 100)
	//fmt.Println(rubles, kopees)

	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\nДата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\nОбщая сумма выплат составит %.0f руб. %.0f коп.\n\nС уважением,\nГл. бух. Иванов А.Е.", surname, name, thirdname, date, rubles, kopees)
}
