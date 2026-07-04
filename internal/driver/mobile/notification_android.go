//go:build android

package mobile

import (
	"fyne.io/fyne/v2/internal/driver/mobile/app"
)

func (*device) SetNotificationCallback(callback func(string)) {
	app.NativeSetNotificationCallback(callback)
}
