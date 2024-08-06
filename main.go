package main

import (
	"fmt"
	esyaml "main/yaml"
)

func main() {
	yamlStr := `
spec:
  name: oldName
  value: 42
`

	name, err := esyaml.GetYAMLValue(yamlStr, "spec.name")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(name)

	updatedYAML, err := esyaml.ReplaceYAMLValue(yamlStr, "spec.name", "newName")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(updatedYAML)

}
