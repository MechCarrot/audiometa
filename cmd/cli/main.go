package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MechCarrot/audiometa/cmd/cli/command"
	"github.com/MechCarrot/audiometa/internal/interfaces"
)

func main() {
	client := &http.Client{}
	cmds := []interfaces.Command{
		command.NewUploadCommand(client),
		command.NewGetCommand(client),
		command.NewListCommand(client),
	}
	log.Print("Current CMD's in main.go: ")
	for _, val := range cmds {
		log.Printf("name: %+v \n", val.Name())
	}
	//log.Printf("Current CMD's in main.go:%+v\n", cmds)
	parser := command.NewParser(cmds)
	if err := parser.Parse(os.Args[1:]); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		os.Exit(1)
	}
}
