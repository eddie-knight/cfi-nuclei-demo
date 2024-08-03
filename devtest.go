package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Control struct {
	ID                  string
	Title               string
	Feature             string
	NISTCSF             string
	MITREATTACK         string
	Objective           string
	ControlMappings     []string
	TestingRequirements []string
}

func main() {
	inputFile := "CCC.OS.md"
	controls, err := parseMarkdownFile(inputFile)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	for _, control := range controls {
		err := writeYAMLFile(control)
		if err != nil {
			fmt.Println("Error writing YAML file:", err)
		}
	}
}

func parseMarkdownFile(filename string) ([]Control, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var controls []Control
	var currentControl *Control
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "## CCC") {
			if currentControl != nil {
				controls = append(controls, *currentControl)
			}
			currentControl = &Control{
				ID:    strings.Fields(line)[1],
				Title: strings.Join(strings.Fields(line)[2:], " "),
			}
		} else if strings.HasPrefix(line, "- Corresponding Feature:") {
			currentControl.Feature = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.HasPrefix(line, "- NIST CSF:") {
			currentControl.NISTCSF = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.HasPrefix(line, "- MITRE ATT&CK TTP:") {
			currentControl.MITREATTACK = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.HasPrefix(line, "### Objective") {
			scanner.Scan()
			currentControl.Objective = scanner.Text()
		} else if strings.HasPrefix(line, "- CCM:") || strings.HasPrefix(line, "- ISO/IEC 27001:2013") || strings.HasPrefix(line, "- NIST SP 800-53") {
			currentControl.ControlMappings = append(currentControl.ControlMappings, strings.TrimSpace(line))
		} else if strings.HasPrefix(line, "1. **CCC.OS") {
			currentControl.TestingRequirements = append(currentControl.TestingRequirements, strings.TrimSpace(line))
		}
	}

	if currentControl != nil {
		controls = append(controls, *currentControl)
	}

	return controls, scanner.Err()
}

func writeYAMLFile(control Control) error {
	filename := fmt.Sprintf("ccc-os-%s.yaml", strings.ToLower(control.ID))
	content := fmt.Sprintf(
		"ID: %s\nTitle: %s\nFeature: %s\nNIST CSF: %s\nMITRE ATT&CK: %s\nObjective: %s\nControl Mappings:\n%s\nTesting Requirements:\n%s\n",
		control.ID, control.Title, control.Feature, control.NISTCSF, control.MITREATTACK, control.Objective,
		strings.Join(control.ControlMappings, "\n"), strings.Join(control.TestingRequirements, "\n"),
	)
	return ioutil.WriteFile(filename, []byte(content), 0644)
}
