package esyaml

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func getNodeValue(node *yaml.Node, path []string) (interface{}, error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("empty path")
	}

	if node.Kind != yaml.DocumentNode {
		return nil, fmt.Errorf("expected document node")
	}

	return getMappingValue(node.Content[0], path)
}

func getMappingValue(node *yaml.Node, path []string) (interface{}, error) {
	if node.Kind != yaml.MappingNode {
		return nil, fmt.Errorf("expected mapping node")
	}

	for i := 0; i < len(node.Content); i += 2 {
		key := node.Content[i]
		value := node.Content[i+1]

		if key.Value == path[0] {
			if len(path) == 1 {
				// We've reached the target field, return its value
				return getNodeScalarValue(value)
			}
			// Continue traversing the path
			return getMappingValue(value, path[1:])
		}
	}

	return nil, fmt.Errorf("path not found: %s", strings.Join(path, "."))
}

func getNodeScalarValue(node *yaml.Node) (interface{}, error) {
	switch node.Kind {
	case yaml.ScalarNode:
		return node.Value, nil
	case yaml.SequenceNode:
		var result []interface{}
		for _, item := range node.Content {
			value, err := getNodeScalarValue(item)
			if err != nil {
				return nil, err
			}
			result = append(result, value)
		}
		return result, nil
	case yaml.MappingNode:
		result := make(map[string]interface{})
		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i].Value
			value, err := getNodeScalarValue(node.Content[i+1])
			if err != nil {
				return nil, err
			}
			result[key] = value
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected node kind: %v", node.Kind)
	}
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

func deleteNode(node *yaml.Node, path []string) error {
	if len(path) == 0 {
		return fmt.Errorf("empty path")
	}

	if node.Kind != yaml.DocumentNode {
		return fmt.Errorf("expected document node")
	}

	return deleteMapping(node.Content[0], path)
}

func deleteMapping(node *yaml.Node, path []string) error {
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("expected mapping node")
	}

	for i := 0; i < len(node.Content); i += 2 {
		key := node.Content[i]
		value := node.Content[i+1]

		if key.Value == path[0] {
			if len(path) == 1 {
				// We've reached the target field, remove it
				node.Content = append(node.Content[:i], node.Content[i+2:]...)
				return nil
			}
			// Continue traversing the path
			return deleteMapping(value, path[1:])
		}
	}

	return fmt.Errorf("path not found: %s", strings.Join(path, "."))
}

func replaceKey(node *yaml.Node, path []string, newKey string) error {
	if len(path) == 0 {
		return fmt.Errorf("empty path")
	}

	if node.Kind != yaml.DocumentNode {
		return fmt.Errorf("expected document node")
	}

	return replaceKeyInMapping(node.Content[0], path, newKey)
}

func replaceKeyInMapping(node *yaml.Node, path []string, newKey string) error {
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("expected mapping node")
	}

	for i := 0; i < len(node.Content); i += 2 {
		key := node.Content[i]
		value := node.Content[i+1]

		if key.Value == path[0] {
			if len(path) == 1 {
				// We've reached the target key, replace it
				key.Value = newKey
				return nil
			}
			// Continue traversing the path
			return replaceKeyInMapping(value, path[1:], newKey)
		}
	}

	return fmt.Errorf("path not found: %s", strings.Join(path, "."))
}
