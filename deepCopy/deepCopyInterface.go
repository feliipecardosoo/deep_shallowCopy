package deepcopy

type DeepCopyInterface interface {
	ExampleSlices()
	ExampleMaps()
	ExampleStructs()
	ExampleJSON()
}

type DeepCopyStruct struct{}
