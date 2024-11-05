package command

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/MechCarrot/audiometa/internal/interfaces"
)

type GetCommand struct {
	fs     *flag.FlagSet
	client interfaces.Client
	id     string
}

func NewGetCommand(client interfaces.Client) *GetCommand {
	gC := &GetCommand{
		fs:     flag.NewFlagSet("get", flag.ContinueOnError),
		client: client,
	}

	gC.fs.StringVar(&gC.id, "id", "", "id of audio file requested")
	return gC
}

func (getComm *GetCommand) Name() string {
	return getComm.fs.Name()
}

func (getComm *GetCommand) ParseFlag(args []string) error {
	if len(args) == 0 {
		fmt.Println(`usage: ./audiometa-cli get -id <id>
		`)
		return fmt.Errorf("missing flags")
	}
	return getComm.fs.Parse(args)
}

func (getComm *GetCommand) Run() error {
	//Http
	if getComm.id == "" {
		return fmt.Errorf("missing id: see audiometa --help")
	}
	params := "id=" + url.QueryEscape(getComm.id)
	path := fmt.Sprintf("http://localhost:8080/get?%s", params)
	payload := &bytes.Buffer{}
	client := getComm.client
	log.Println("In getcommand")
	req, err := http.NewRequest(http.MethodGet, path, payload)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error reading response, %s\n", err.Error())
		return err
	}
	log.Println("Waiting response")
	log.Println(string(b))
	log.Println(b)
	fmt.Println(string(b))
	return nil
}
