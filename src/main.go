package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Config struct {
	Commands []string
}

func main() {
	var commandFileName string
	flag.StringVar(&commandFileName, "config", "pc.json", "Default value: pc.json")
	flag.Parse()

	config := getConfig(commandFileName)
	runCommands(config.Commands)
}

func getConfig(configFileName string) Config {
	_, err := os.Stat(configFileName)

	config := Config{}

	if errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Config file does not exist: %s", configFileName)
	}

	file, _ := os.ReadFile(configFileName)
	_ = json.Unmarshal(file, &config)
	return config
}

func runCommands(commands []string) {
	var wg sync.WaitGroup
	for i := 0; i < len(commands); i++ {
		wg.Add(1)
		go func(command string) {
			defer wg.Done()
			runCommand(command)
		}(commands[i])
	}
	wg.Wait()
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func runCommand(command string) {
	parts := strings.Split(command, " ")
	name := parts[0]

	if !commandExists(name) {
		log.Printf("Could not run: %s : %s\n", color.CyanString(command), color.RedString("Command does not exist."))
		return
	}

	arg := parts[1:]
	fmt.Println("⚙", command)

	cmd := exec.Command(name, arg...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	runErr := cmd.Run()
	if runErr != nil {
		log.Printf("Could not run: %s : %s\n", color.CyanString(command), color.RedString(runErr.Error()))
		return
	}
}
