package command

import (
	"flag"
	"fmt"

	"github.com/MechCarrot/audiometa/internal/interfaces"
)

type UploadCommand struct {
	fs       *flag.FlagSet
	client   interfaces.Client
	filename string
}

func NewUploadCommand(client interfaces.Client) *UploadCommand {
	cmd := &UploadCommand{
		fs:     flag.NewFlagSet("upload", flag.ContinueOnError),
		client: client,
	}
	cmd.fs.StringVar(&cmd.filename, "filename", "", "filename of audio file to upload")
	return cmd
}

func (cmd *UploadCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd *UploadCommand) ParseFlag(args []string) error {
	if len(args) == 0 {
		fmt.Println(`usage: ./audiometa-cli upload -id <id>`)
		return fmt.Errorf("missing flags")
	}

	return cmd.fs.Parse(args)
}

func (cmd *UploadCommand) Run() error {
	/*if cmd.filename == "" {
		return fmt.Errorf("filename shouldn't be empty")
	}

	path := "localhost:8080/upload"
	payload := &bytes.Buffer{}

	file, err := os.Open(cmd.filename)
	if err != nil {
		return nil
	}
	defer file.Close()
	*/

	return nil
}
