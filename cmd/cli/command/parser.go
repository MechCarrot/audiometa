package command

import (
	"fmt"
	"log"

	"github.com/MechCarrot/audiometa/internal/interfaces"
)

type Parser struct {
	commands []interfaces.Command
}

func (p *Parser) GetCommandsLog() {
	log.Println("Parser GetCommandsLog is working")
	for ind, val := range p.commands {
		log.Println(ind, val)
	}
}
func NewParser(commands []interfaces.Command) *Parser {
	return &Parser{commands: commands}
}

func (p *Parser) Parse(args []string) error {
	log.Println("Parsing arguments...")
	p.GetCommandsLog()
	if len(args) < 1 {
		log.Fatal("You need at least 2 arguments")
		return fmt.Errorf("%s", helpMessage)
	}
	subcommand := args[0]
	log.Println("Current full args:", args)
	log.Println("Current Parser commands:", p.commands)
	for _, cmd := range p.commands {
		log.Println("Current commands input", cmd.Name())
	}
	for _, cmd := range p.commands {
		log.Printf("Bruting cmds, current cmd: %T\n", cmd)
		if cmd.Name() == subcommand {
			cmd.ParseFlag(args[1:])
			log.Println("Command", cmd, "is running!")
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand %s", subcommand)
}
