# esyaml

`esyaml` is a Go library for manipulating YAML data. It provides functions to get, set, delete, replace, and insert values in YAML documents.

## Features

- **Get YAML Value**: Retrieve a value from a YAML document using a dot-separated path.
- **Set YAML Value**: Update a value in a YAML document at a specified path.
- **Delete YAML Field**: Remove a field from a YAML document.
- **Replace YAML Key**: Replace a key in a YAML document.
- **Insert YAML Value**: Insert a new value into a YAML document.

## Installation

To use this library in your Go project, you need to install the `gopkg.in/yaml.v3` package:

```sh
go get gopkg.in/yaml.v3