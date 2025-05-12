package entity_test

import (
	"testing"

	"github.com/z6wdc/go-avro/internal/entity"
)

func TestIsValidNotificationType(t *testing.T) {
	tests := []struct {
		name     string
		input    entity.NotificationType
		expected bool
	}{
		{"valid email", entity.NotificationEmail, true},
		{"valid push", entity.NotificationPush, true},
		{"valid sms", entity.NotificationSMS, true},
		{"invalid type", entity.NotificationType("fax"), false},
		{"empty string", entity.NotificationType(""), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := entity.IsValidNotificationType(tc.input)
			if actual != tc.expected {
				t.Errorf("input %q: expected %v, got %v", tc.input, tc.expected, actual)
			}
		})
	}
}
