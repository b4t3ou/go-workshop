package main

import "fmt"

type LifeWorks struct {
	Language string
	Name string
}

func (l LifeWorks) getName() string {
	return l.Name
}

type WorkAngel struct {
	Language string
	Name string
	LifeWorks
}

func (w WorkAngel) getName() string {
	return w.Name
}

func main() {
	company := WorkAngel{
		"en-GB",
		"WorkAngel",
		LifeWorks{"en-US", "LifeWorks"},
	}

	fmt.Println(company.Language)
	fmt.Println(company.getName())
	fmt.Println(company.LifeWorks.Language)
	fmt.Println(company.LifeWorks.getName())
}
