//go:generate mockgen -source=notification_decoder.go -destination=../../mocks/notification_decoder_mock.go -package=mocks

package avro

import "github.com/z6wdc/go-avro/internal/entity"

type NotificationDecoder interface {
    Decode([]byte) (*entity.Notification, error)
}
