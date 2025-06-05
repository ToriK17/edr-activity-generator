package activity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// mock struct
type sample struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

func TestToCSVRecord(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		want    []string
		wantErr bool
	}{
		{
			name:  "happy path",
			input: sample{Foo: "hello", Bar: 42},
			want:  []string{"hello", "42"},
		},
		{
			name:    "non-struct input",
			input:   "not a struct",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := toCSVRecord(tc.input)
			if (err != nil) != tc.wantErr {
				t.Fatalf("error expectation mismatch: %v", err)
			}
			if !tc.wantErr && len(got) != len(tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestWriteLog_JSONAndUnsupported(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "out.json")

	// JSON happy-path
	err := logWriter(sample{Foo: "baz", Bar: 7}, tmpFile, "json")
	if err != nil {
		t.Fatalf("writeLog json unexpected error: %v", err)
	}

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("reading file: %v", err)
	}

	var got sample
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal json: %v", err)
	}
	if got.Foo != "baz" || got.Bar != 7 {
		t.Fatalf("round-trip mismatch: %+v", got)
	}

	// Unsupported format should error
	if err := logWriter(sample{}, tmpFile, "ini"); err == nil {
		t.Fatalf("expected error for unsupported format, got nil")
	}
}
