package main

import "fmt"

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	// fmt.Println("Counter selesai")

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("Counter selesai")

	names := []string{"Jaka", "Kelana", "Umbara"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, name := range names {
		fmt.Println("Index", i, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
