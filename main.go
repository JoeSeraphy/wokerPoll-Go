package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	id    int
	value int
}
type Result struct {
	jobID int
	total int
}

func wokers(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {

	defer wg.Done()
	for j := range jobs {
		fmt.Println("Woker iniciando job", j.id, j.value)
		time.Sleep(time.Millisecond * 500)
		results <- Result{jobID: j.id, total: j.value * 2}
	}

}
func main() {
	const numJobs = 10
	const numWokers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numWokers)
	var wg sync.WaitGroup

	for w := 1; w <= numWokers; w++ {
		wg.Add(1)
		go wokers(w, jobs, results, &wg)
	}
	for j := 1; j <= numJobs; j++ {
		wg.Add(1)
		jobs <- Job{id: j, value: j * 10}
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("\n --- Resultado Finais ---")
	for res := range results {
		fmt.Printf("Job %d finalizado. Resultado: %d\n", res.jobID, res.total)
		wg.Done()
	}

}
