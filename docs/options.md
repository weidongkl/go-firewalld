```go
type Options struct {
	DbusRuntimePath        dbus.ObjectPath
	dbusRuntimeInterface   string
	DbusPermanentPath      dbus.ObjectPath
	dbusPermanentInterface string

	Permanent bool
}
```

通过传入Permanent参数来设置查看运行时还是配置中的firewalld信息