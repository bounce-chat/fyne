package mobile

type NotificationCallback interface {
	SetNotificationCallback(func(string))
}
