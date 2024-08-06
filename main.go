package main

import (
	"fmt"

	esyaml "main/yaml" // Replace with the actual import path
)

func main() {
	// Define the original YAML string.
	yamlStr := `
spec:
  name: oldName
  value: 42
`
	fmt.Println("Original YAML:")
	fmt.Println(yamlStr)

	// --- Test get ---
	fmt.Println("\n--- Test get ---")
	name, err := esyaml.GetYAMLValue(yamlStr, "spec.name")
	if err != nil {
		fmt.Println("Error getting value:", err)
		return
	}
	fmt.Println("Value at spec.name:", name)

	// --- Test update ---
	fmt.Println("\n--- Test update ---")
	updatedYAML, err := esyaml.ReplaceYAMLValue(yamlStr, "spec.name", "newName")
	if err != nil {
		fmt.Println("Error updating value:", err)
		return
	}
	fmt.Println("Updated YAML:")
	fmt.Println(updatedYAML)

	// --- Test delete ---
	fmt.Println("\n--- Test delete ---")
	deletedYAML, err := esyaml.DeleteYAMLField(yamlStr, "spec.name")
	if err != nil {
		fmt.Println("Error deleting field:", err)
		return
	}
	fmt.Println("Deleted YAML:")
	fmt.Println(deletedYAML)

	// --- Test replace key ---
	fmt.Println("\n--- Test replace key ---")
	replacedYAML, err := esyaml.ReplaceYAMLKey(yamlStr, "spec.name", "title")
	if err != nil {
		fmt.Println("Error replacing key:", err)
		return
	}
	fmt.Println("Replaced YAML:")
	fmt.Println(replacedYAML)

	// --- Test replace value ---
	fmt.Println("\n--- Test replace value ---")
	replacedValueYAML, err := esyaml.ReplaceYAMLValue(yamlStr, "spec.value", "new-value")
	if err != nil {
		fmt.Println("Error replacing value:", err)
		return
	}
	fmt.Println("Replaced YAML:")
	fmt.Println(replacedValueYAML)
}
