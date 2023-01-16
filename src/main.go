package main

import (
	"fmt"

	"github.com/ief2i-florent/go-f22/src/cours"
)

func isAlive(be cours.Be) {
	if be.IsAlive() {
		fmt.Println("Il est vivant")
	}
}

func main() {

	//result, err := cours.Division(10, 5)
	//fmt.Println(*result, err)

	cours.ArgList("toto", "tata", "titi", "tutu")
	cours.While()

	/*
		cours.Master(func() string {
			time.Sleep(1 * time.Second)
			return "function1"
		}, func() string {
			time.Sleep(10 * time.Second)
			return "function2"
		}, func() string {
			time.Sleep(5 * time.Second)
			return "function3"
		}, func() string {
			time.Sleep(3 * time.Second)
			return "function4"
		})
	*/

	dog := cours.MakeAnimal("Florent", 34)
	user := cours.MakeUser("Florent", 34)
	user.Hello()
	fmt.Println(dog)
	fmt.Println(user)

	isAlive(dog)

}
