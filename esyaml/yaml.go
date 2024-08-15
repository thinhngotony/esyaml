package esyaml

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func SetYAMLValue(yamlStr, path string, newValue interface{}) (string, error) {
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

func GetYAMLValue(yamlStr, path string) (interface{}, error) {
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yamlStr), &node)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	pathParts := strings.Split(path, ".")
	value, err := getNodeValue(&node, pathParts)
	if err != nil {
		return nil, fmt.Errorf("failed to get node value: %w", err)
	}

	return value, nil
}

func DeleteYAMLField(yamlStr, path string) (string, error) {
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yamlStr), &node)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	pathParts := strings.Split(path, ".")
	err = deleteNode(&node, pathParts)
	if err != nil {
		return "", fmt.Errorf("failed to delete node: %w", err)
	}

	updatedYAML, err := yaml.Marshal(&node)
	if err != nil {
		return "", fmt.Errorf("failed to marshal updated YAML: %w", err)
	}

	return string(updatedYAML), nil
}

func ReplaceYAMLKey(yamlStr, oldKeyPath, newKey string) (string, error) {
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yamlStr), &node)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	pathParts := strings.Split(oldKeyPath, ".")
	err = replaceKey(&node, pathParts, newKey)
	if err != nil {
		return "", fmt.Errorf("failed to replace key: %w", err)
	}

	updatedYAML, err := yaml.Marshal(&node)
	if err != nil {
		return "", fmt.Errorf("failed to marshal updated YAML: %w", err)
	}

	return string(updatedYAML), nil
}

func InsertYAML(yamlStr, path string, newValue interface{}) (string, error) {
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yamlStr), &node)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	pathParts := strings.Split(path, ".")
	err = insertNode(&node, pathParts, newValue)
	if err != nil {
		return "", fmt.Errorf("failed to insert node: %w", err)
	}

	updatedYAML, err := yaml.Marshal(&node)
	if err != nil {
		return "", fmt.Errorf("failed to marshal updated YAML: %w", err)
	}

	return string(updatedYAML), nil
}

func MustSetYAMLValue(yamlStr, fieldName string, newValue interface{}) (string, error) {
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yamlStr), &node)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	updateAllOccurrences(&node, fieldName, newValue)

	updatedYAML, err := yaml.Marshal(&node)
	if err != nil {
		return "", fmt.Errorf("failed to marshal updated YAML: %w", err)
	}

	return string(updatedYAML), nil
}

func MustPrependYAMLValue(yamlStr, fieldName string, prependValue string) (string, error) {
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yamlStr), &node)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	prependAllOccurrences(&node, fieldName, prependValue)

	updatedYAML, err := yaml.Marshal(&node)
	if err != nil {
		return "", fmt.Errorf("failed to marshal updated YAML: %w", err)
	}

	return string(updatedYAML), nil
}

// func MustReplaceTemplateYAMLValue(yamlStr, placeholder, replaceValue string) (string, error) {
// 	var node yaml.Node
// 	err := yaml.Unmarshal([]byte(yamlStr), &node)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to unmarshal YAML: %w", err)
// 	}

// 	replaceAllOccurrences(&node, placeholder, replaceValue)

// 	updatedYAML, err := yaml.Marshal(&node)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to marshal updated YAML: %w", err)
// 	}

// 	return string(updatedYAML), nil
// }

// func replaceAllOccurrences(node *yaml.Node, placeholder, replaceValue string) {
// 	switch node.Kind {
// 	case yaml.DocumentNode, yaml.SequenceNode:
// 		for i := range node.Content {
// 			replaceAllOccurrences(node.Content[i], placeholder, replaceValue)
// 		}
// 	case yaml.MappingNode:
// 		for i := 0; i < len(node.Content); i += 2 {
// 			value := node.Content[i+1]
// 			replaceAllOccurrences(value, placeholder, replaceValue)
// 		}
// 	case yaml.ScalarNode:
// 		node.Value = strings.ReplaceAll(node.Value, "{{"+placeholder+"}}", replaceValue)
// 	}
// }

func AddTmlValue(yamlStr, placeholder, replaceValue string) (string, error) {
	// Convert placeholder to lowercase for case-insensitive matching
	placeholderLower := strings.ToLower(placeholder)

	// Split the YAML string into lines
	lines := strings.Split(yamlStr, "\n")

	for i, line := range lines {
		// Find the start and end positions of the placeholder
		start := strings.Index(line, "{{")
		end := strings.Index(line, "}}")

		// If both {{ and }} are found in the line
		if start != -1 && end != -1 && end > start {
			// Extract the content between {{ and }}
			fullPlaceholder := line[start : end+2]
			content := line[start+2 : end]

			// Remove all whitespace and convert to lowercase
			contentCleaned := strings.ToLower(strings.ReplaceAll(content, " ", ""))

			// If the content matches the placeholder (case-insensitive)
			if contentCleaned == placeholderLower {
				// Replace the entire {{placeholder}} (including possible spaces) with the new value
				lines[i] = strings.Replace(line, fullPlaceholder, replaceValue, 1)
			}
		}
	}

	// Join the lines back into a single string
	return strings.Join(lines, "\n"), nil
}
