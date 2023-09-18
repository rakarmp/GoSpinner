package main

import (
	"fmt"
	"time"

	gospinner "github.com/rakarmp/GoSpinner/GoSpinner"
)

func main() {
	spinner := gospinner.NewLoadingSpinner()

	// Default spinner
	spinner.Start("default", 5*time.Second)
	fmt.Println("Done!")

	// Arrow spinner
	spinner.Start("arrow", 3*time.Second)
	fmt.Println("Done!")

	// Circle spinner
	spinner.Start("circle", 7*time.Second)
	fmt.Println("Done!")

	// Star spinner
	spinner.Start("star", 4*time.Second)
	fmt.Println("Done!")
}
