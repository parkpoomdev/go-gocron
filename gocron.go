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
