package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
)

type Config struct {
	Commands []string
}

func main() {
	var commandFileName string
	flag.StringVar(&commandFileName, "config", "dev.json", "Default value: dev.json")
	flag.Parse()

	config := getConfig(commandFileName)
	runCommands(config.Commands)
}

func getConfig(configFileName string) Config {
	_, error := os.Stat(configFileName)

	config := Config{}

	if errors.Is(error, os.ErrNotExist) {
		fmt.Println("Config file does not exist:", configFileName)
		os.Exit(1)
	}

	file, _ := os.ReadFile(configFileName)
	_ = json.Unmarshal(file, &config)
	return config
}

func runCommands(commands []string) {
	var wg sync.WaitGroup
	wg.Add(1)
	listenForExitSignal()
	for i := 0; i < len(commands); i++ {
		wg.Add(1)
		go runCommand(commands[i])
	}
	wg.Wait()
}

func runCommand(command string) {
	parts := strings.Split(command, " ")
	name := parts[0]
	arg := parts[1:]
	fmt.Println("⚙", command)
	output, err := exec.Command(name, arg...).CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))
}

func listenForExitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("Goodbye")
			os.Exit(0)
		}
	}()
}
