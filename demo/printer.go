package demo

type Person struct {
	Name    string
	Age     int
	realAge int
}

func NewPerson(name string, age int, realAge int) *Person {
	return &Person{Name: name, Age: age, realAge: realAge}
}

func Add(person *Person) {
	//对象传递需要使用指针
	person.Age += 10
	person.realAge += 10
}

func (person *Person) Add() {
	//函数与结构体关联
	person.Age += 10
	person.realAge += 10
}
