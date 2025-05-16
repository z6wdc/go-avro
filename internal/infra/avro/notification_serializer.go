package avro

import (
	"fmt"
	"os"

	"github.com/linkedin/goavro/v2"
	"github.com/z6wdc/go-avro/internal/entity"
)

type NotificationSerializer struct {
	codec *goavro.Codec
}

func NewNotificationSerializer(schemaPath string) (*NotificationSerializer, error) {
	schemaData, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema: %w", err)
	}

	codec, err := goavro.NewCodec(string(schemaData))
	if err != nil {
		return nil, fmt.Errorf("failed to create codec: %w", err)
	}

	return &NotificationSerializer{codec: codec}, nil
}

func (s *NotificationSerializer) Encode(n *entity.Notification) ([]byte, error) {
	native := map[string]interface{}{
		"id":      n.ID,
		"userId":  n.UserID,
		"message": n.Message,
	}

	binary, err := s.codec.BinaryFromNative(nil, native)
	if err != nil {
		return nil, fmt.Errorf("encode error: %w", err)
	}

	return binary, nil
}

func (s *NotificationSerializer) Decode(data []byte) (*entity.Notification, error) {
	native, _, err := s.codec.NativeFromBinary(data)
	if err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	m := native.(map[string]interface{})

	return &entity.Notification{
		ID:      m["id"].(string),
		UserID:  int(m["userId"].(int32)),
		Message: m["message"].(string),
	}, nil
}
