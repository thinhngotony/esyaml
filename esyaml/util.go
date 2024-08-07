package esyaml

import (
	"fmt"
	"math"
	"reflect"
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
				return updateNodeValue(value, newValue)
			}
			// Continue traversing the path
			return updateMapping(value, path[1:], newValue)
		}
	}

	return fmt.Errorf("path not found: %s", strings.Join(path, "."))
}

func updateNodeValue(node *yaml.Node, newValue interface{}) error {
	v := reflect.ValueOf(newValue)

	switch v.Kind() {
	case reflect.String:
		node.Kind = yaml.ScalarNode
		node.Tag = "!!str"
		node.Value = v.String()

	case reflect.Bool:
		node.Kind = yaml.ScalarNode
		node.Tag = "!!bool"
		node.Value = fmt.Sprintf("%t", v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		node.Kind = yaml.ScalarNode
		node.Tag = "!!int"
		node.Value = fmt.Sprintf("%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		node.Kind = yaml.ScalarNode
		node.Tag = "!!int"
		node.Value = fmt.Sprintf("%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		node.Kind = yaml.ScalarNode
		node.Tag = "!!float"
		f := v.Float()
		if math.IsNaN(f) {
			node.Value = ".nan"
		} else if math.IsInf(f, 1) {
			node.Value = ".inf"
		} else if math.IsInf(f, -1) {
			node.Value = "-.inf"
		} else {
			node.Value = fmt.Sprintf("%g", f)
		}

	case reflect.Slice, reflect.Array:
		node.Kind = yaml.SequenceNode
		node.Tag = "!!seq"
		node.Content = make([]*yaml.Node, v.Len())
		for i := 0; i < v.Len(); i++ {
			itemNode := &yaml.Node{}
			if err := updateNodeValue(itemNode, v.Index(i).Interface()); err != nil {
				return err
			}
			node.Content[i] = itemNode
		}

	case reflect.Map:
		node.Kind = yaml.MappingNode
		node.Tag = "!!map"
		node.Content = make([]*yaml.Node, 0, v.Len()*2)
		for _, key := range v.MapKeys() {
			keyNode := &yaml.Node{}
			valueNode := &yaml.Node{}
			if err := updateNodeValue(keyNode, key.Interface()); err != nil {
				return err
			}
			if err := updateNodeValue(valueNode, v.MapIndex(key).Interface()); err != nil {
				return err
			}
			node.Content = append(node.Content, keyNode, valueNode)
		}

	case reflect.Ptr:
		if v.IsNil() {
			node.Kind = yaml.ScalarNode
			node.Tag = "!!null"
			node.Value = "null"
		} else {
			return updateNodeValue(node, v.Elem().Interface())
		}

	default:
		fmt.Printf("unsupported type: %T", newValue)
		return fmt.Errorf("unsupported type: %T", newValue)
	}

	return nil
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
