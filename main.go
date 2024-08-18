// Exercise 4.11
// Build a tool that lets users create, read, update, and close GitHub issues from the command line, invoking their preferred text editor when substantial text input is required.

package main

import (
	"log"
	"os"

	"exercise-4.11/github"
	"github.com/joho/godotenv"
)

type config struct {
	client github.Client
	editor string
}

type cliCommand struct {
	command     string
	description string
	callback    func(*config, []string) error
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("you need to include a command name and appropriate parameters")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	client := github.NewClient(os.Getenv("GITHUB_TOKEN"), os.Getenv("GITHUB_USER"))
	config := &config{
		client: client,
		editor: os.Getenv("EDITOR"),
	}

	commands := genCommands()
	command, ok := commands[os.Args[1]]
	if !ok {
		log.Fatal("Unknown command. Use 'help' to see available commands.")
	}

	params := []string{}
	if command.command != "help" && len(os.Args) > 2 {
		params = os.Args[2:]
	}

	err = command.callback(config, params)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func genCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			command:     "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"read": {
			command:     "read",
			description: "Returns info for an issue. Must pass a repo name and an issue id.",
			callback:    commandRead,
		},
		"create": {
			command:     "create",
			description: "Creates a new issue. Must provide a title and body. Can also optionally provide a list of labels.",
			callback:    commandCreate,
		},
		"update": {
			command:     "update",
			description: "Edit an issue. Must pass a repo name, issue id, and the field to edit which can be 'title', 'body' or 'labels'.",
			callback:    commandUpdate,
		},
		"close": {
			command:     "close",
			description: "Close an issue. Must pass a repo name and issue id.",
			callback:    commandClose,
		},
	}
}
