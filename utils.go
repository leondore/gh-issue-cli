package main

import (
	"fmt"
	"os"
	"os/exec"
)

func openEditor(c *config, initialContent string) ([]byte, error) {
	editor := "nano"
	if len(c.editor) > 0 {
		editor = c.editor
	}

	dir := os.TempDir()
	file, err := os.CreateTemp(dir, "tmpIssueContent")
	if err != nil {
		return nil, fmt.Errorf("error while creating tempFile: %s", err.Error())
	}
	defer file.Close()

	err = os.WriteFile(file.Name(), []byte(initialContent), 0666)
	if err != nil {
		return nil, fmt.Errorf("error while creating tempFile: %s", err.Error())
	}

	path, err := exec.LookPath(editor)
	if err != nil {
		return nil, fmt.Errorf("couldn't find path to %s: %s", editor, err.Error())
	}

	cmd := exec.Command(path, file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start editor: %s", err.Error())
	}
	fmt.Printf("Waiting for command to finish.\n")
	err = cmd.Wait()
	if err != nil {
		return nil, fmt.Errorf("command finished with error: %s", err.Error())
	}

	content, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}

	return content, nil
}
