package Prototype

import "fmt"

type Folder struct {
	children []InterfaceNode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() InterfaceNode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []InterfaceNode
	for _, i := range f.children {
		copyItem := i.clone()
		tempChildren = append(tempChildren, copyItem)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
