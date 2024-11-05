package command

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

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
		fmt.Println(`usage: ./audiometa upload -id <id>`)
		return fmt.Errorf("missing flags")
	}

	return cmd.fs.Parse(args)
}

func (cmd *UploadCommand) Run() error {
	if cmd.filename == "" {
		return fmt.Errorf("filename shouldn't be empty")
	}
	log.Println("Uploading file to the server:", cmd.filename)
	payload := &bytes.Buffer{}

	file, err := os.Open(cmd.filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	multipartWriter := multipart.NewWriter(payload)
	partWriter, err := multipartWriter.CreateFormFile("file", filepath.Base(cmd.filename))

	if err != nil {
		return err
	}

	_, err = io.Copy(partWriter, file)
	if err != nil {
		return err
	}

	err = multipartWriter.Close()
	if err != nil {
		return err
	}

	client := cmd.client
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/upload", payload)
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", multipartWriter.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println("audio file id - ", string(body))
	return nil
}
