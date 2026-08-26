package printer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

func hasNoResources(namespace string, length int) bool {
	if length == 0 {
		fmt.Printf("No resources found in %s namespace.\n", namespace)
		return true
	}
	return false
}

func printJsonFormat(out io.Writer, responseBody []byte) error {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, responseBody, "", "  "); err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}
	_, err := out.Write(prettyJSON.Bytes())
	return err
}

func printYamlFormat(out io.Writer, responseBody []byte) error {
	var yamlData any
	if err := json.Unmarshal(responseBody, &yamlData); err != nil {
		return fmt.Errorf("failed to parse JSON for YAML conversion: %w", err)
	}
	yamlBytes, err := yaml.Marshal(yamlData)
	if err != nil {
		return fmt.Errorf("failed to convert JSON to YAML: %w", err)
	}
	_, err = out.Write(yamlBytes)
	return err
}
