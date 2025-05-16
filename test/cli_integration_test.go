package test

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/z6wdc/go-avro/pkg/util"
)

func TestCLICanEncodeAndDecodeNotification(t *testing.T) {
	root := util.GetProjectRoot()
	bin := filepath.Join(root, "go-avro")

	inputPath := filepath.Join(root, "testdata", "input.json")
	encodedPath := filepath.Join(root, "testdata", "output.avro")
	decodedPath := filepath.Join(root, "testdata", "output.json")

	_ = os.Remove(encodedPath)
	_ = os.Remove(decodedPath)

	cmd := exec.Command(bin, "encode", "--input", inputPath, "--output", encodedPath)
	cmd.Dir = root
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("encode failed: %v\n%s", err, string(output))
	}

	if _, err := os.Stat(encodedPath); err != nil {
		t.Fatalf("encoded file not found: %v", err)
	}

	cmd = exec.Command(bin, "decode", "--input", encodedPath, "--output", decodedPath)
	cmd.Dir = root
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("decode failed: %v\n%s", err, string(output))
	}

	data, err := os.ReadFile(decodedPath)
	if err != nil {
		t.Fatalf("failed to read decoded output: %v", err)
	}

	var result struct {
		ID      string `json:"id"`
		UserID  int    `json:"userId"`
		Message string `json:"message"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}

	if result.ID != "cli-123" {
		t.Errorf("expected ID = cli-123, got %s", result.ID)
	}
	if result.UserID != 99 {
		t.Errorf("expected UserID = 99, got %d", result.UserID)
	}
	if result.Message != "Hello, Avro CLI!" {
		t.Errorf("expected Message = Hello, Avro CLI!, got %s", result.Message)
	}

}
