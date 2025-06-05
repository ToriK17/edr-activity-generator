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
		// add a YAML doc separator if the file is not empty
		if fi, _ := file.Stat(); fi.Size() != 0 {
			if _, err := file.WriteString("---\n"); err != nil {
				return err
			}
		}

		data, err := yaml.Marshal(entry)
		if err != nil {
			return fmt.Errorf("failed to marshal YAML: %w", err)
		}
		// Write the doc and ensure it ends with a newline
		if _, err := file.Write(data); err != nil {
			return err
		}

		return err

	case "csv":
		writer := csv.NewWriter(file)
		defer writer.Flush()

		// If the file is empty, write a header row first.
		if fi, _ := file.Stat(); fi.Size() == 0 {
			header, err := toCSVHeader(entry)
			if err != nil {
				return err
			}
			if err := writer.Write(header); err != nil {
				return err
			}
		}

		record, err := toCSVRecord(entry)
		if err != nil {
			return err
		}
		return writer.Write(record)

	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}
