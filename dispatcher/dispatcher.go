// Package dispatcher implements spawning and executing of workers
package dispatcher

import (
	"fmt"
	"strings"

	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/workers"
)

func Run(startingWorkerCount int, maxWorkerCount int, totalJobCount int) {
	testCount := 1
	for i := startingWorkerCount; i <= maxWorkerCount; i++ {
		currentWorkers := i

		if config.Debug {
			text := fmt.Sprintf("|  Spawning workers for test %v of %v  |", testCount, (maxWorkerCount - startingWorkerCount + 1))
			divider := strings.Repeat("-", len(text))
			fmt.Println()
			config.ConsoleCyan.Println(divider)
			config.ConsoleCyan.Printf("|  Spawning workers for test %v of %v  |\n", testCount, (maxWorkerCount - startingWorkerCount + 1))
			config.ConsoleCyan.Println(divider)
			fmt.Println()
		}

		elapsed := workers.Workers(currentWorkers, totalJobCount, "emptySleepJob")
		workers.WorkerStats = append(workers.WorkerStats, workers.WorkerStat{Workers: currentWorkers, ExecutionTime: elapsed})

		testCount++
	}
}
