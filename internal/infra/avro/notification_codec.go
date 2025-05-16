package avro

import (
    "fmt"
    "os"

    "github.com/linkedin/goavro/v2"
    "github.com/z6wdc/go-avro/internal/entity"
)

type NotificationCodec struct {
    codec *goavro.Codec
}

func NewNotificationCodec(schemaPath string) (*NotificationCodec, error) {
    schemaData, err := os.ReadFile(schemaPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read schema: %w", err)
    }

    codec, err := goavro.NewCodec(string(schemaData))
    if err != nil {
        return nil, fmt.Errorf("failed to create codec: %w", err)
    }

    return &NotificationCodec{codec: codec}, nil
}

func (c *NotificationCodec) Encode(n *entity.Notification) ([]byte, error) {
    native := map[string]interface{}{
        "id":      n.ID,
        "userId":  n.UserID,
        "message": n.Message,
    }

    binary, err := c.codec.BinaryFromNative(nil, native)
    if err != nil {
        return nil, fmt.Errorf("encode error: %w", err)
    }

    return binary, nil
}

func (c *NotificationCodec) Decode(data []byte) (*entity.Notification, error) {
    native, _, err := c.codec.NativeFromBinary(data)
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

// Interface assertions
var _ NotificationEncoder = (*NotificationCodec)(nil)
var _ NotificationDecoder = (*NotificationCodec)(nil)
