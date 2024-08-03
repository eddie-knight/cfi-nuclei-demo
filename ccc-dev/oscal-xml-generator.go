// package main

// import (
// 	"encoding/xml"
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
// 	XMLName     xml.Name    `xml:"component-definition"`
// 	Xmlns       string      `xml:"xmlns,attr"`
// 	Title       string      `xml:"title"`
// 	Description string      `xml:"description"`
// 	Version     string      `xml:"version"`
// 	Components  []Component `xml:"components>component"`
// }

// type Component struct {
// 	UUID                   string                  `xml:"uuid,attr"`
// 	Title                  string                  `xml:"title"`
// 	Description            string                  `xml:"description"`
// 	ControlImplementations []ControlImplementation `xml:"control-implementations>control-implementation"`
// }

// type ControlImplementation struct {
// 	UUID                    string                   `xml:"uuid,attr"`
// 	Source                  string                   `xml:"source"`
// 	Description             string                   `xml:"description"`
// 	ImplementedRequirements []ImplementedRequirement `xml:"implemented-requirements>implemented-requirement"`
// }

// type ImplementedRequirement struct {
// 	UUID           string          `xml:"uuid,attr"`
// 	ControlID      string          `xml:"control-id"`
// 	Prose          string          `xml:"prose"`
// 	TestProcedures []TestProcedure `xml:"test-procedures>test-procedure"`
// }

// type TestProcedure struct {
// 	UUID  string `xml:"uuid,attr"`
// 	Prose string `xml:"prose"`
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
// 	oscal := OSCALComponentDefinition{
// 		Xmlns:       "http://csrc.nist.gov/ns/oscal/1.0",
// 		Title:       data.Title,
// 		Description: data.Title,
// 		Version:     data.Version,
// 	}

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

// 	// Marshal to XML
// 	output, err := xml.MarshalIndent(oscal, "", "  ")
// 	if err != nil {
// 		log.Fatalf("error: %v", err)
// 	}

// 	// Write to XML file
// 	err = ioutil.WriteFile("CCC.OS.xml", output, 0644)
// 	if err != nil {
// 		log.Fatalf("error: %v", err)
// 	}

// 	fmt.Println("OSCAL component definition XML generated successfully!")
// }
