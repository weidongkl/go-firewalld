/*
Copyright © 2024 weidongkl <weidongkx@gmail.com>
*/

package firewalld

// Options keeps the settings to set up firewalld connection.
type Options struct {
	//dbusRuntimePath          dbus.ObjectPath
	//dbusRuntimeInterface     string
	//dbusRuntimeZoneInterface string
	//dbusPermanentPath        dbus.ObjectPath
	//dbusPermanentInterface   string
	Zone      string
	Permanent bool
}

func (opt *Options) init() {
	//if opt.dbusRuntimePath == "" {
	//	opt.dbusRuntimePath = dbusRuntimePath
	//}
	//if opt.dbusPermanentPath == "" {
	//	opt.dbusPermanentPath = dbusPermanentPath
	//}
	//opt.dbusRuntimeInterface = dbusRuntimeInterface
	//opt.dbusPermanentInterface = dbusPermanentInterface
}
