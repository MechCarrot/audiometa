package command

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/MechCarrot/audiometa/internal/interfaces"
)

type ListCommand struct {
	fs     *flag.FlagSet
	client interfaces.Client
}

func NewListCommand(client interfaces.Client) *ListCommand {
	lst := &ListCommand{
		fs:     flag.NewFlagSet("list", flag.ContinueOnError),
		client: client,
	}
	return lst
}

func (lst *ListCommand) Name() string {
	return lst.fs.Name()
}

func (lst *ListCommand) ParseFlag(args []string) error {
	return lst.fs.Parse(args)
}

func (lst *ListCommand) Run() error {
	path := "http://localhost:8080/list"
	payload := &bytes.Buffer{}
	client := lst.client

	req, err := http.NewRequest(http.MethodGet, path, payload)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(respBody))
	return nil
}
