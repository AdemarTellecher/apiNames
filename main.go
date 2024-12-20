package main

import (
	goroutines "apiNames/src/routines"
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	AllNames, err := goroutines.Concurrency(50, 50)
	if err != nil {
		fmt.Println("Erro ao obter nomes...")
	}
	for i, Names := range AllNames {

		fmt.Println(i, Names)
	}
	fmt.Println("Quantidade total de nomes:", len(AllNames))
	fmt.Println("Tempo estimado", time.Since(t0))
}
