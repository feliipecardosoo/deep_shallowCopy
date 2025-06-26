package shallowcopy

import "fmt"

func (s ShallowCopy) ExampleArrays() {
	// Exemplo de cópia rasa com slice/array
	sliceOriginal := []int{1, 2, 3}
	sliceCopia := sliceOriginal // shallow copy, ambos apontam para o mesmo array subjacente
	sliceCopia[0] = 100
	fmt.Println(sliceOriginal) // Output: [100 2 3] (original também foi modificado)
}

func (s ShallowCopy) ExampleMaps() {
	// Exemplo de cópia rasa com maps
	mapOriginal := map[string]int{"a": 1, "b": 2}
	mapCopia := mapOriginal // shallow copy, ambos apontam para o mesmo map
	mapCopia["a"] = 100
	fmt.Println(mapOriginal) // Output: map[a:100 b:2] (original também foi modificado)
}

func (s ShallowCopy) ExampleStructs() {
	// Exemplo de cópia rasa com structs
	// Structs que contêm slices, maps ou outros tipos referenciados também
	type Pessoa struct {
		Nome    string
		Hobbies []string
	}

	p1 := Pessoa{Nome: "Ana", Hobbies: []string{"Ler", "Cozinhar"}}
	p2 := p1 // shallow copy

	p2.Hobbies[0] = "Correr"
	fmt.Println(p1.Hobbies) // Output: [Correr Cozinhar] (hobbies de p1 também foram modificados)
}
