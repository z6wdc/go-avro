package usecase

import (
	"fmt"

	"github.com/z6wdc/go-avro/internal/entity"
	"github.com/z6wdc/go-avro/internal/infra/avro"
)

type EncodeNotificationUseCase struct {
	Encoder avro.NotificationEncoder
}

func NewEncodeNotificationUseCase(encoder avro.NotificationEncoder) *EncodeNotificationUseCase {
	return &EncodeNotificationUseCase{Encoder: encoder}
}

func (u *EncodeNotificationUseCase) Execute(n *entity.Notification) ([]byte, error) {
	encoded, err := u.Encoder.Encode(n)
	if err != nil {
		return nil, fmt.Errorf("encode error: %w", err)
	}
	return encoded, nil
}
