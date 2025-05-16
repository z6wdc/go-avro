package usecase

import (
	"fmt"

	"github.com/z6wdc/go-avro/internal/entity"
	"github.com/z6wdc/go-avro/internal/infra/avro"
)

type DecodeNotificationUseCase struct {
	Decoder avro.NotificationDecoder
}

func NewDecodeNotificationUseCase(decoder avro.NotificationDecoder) *DecodeNotificationUseCase {
	return &DecodeNotificationUseCase{Decoder: decoder}
}

func (u *DecodeNotificationUseCase) Execute(data []byte) (*entity.Notification, error) {
	decoded, err := u.Decoder.Decode(data)
	if err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}
	return decoded, nil
}
