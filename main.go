// main.go
package main

import (
	deepcopy "deep_shallowcopy/deepCopy"
	shallowcopy "deep_shallowcopy/shallowCopy"
)

func main() {
	var copiaShallow shallowcopy.ShallowCopyInterface = shallowcopy.ShallowCopy{}

	copiaShallow.ExampleArrays()
	copiaShallow.ExampleMaps()
	copiaShallow.ExampleStructs()

	var copiaDeep deepcopy.DeepCopyInterface = deepcopy.DeepCopyStruct{}

	copiaDeep.ExampleSlices()
	copiaDeep.ExampleMaps()
	copiaDeep.ExampleStructs()
	copiaDeep.ExampleJSON()

}
