# go-gocron

This repository is a beginner's guide to using `gocron` in Go language. `gocron` is a powerful cron library for Go that allows you to schedule and run periodic tasks.

## Installation

To use `gocron`, you need to have Go installed on your system. If you haven't installed Go, you can download it from the official Go website: https://golang.org/

Once Go is installed, you can install `gocron` by running the following command:

```shell
go get github.com/go-co-op/gocron
```

Getting Started

To start using gocron, you need to import the gocron package in your Go program:
```shell
import "github.com/go-co-op/gocron"
```

Scheduling Tasks

## Example

```shell
package main

import (
 "fmt"
 "time"

 "github.com/go-co-op/gocron"
)

// Peint Message function for example Hello World!
func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func runCronJobs() {
	// Define Instance
	s := gocron.NewScheduler(time.UTC)

	// Print messange every single second
	s.Every(1).Seconds().Do(func() {

		hello("John Doe")

	})

	// Exeure the instance
	s.StartBlocking()
}

// Start Main Function
func main() {
	runCronJobs()
}
```

## Running Command
```shell
go run server.go
```

## Result
```shell
Hi, John Doe
Hi, John Doe
Hi, John Doe
Hi, John Doe
Hi, John Doe
Hi, John Doe
Hi, John Doe
Hi, John Doe
Hi, John Doe
Hi, John Doe
```

