package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// channels := make(chan int)
	// channels <- 15
	// valor <- channels

	start := time.Now()

	apis := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://management.azure.com",
		"https://api.github.com/users/octocat",
	}

	ch := make(chan string)
	for _, api := range apis {
		go checkApi(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	// time.Sleep(5 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Tiempo total de ejecución: %s\n", elapsed)
}

func checkApi(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("Error al acceder a la API %s: %v\n", api, err)
		return
	}

	ch <- fmt.Sprintf("API %s está disponible\n", api)
}
