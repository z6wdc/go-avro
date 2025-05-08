package entity

import "testing"

func TestIsValidNotificationType(t *testing.T) {
    tests := []struct {
        name     string
        input    NotificationType
        expected bool
    }{
        {"valid email", NotificationEmail, true},
        {"valid push", NotificationPush, true},
        {"valid sms", NotificationSMS, true},
        {"invalid type", NotificationType("fax"), false},
        {"empty string", NotificationType(""), false},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            actual := IsValidNotificationType(tc.input)
            if actual != tc.expected {
                t.Errorf("input %q: expected %v, got %v", tc.input, tc.expected, actual)
            }
        })
    }
}
