package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type result struct {
	status int
	err    error
}

func worker(url string, requests int, results chan<- result, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < requests; i++ {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			results <- result{status: 0, err: err}
			continue
		}
		results <- result{status: resp.StatusCode, err: nil}
		resp.Body.Close()
		elapsed := time.Since(start)
		log.Printf("Request %d took %s with status %d", i, elapsed, resp.StatusCode)
	}
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	totalRequests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 0, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" || *totalRequests == 0 || *concurrency == 0 {
		flag.Usage()
		return
	}

	results := make(chan result, *totalRequests)
	var wg sync.WaitGroup

	start := time.Now()

	requestsPerWorker := *totalRequests / *concurrency
	extraRequests := *totalRequests % *concurrency

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		requests := requestsPerWorker
		if i < extraRequests {
			requests++
		}
		go worker(*url, requests, results, &wg)
	}

	wg.Wait()
	close(results)

	elapsed := time.Since(start)

	statusCount := make(map[int]int)
	totalCount := 0
	for res := range results {
		if res.err != nil {
			fmt.Println("Error:", res.err)
			continue
		}
		statusCount[res.status]++
		totalCount++
	}

	fmt.Println("Relatório de Teste de Carga")
	fmt.Printf("Tempo total gasto: %s\n", elapsed)
	fmt.Printf("Total de requests realizados: %d\n", totalCount)
	fmt.Println("Distribuição de códigos de status HTTP:")
	for status, count := range statusCount {
		fmt.Printf("Status %d: %d\n", status, count)
	}
}
