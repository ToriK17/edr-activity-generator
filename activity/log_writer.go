package activity

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func logWriter(entry any, outputPath, format string) error {
	file, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	switch format {
	case "json":
		encoder := json.NewEncoder(file)
		return encoder.Encode(entry)

	case "yaml":
		data, err := yaml.Marshal(entry)
		if err != nil {
			return fmt.Errorf("failed to marshal YAML: %w", err)
		}
		_, err = file.Write(data)
		return err

	case "csv":
		writer := csv.NewWriter(file)
		defer writer.Flush()

		// For now, support a very basic CSV with `map[string]string` only
		record, err := toCSVRecord(entry)
		if err != nil {
			return err
		}
		return writer.Write(record)

	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}
