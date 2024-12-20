package goroutines

import (
	"apiNames/src/pages"
	"fmt"
	"log"
	"strconv"
)

func Concurrency(totalPaginas, Concurrency int) ([]string, error) {

	chSemaphore := make(chan struct{}, Concurrency)
	allNames := make([]string, 0)
	chNames := make(chan []string)

	go func() {
		for names := range chNames {
			if names != nil {
				allNames = append(allNames, names...)
			} else {
				fmt.Println("Erro ao adicionar os nomes na variavel (allNames)")
			}
		}
	}()

	for i := 1; i <= totalPaginas; i++ {
		chSemaphore <- struct{}{}
		go func(pagina int) {
			defer func() { <-chSemaphore }()
			nomes, err := pages.GetNomes(strconv.Itoa(pagina))
			if err != nil {
				chNames <- nil
				log.Fatal(err)
				return
			}
			chNames <- nomes
		}(i)
	}
	for i := 0; i < Concurrency; i++ {
		chSemaphore <- struct{}{}
	}
	close(chNames)
	return allNames, nil
}
