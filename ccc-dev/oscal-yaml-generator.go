package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"

// 	"gopkg.in/yaml.v2"
// )

// // YAMLData represents the structure of the input YAML file.
// type YAMLData struct {
// 	ID       string    `yaml:"id"`
// 	Title    string    `yaml:"title"`
// 	Version  string    `yaml:"version"`
// 	Controls []Control `yaml:"controls"`
// }

// // Control represents the structure of each control within the YAML file.
// type Control struct {
// 	ID               string              `yaml:"id"`
// 	FeatureID        string              `yaml:"feature_id"`
// 	Title            string              `yaml:"title"`
// 	Objective        string              `yaml:"objective"`
// 	NISTCSF          string              `yaml:"nist_csf"`
// 	MITREAttack      string              `yaml:"mitre_attack"`
// 	ControlMappings  map[string][]string `yaml:"control_mappings"`
// 	TestRequirements map[string]string   `yaml:"test_requirements"`
// }

// // OSCAL Component Definition structures
// type OSCALComponentDefinition struct {
// 	Metadata struct {
// 		Title       string `yaml:"title"`
// 		Description string `yaml:"description"`
// 		Version     string `yaml:"version"`
// 	} `yaml:"metadata"`
// 	Components []Component `yaml:"components"`
// }

// type Component struct {
// 	UUID                   string                  `yaml:"uuid"`
// 	Title                  string                  `yaml:"title"`
// 	Description            string                  `yaml:"description"`
// 	ControlImplementations []ControlImplementation `yaml:"control-implementations"`
// }

// type ControlImplementation struct {
// 	UUID                    string                   `yaml:"uuid"`
// 	Source                  string                   `yaml:"source"`
// 	Description             string                   `yaml:"description"`
// 	ImplementedRequirements []ImplementedRequirement `yaml:"implemented-requirements"`
// }

// type ImplementedRequirement struct {
// 	UUID           string          `yaml:"uuid"`
// 	ControlID      string          `yaml:"control-id"`
// 	Prose          string          `yaml:"prose"`
// 	TestProcedures []TestProcedure `yaml:"test-procedures"`
// }

// type TestProcedure struct {
// 	UUID  string `yaml:"uuid"`
// 	Prose string `yaml:"prose"`
// }

// func main() {
// 	// Read the YAML file
// 	yamlFile, err := ioutil.ReadFile("yaml/CCC.OS.yaml")
// 	if err != nil {
// 		log.Fatalf("error: %v", err)
// 	}

// 	// Unmarshal the YAML file into a struct
// 	var data YAMLData
// 	err = yaml.Unmarshal(yamlFile, &data)
// 	if err != nil {
// 		log.Fatalf("error: %v", err)
// 	}

// 	// Create OSCAL component definition
// 	oscal := OSCALComponentDefinition{}
// 	oscal.Metadata.Title = data.Title
// 	oscal.Metadata.Description = data.Title
// 	oscal.Metadata.Version = data.Version

// 	for _, control := range data.Controls {
// 		component := Component{
// 			UUID:        control.ID,
// 			Title:       control.Title,
// 			Description: control.Objective,
// 		}

// 		impl := ControlImplementation{
// 			UUID:        control.ID,
// 			Source:      "https://example.com/source", // Update as needed
// 			Description: control.Title,
// 		}

// 		implReq := ImplementedRequirement{
// 			UUID:      control.ID,
// 			ControlID: control.ID,
// 			Prose:     control.Objective,
// 		}

// 		for trID, trDesc := range control.TestRequirements {
// 			testProcedure := TestProcedure{
// 				UUID:  trID,
// 				Prose: trDesc,
// 			}
// 			implReq.TestProcedures = append(implReq.TestProcedures, testProcedure)
// 		}

// 		impl.ImplementedRequirements = append(impl.ImplementedRequirements, implReq)
// 		component.ControlImplementations = append(component.ControlImplementations, impl)
// 		oscal.Components = append(oscal.Components, component)
// 	}

// 	// Marshal to YAML
// 	output, err := yaml.Marshal(&oscal)
// 	if err != nil {
// 		log.Fatalf("error: %v", err)
// 	}

// 	// Write to YAML file
// 	err = ioutil.WriteFile("CCC.OS.oscal.yaml", output, 0644)
// 	if err != nil {
// 		log.Fatalf("error: %v", err)
// 	}

// 	fmt.Println("OSCAL component definition YAML generated successfully!")
// }
