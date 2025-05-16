package avro_test

import (
	"path/filepath"
	"testing"

	"github.com/z6wdc/go-avro/internal/entity"
	"github.com/z6wdc/go-avro/internal/infra/avro"
	"github.com/z6wdc/go-avro/internal/test"
)

func TestNotificationSerializer(t *testing.T) {
	// Specify the path to the Avro schema
	schemaPath := filepath.Join(test.GetProjectRoot(), "internal", "infra", "avro", "schema_v1.avsc")

	// Create a serializer
	serializer, err := avro.NewNotificationSerializer(schemaPath)
	if err != nil {
		t.Fatalf("failed to create serializer: %v", err)
	}

	// Input data for testing
	input := &entity.Notification{
		ID:      "n001",
		UserID:  42,
		Message: "Hello, AVRO!",
	}

	// Encode the input data into Avro binary
	data, err := serializer.Encode(input)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	if len(data) == 0 {
		t.Error("encoded data should not be empty")
	}

	// Decode the binary back into a Notification entity
	output, err := serializer.Decode(data)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Validate the decoded result matches the original
	if output.ID != input.ID {
		t.Errorf("ID mismatch: got %s, want %s", output.ID, input.ID)
	}
	if output.UserID != input.UserID {
		t.Errorf("UserID mismatch: got %d, want %d", output.UserID, input.UserID)
	}
	if output.Message != input.Message {
		t.Errorf("Message mismatch: got %s, want %s", output.Message, input.Message)
	}
}
