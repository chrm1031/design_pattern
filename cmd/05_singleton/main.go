package singleton

import "fmt"

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

/*
実行結果
Creating single instance now.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
*/
