package deepcopy

import (
	"encoding/json"
	"fmt"
)

// Definição da struct Pessoa no nível de pacote
type Pessoa struct {
	Nome   string
	Scores []int
	Props  map[string]string
}

// ExampleSlices - exemplo de deep copy com slices
func (s DeepCopyStruct) ExampleSlices() {
	sliceOriginal := []int{1, 2, 3}
	sliceCopia := make([]int, len(sliceOriginal))
	copy(sliceCopia, sliceOriginal)
	sliceCopia[0] = 100
	fmt.Println(sliceOriginal) // [1 2 3]
}

// ExampleMaps - exemplo de deep copy com maps
func (s DeepCopyStruct) ExampleMaps() {
	mapOriginal := map[string]int{"a": 1, "b": 2}
	mapCopia := make(map[string]int, len(mapOriginal))
	for k, v := range mapOriginal {
		mapCopia[k] = v
	}
	mapCopia["a"] = 100
	fmt.Println(mapOriginal) // map[a:1 b:2]
}

// ExampleStructs - exemplo de deep copy com structs
func (s DeepCopyStruct) ExampleStructs() {
	p1 := Pessoa{
		Nome:   "Ana",
		Scores: []int{10, 20},
		Props:  map[string]string{"a": "x"},
	}
	p2 := deepCopyPessoa(p1)

	p2.Scores[0] = 100
	p2.Props["a"] = "y"

	fmt.Println(p1.Scores, p1.Props) // [10 20] map[a:x]
	fmt.Println(p2.Scores, p2.Props) // [100 20] map[a:y]
}

// Função auxiliar para deep copy de Pessoa
func deepCopyPessoa(orig Pessoa) Pessoa {
	copia := orig // copia superficial

	// Deep copy do slice
	copia.Scores = make([]int, len(orig.Scores))
	copy(copia.Scores, orig.Scores)

	// Deep copy do map
	copia.Props = make(map[string]string, len(orig.Props))
	for k, v := range orig.Props {
		copia.Props[k] = v
	}

	return copia
}

// ExampleJSON - exemplo usando serialização JSON (deep copy)
func (s DeepCopyStruct) ExampleJSON() {
	type Car struct {
		Brand  string
		Model  string
		Owners []string
		Engine struct {
			Capacity int
			Power    int
		}
	}

	car1 := Car{
		Brand:  "Tesla",
		Model:  "Model X",
		Owners: []string{"Alice", "Bob"},
		Engine: struct {
			Capacity int
			Power    int
		}{Capacity: 100, Power: 300},
	}
	var car2 Car
	byt, _ := json.Marshal(car1)
	json.Unmarshal(byt, &car2)

	car2.Owners[0] = "John"
	car2.Engine.Capacity = 200

	fmt.Println("Car1:", car1)
	fmt.Println("Car2:", car2)
}
