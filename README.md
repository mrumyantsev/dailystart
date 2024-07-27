# Dailyt

**Dailyt** is a Go library, that provides a functions for starting a daily task (or routine) on the particular time of the day.

## Usage

``` Golang
package main

import (
	"fmt"
	"time"

	"github.com/mrumyantsev/dailyt"
)

func main() {
	// A daily time for starting the task.
	dailyTaskTime := "14:36:30"

	// Visual time printer.
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for t := range ticker.C {
			fmt.Println(t.Format(time.TimeOnly))
		}
	}()

	// An infinite loop, which provides daily task launch.
	for {
		dailyt.SleepUntil(dailyTaskTime)

		task()
	}
}

func task() {
	fmt.Println("Task has been launched!")
}
```
