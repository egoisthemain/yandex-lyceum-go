package main

import "fmt"

type Animal interface {
	MakeSound() string
	GetName() string
	GetInfo() string
}

type animal struct {
	name    string
	species string
	age     int
	sound   string
}

func NewAnimal(name, species string, age int, sound string) Animal {
	return &animal{name: name, species: species, age: age, sound: sound}
}

func (ani *animal) MakeSound() string {
	return ani.sound
}

func (ani *animal) GetName() string {
	return ani.name
}

func (ani *animal) GetInfo() string {
	return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", ani.name, ani.species, ani.age)
}

type ZooKeeper struct {
}

func (keeper *ZooKeeper) Feed(animal Animal) {
	fmt.Printf("Смотритель зоопарка кормит %s. %s!", animal.GetName(), animal.MakeSound())
}

func ZooShow(animals []Animal) {
	for _, ani := range animals {
		fmt.Println(ani.GetInfo())
		fmt.Println(ani.MakeSound())
	}
}
