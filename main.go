package main

import (
	"fmt"

	"main/esyaml" // Replace with the actual import path
)

const test = `apiVersion: apps/v1 
kind: Deployment
metadata:
  name: ubuntu-deployment
spec:
  selector:
    matchLabels:
      app: ubuntu
  replicas: 10 # amount of pods must be > 1
  template:
    metadata:
      labels:
        app: ubuntu
    spec:
      containers:
      - name: ubuntu
        image: ubuntu
        command:
        - sleep
        - "infinity"
        volumeMounts:
        - mountPath: /app/folder
          name: volume
      volumes:
      - name: volume
        persistentVolumeClaim:
          claimName: new-value
      - name: volume2
        persistentVolumeClaim:
          claimName: new-value		  
`

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
	updatedYAML, err := esyaml.SetYAMLValue(yamlStr, "spec.name", "newName")
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
	replacedValueYAML, err := esyaml.SetYAMLValue(yamlStr, "spec.value", "new-value")
	if err != nil {
		fmt.Println("Error replacing value:", err)
		return
	}
	fmt.Println("Replaced YAML:")
	fmt.Println(replacedValueYAML)

	// --- Test insert value ---
	fmt.Println("\n--- Test replace value ---")
	insertedValueYAML, err := esyaml.InsertYAML(yamlStr, "spec.replicas", 2)
	if err != nil {
		fmt.Println("Error replacing value:", err)
		return
	}
	fmt.Println("Inserted YAML:")
	fmt.Println(insertedValueYAML)

	// --- Test replace value ---
	fmt.Println("\n--- Test must set value ---")
	mustSetValueYAML, err := esyaml.MustSetYAMLValue(test, "claimName", "new-value")
	if err != nil {
		fmt.Println("Error must set:", err)
		return
	}
	fmt.Println("Must set YAML:")
	fmt.Println(mustSetValueYAML)

	// --- Test must prepend value ---
	fmt.Println("\n--- Test must prepend value ---")
	mustPrependValueYAML, err := esyaml.MustPrependYAMLValue(test, "claimName", "new-value-")
	if err != nil {
		fmt.Println("Error must prepend set:", err)
		return
	}
	fmt.Println("Must prepend YAML:")
	fmt.Println(mustPrependValueYAML)

}
