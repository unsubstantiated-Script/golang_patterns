package Singleton

import "fmt"

func MainSingleton() {
	fmt.Println("Hello singleton")
	for i := 0; i < 30; i++ {
		go getInstance()
	}
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
