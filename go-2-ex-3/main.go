package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := make(map[int]float32);

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 117);
	// TODO: add one
	modules[254] = 245;
	// TODO: replace one
	modules[104] = 200;
	fmt.Println(modules)
}
