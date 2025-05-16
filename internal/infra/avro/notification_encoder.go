//go:generate mockgen -source=notification_encoder.go -destination=../../mocks/notification_encoder_mock.go -package=mocks

package avro

import "github.com/z6wdc/go-avro/internal/entity"

type NotificationEncoder interface {
    Encode(*entity.Notification) ([]byte, error)
}
