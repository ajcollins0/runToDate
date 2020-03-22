package main

import (
	"fmt"

	"./arg"
	"./stravaApi"
)

func main() {
	params := arg.ParseArgs()
	fmt.Printf("%0.2f \n", stravaApi.DistanceRun(*params))
}
