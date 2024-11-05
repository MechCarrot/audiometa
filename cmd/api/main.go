package main

import (
	"flag"
	"fmt"

	metadataservice "github.com/MechCarrot/audiometa/services/metadata"
)

func main() {
	//Flag variables
	var port int

	//Flag definition
	flag.IntVar(&port, "p", 8080, "Port for metadata service")

	flag.Parse()

	fmt.Printf("Starting API at localhost:%d\n", port)
	metadataservice.Run(port)
	fmt.Printf("API is working")
}
