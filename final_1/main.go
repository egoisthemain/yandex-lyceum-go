package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintQueue(queue []string) {
	for i, name := range queue {
		fmt.Printf("%d. %s\n", i+1, name)
	}
}

func main() {
	var currentAction string
	var freePlaces = 5

	reader := bufio.NewReader(os.Stdin)

	places := []string{"-", "-", "-", "-", "-"}

	for {
		currentAction, _ = reader.ReadString('\n')
		currentAction = strings.TrimSpace(currentAction)
		if currentAction == "" {
			continue
		}
		//fmt.Printf("currentAction: %s\n", currentAction)
		if currentAction == "конец" {
			PrintQueue(places)
			break
		} else if currentAction == "очередь" {
			PrintQueue(places)
			continue
		} else if currentAction == "количество" {
			fmt.Printf("Осталось свободных мест: %d\n", freePlaces)
			fmt.Printf("Всего человек в очереди: %d\n", 5-freePlaces)
			continue
		}
		nameAndPlace := strings.Fields(currentAction)
		pos, _ := strconv.Atoi(nameAndPlace[1])
		//fmt.Println(nameAndPlace, len(nameAndPlace))
		if len(nameAndPlace) != 2 {
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", pos)
			continue
		} else if pos < 1 || pos > 5 {
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", pos)
			continue
		} else if freePlaces == 0 {
			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", pos)
			continue
		}
		index, _ := strconv.Atoi(nameAndPlace[1])
		if places[index-1] != "-" {
			fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", pos)
			continue
		}
		places[index-1] = nameAndPlace[0]
		freePlaces -= 1
	}

}
