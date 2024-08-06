package esyaml

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func ReplaceYAMLValue(yamlStr, path string, newValue interface{}) (string, error) {
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
