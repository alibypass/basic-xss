package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {

	target := flag.String("d", "", "string of target")
	flag.Parse()

	// Define the Bash command you want to run
	cmd := exec.Command("bash", "-c", fmt.Sprintf("ggospider  %s >> gospiderresult.txt && cat gospiderresult.txt | httpx >> final.txt", *target))

	// Run the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Output: %s\n", output)

}

// func bash() {
// 	// Define the Bash command you want to run
// 	cmd := exec.Command("bash", "-c", "ggospider ", target)

// 	// Run the command and capture its output
// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}

// 	// Print the output
// 	fmt.Printf("Output: %s\n", output)
// }
