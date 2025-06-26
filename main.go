// main.go
package main

import shallowcopy "deep_shallowcopy/shallowCopy"

func main() {
	var copia shallowcopy.ShallowCopyInterface = shallowcopy.ShallowCopy{}

	copia.ExampleArrays()
	copia.ExampleMaps()
	copia.ExampleStructs()
}
