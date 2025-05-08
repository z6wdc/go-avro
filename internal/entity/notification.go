package entity

type NotificationType string

const (
    NotificationEmail NotificationType = "email"
    NotificationPush  NotificationType = "push"
    NotificationSMS   NotificationType = "sms"
)

func IsValidNotificationType(t NotificationType) bool {
    switch t {
    case NotificationEmail, NotificationPush, NotificationSMS:
        return true
    default:
        return false
    }
}

type Notification struct {
    ID      string
    UserID  int
    Message string
    Type    NotificationType
}
