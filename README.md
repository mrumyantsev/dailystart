# Dailystart

**Dailystart** is a Go library, that provides a functions for starting a daily task (or routine) on the particular time of the day.

## Usage

``` Golang
package main

import (
	"fmt"
	"time"

	"github.com/mrumyantsev/dailystart"
)

func main() {
	// A day time for starting a daily task.
	dailyTaskTime := "14:36:30"

	// Visual time printer.
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for t := range ticker.C {
			fmt.Println(t.Format(time.TimeOnly))
		}
	}()

	// An infinite loop, which provides daily task starting.
	for {
		dailystart.SleepUntil(dailyTaskTime)

		task()
	}
}

func task() {
	fmt.Println("Task started")
}

```
