package Prototype

import "fmt"

func PrototypeMain() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []InterfaceNode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []InterfaceNode{folder1, file2, file3},
		name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")

}
