package main

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func replaceYAMLValue(yamlStr, path string, newValue interface{}) (string, error) {
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yamlStr), &node)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	pathParts := strings.Split(path, ".")
	err = updateNode(&node, pathParts, newValue)
	if err != nil {
		return "", fmt.Errorf("failed to update node: %w", err)
	}

	updatedYAML, err := yaml.Marshal(&node)
	if err != nil {
		return "", fmt.Errorf("failed to marshal updated YAML: %w", err)
	}

	return string(updatedYAML), nil
}

func updateNode(node *yaml.Node, path []string, newValue interface{}) error {
	if len(path) == 0 {
		return fmt.Errorf("empty path")
	}

	if node.Kind != yaml.DocumentNode {
		return fmt.Errorf("expected document node")
	}

	return updateMapping(node.Content[0], path, newValue)
}

func updateMapping(node *yaml.Node, path []string, newValue interface{}) error {
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("expected mapping node")
	}

	for i := 0; i < len(node.Content); i += 2 {
		key := node.Content[i]
		value := node.Content[i+1]

		if key.Value == path[0] {
			if len(path) == 1 {
				// We've reached the target field, update its value
				value.SetString(fmt.Sprintf("%v", newValue))
				return nil
			}
			// Continue traversing the path
			return updateMapping(value, path[1:], newValue)
		}
	}

	return fmt.Errorf("path not found: %s", strings.Join(path, "."))
}

func main() {
	yamlStr := `
spec:
  name: oldName
  value: 42
`

	updatedYAML, err := replaceYAMLValue(yamlStr, "spec.name", "newName")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(updatedYAML)
}

