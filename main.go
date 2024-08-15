package main

import (
	"fmt"

	"main/esyaml" // Replace with the actual import path
)

const deployment = `apiVersion: apps/v1 
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

const template = `apiVersion: apps/v1
# ở đây phiên bản cũ hơn của kubernetes có dạng extensions/v1beta1
kind: Deployment
# kind là loại Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 2
  # replica ở đây sẽ tạo ra 2 pods luôn luôn chạy, khi một số pods bị down hay chết hay bất kì lý do nào đó sẽ tự động tạo lại số lượng pods bằng 2
  selector:
    matchLabels:
      app: nginx-deployment
  template:
    metadata:
      labels:
        app: nginx-deployment
    spec:
      containers:
      - name: nginx-deployment
      - tony: {{Add-on.Name}}
      # image của container docker
        image: nginx
        ports:
        # port bên trong container
        - containerPort: 8080
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
	mustSetValueYAML, err := esyaml.MustSetYAMLValue(deployment, "claimName", "new-value")
	if err != nil {
		fmt.Println("Error must set:", err)
		return
	}
	fmt.Println("Must set YAML:")
	fmt.Println(mustSetValueYAML)

	// --- Test must prepend value ---
	fmt.Println("\n--- Test must prepend value ---")
	mustPrependValueYAML, err := esyaml.MustPrependYAMLValue(deployment, "claimName", "new-value-")
	if err != nil {
		fmt.Println("Error must prepend set:", err)
		return
	}
	fmt.Println("Must prepend YAML:")
	fmt.Println(mustPrependValueYAML)

	// --- Test parse template value ---
	fmt.Println("\n--- Test parse template value ---")
	parsedYAML, err := esyaml.AddTmlValue(template, "add-on.Name", "added")
	if err != nil {
		fmt.Println("Error parse template set:", err)
		return
	}
	fmt.Println("Parsed template YAML:")
	fmt.Println(parsedYAML)

}
