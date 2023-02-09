package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ttacon/chalk"
)

type AuthenticateData struct {
	User           string `json:"user"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	PrivateKeyPath string `json:"private_key_path"`
	Password       string `json:"password"`
	AuthType       string `json:"auth_type"`
}

type Step struct {
	Description string   `json:"description"`
	Commands    []string `json:"commands"`
}

type Workflow struct {
	ServerInfo AuthenticateData `json:"server_info"`
	Steps      []Step           `json:"steps"`
}

func checkFileExists(path string) {
	filepath := ".workflow.kloner.json"

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		log.Fatal(chalk.Red, "`.workflow.kloner.json` file does not exist")
	}
}

func ParseWorkflowFile(path string) *Workflow {
	checkFileExists(path)

	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(chalk.Red, "Failed to load file %v", err, chalk.Cyan, "\n File name %s", path)
	}

	var workflow Workflow

	err = json.Unmarshal(file, &workflow)

	if err != nil {
		log.Fatal(chalk.Red, "Something happened parsing the .workflow.kloner.json file %s", err)
	}

	return &workflow
}
