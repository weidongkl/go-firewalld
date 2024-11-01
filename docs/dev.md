### [Constants](https://pkg.go.dev/gitee.com/weidongkl/go-firewalld#pkg-constants)

This section is empty.

###  [Variables](https://pkg.go.dev/gitee.com/weidongkl/go-firewalld#pkg-variables)

```go
var (
	NotSupportPermanentErr = errors.New("this method not supported permanent call")
	UnimplementedErr       = errors.New("this method is not yet implemented")
	NotSupportRuntimeErr   = errors.New("this method not supported Runtime call")
)
```

### [Functions](https://pkg.go.dev/gitee.com/weidongkl/go-firewalld#pkg-functions)

#### [func Version](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/version.go#L8)

```go
func Version() string
```

Version is the current release version.

### [Types](https://pkg.go.dev/gitee.com/weidongkl/go-firewalld#pkg-types)

#### type [ActivateZone](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/defines.go#L54) 

```go
type ActivateZone struct {
	Interfaces []string
	Sources    []string
}
```

#### type [Client](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/firewalld.go#L16)

```go
type Client struct {
	// contains filtered or unexported fields
}
```

#### func [NewClient](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/firewalld.go#L25)

```go
func NewClient(opt *Options) (*Client, error)
```

#### func (*Client) [AddForwardPort](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L190) 

```go
func (c *Client) AddForwardPort(zone, port, protocol, toPort, toAddress string, timeout int) error
```

AddForwardPort Add the IPv4 forward port into zone. If zone is empty, use default zone. The port can either be a single port number portid or a port range portid-portid. The protocol can either be tcp or udp. The destination address is a simple IP address. If timeout(The timeout configuration does not take effect for permanent configuration) is non-zero, the operation will be active only for the amount of seconds.

#### func (*Client) [AddInterface](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L207) 

```go
func (c *Client) AddInterface(zone, interFace string) error
```

AddInterface Bind interface with zone.

#### func (*Client) [AddPort](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L226) 

```go
func (c *Client) AddPort(zone, port, protocol string, timeout int) error
```

AddPort when the timeout((The timeout configuration does not take effect for permanent configuration) is set to 0, the timeout is ignored.

#### func (*Client) [AddProtocol](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L244) 

```go
func (c *Client) AddProtocol(zone, protocol string, timeout int) error
```

AddProtocol add protocol into zone. The protocol can be any protocol supported by the system. Please have a look at /etc/protocols for supported protocols.

#### func (*Client) [AddRichRule](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L262) 

```go
func (c *Client) AddRichRule(zone, rule string, timeout int) error
```

AddRichRule add rule to list of rich-language rules in zone.

#### func (*Client) [AddService](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L281) 

```go
func (c *Client) AddService(zone, service string, timeout int) error
```

AddService Add service into zone. If timeout is non-zero, the operation will be active only for the amount of seconds.

#### func (*Client) [AddSource](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L299) 

```go
func (c *Client) AddSource(zone, source string, timeout int) error
```

AddSource add source to list of source addresses bound to zone.

#### func (*Client) [AddSourcePort](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L317) 

```go
func (c *Client) AddSourcePort(zone, port, protocol string, timeout int) error
```

AddSourcePort add (port, protocol) to list of source ports of zone.

#### func (*Client) [AddZone](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L73)

```go
func (c *Client) AddZone(zoneSet ZoneSetting) (err error)
```

AddZone Add zone with given settings into permanent configuration.

#### func (*Client) [CallMethod](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/firewalld.go#L51)

```go
func (c *Client) CallMethod(method string, args ...interface{}) (*dbus.Call, error)
```

#### func (*Client) [CallPermanentServiceMethod](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/firewalld.go#L74) 

```go
func (c *Client) CallPermanentServiceMethod(svcId int, method string, args ...interface{}) (*dbus.Call, error)
```

#### func (*Client) [CallPermanentServiceMethod2](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/firewalld.go#L80)
```go
func (c *Client) CallPermanentServiceMethod2(svc string, method string, args ...interface{}) (*dbus.Call, error)
```

#### func (*Client) [CallPermanentZoneMethod](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/firewalld.go#L59) 

```go
func (c *Client) CallPermanentZoneMethod(zoneId int, method string, args ...interface{}) (*dbus.Call, error)
```

#### func (*Client) [CallPermanentZoneMethod2](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/firewalld.go#L65) 

```go
func (c *Client) CallPermanentZoneMethod2(zone string, method string, args ...interface{}) (*dbus.Call, error)
```

#### func (*Client) [CallRuntimeZoneMethod](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/firewalld.go#L55) 

```go
func (c *Client) CallRuntimeZoneMethod(method string, args ...interface{}) (*dbus.Call, error)
```

#### func (*Client) [CheckPermanentConfig](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L36) 

```go
func (c *Client) CheckPermanentConfig() (err error)
```

CheckPermanentConfig Run checks on the permanent configuration. This is most useful if changes were made manually to configuration files.

#### func (*Client) [Close](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/firewalld.go#L47)

```go
func (c *Client) Close() error
```

#### func (*Client) [GetActiveZones](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L336) 

```go
func (c *Client) GetActiveZones() (azs map[string]ActivateZone, err error)
```

GetActiveZones Return dictionary of currently active zones altogether with interfaces and sources used in these zones. Active zones are zones, that have a binding to an interface or source.

#### func (*Client) [GetDefaultZone](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L630) 

```go
func (c *Client) GetDefaultZone() (defaultZone string, err error)
```

GetDefaultZone Return default zone.

#### func (*Client) [GetForwardPorts](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L357)

```go
func (c *Client) GetForwardPorts(zone string) (fps ForwardPorts, err error)
```

GetForwardPorts Get list of (port, protocol, toport, toaddr) defined in zone.

#### func (*Client) [GetInterfaces](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L374)

```go
func (c *Client) GetInterfaces(zone string) (Interfaces []string, err error)
```

GetInterfaces Return array of interfaces (s) previously bound with zone.

#### func (*Client) [GetPorts](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L391)

```go
func (c *Client) GetPorts(zone string) (ports Ports, err error)
```

GetPorts Return array of ports (2-tuple of port and protocol) previously enabled in zone

#### func (*Client) [GetProtocols](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L413)

```go
func (c *Client) GetProtocols(zone string) (protocols []string, err error)
```

GetProtocols Return array of protocols (s) previously enabled in zone.

#### func (*Client) [GetRichRules](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L430) 

```go
func (c *Client) GetRichRules(zone string) (richRules []string, err error)
```

GetRichRules Get list of rich-language rules in zone.

#### func (*Client) [GetServiceByName](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L105)

```go
func (c *Client) GetServiceByName(service string) (path string, err error)
```

GetServiceByName Return object path (permanent configuration) of service with given name.

#### func (*Client) [GetServiceNames](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L118) 
```go
func (c *Client) GetServiceNames() (names []string, err error)
```

GetServiceNames Return list of service names (permanent configuration).

#### func (*Client) [GetServiceSettings](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L170) 

```go
func (c *Client) GetServiceSettings(svc string) (svcSet ServiceSetting, err error)
```

GetServiceSettings Return permanent settings of a service.

#### func (*Client) [GetServices](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L447) 

```go
func (c *Client) GetServices(zone string) (services []string, err error)
```

GetServices Get list of service names used in zone.

#### func (*Client) [GetSourcePorts](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L464)
```go
func (c *Client) GetSourcePorts(zone string) (ports Ports, err error)
```

GetSourcePorts Get list of (port, protocol) defined in zone.

#### func (*Client) [GetSources](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L486) 

```go
func (c *Client) GetSources(zone string) (sources []string, err error)
```

GetSources Get list of source addresses bound to zone.

#### func (*Client) [GetZoneByName](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L131) 

```go
func (c *Client) GetZoneByName(zone string) (path string, err error)
```

GetZoneByName Return object path (permanent configuration) of zone with given name.

#### func (*Client) [GetZoneNames](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L144) 

```go
func (c *Client) GetZoneNames() (names []string, err error)
```

GetZoneNames Return list of zone names (permanent configuration).

#### func (*Client) [GetZoneOfSource](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L157) 

```go
func (c *Client) GetZoneOfSource(source string) (zoneName string, err error)
```

GetZoneOfSource Return name of zone the source is bound to or empty string.

#### func (*Client) [GetZoneSettings](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L671)

```go
func (c *Client) GetZoneSettings(zone string) (zs ZoneSetting, err error)
```

#### func (*Client) [GetZones](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L644)

```go
func (c *Client) GetZones() (zones []string, err error)
```

GetZones Return array of names (s) of predefined zones known to current runtime environment.

#### func (*Client) [ListServices](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L45) 

```go
func (c *Client) ListServices() (services []string, err error)
```

ListServices Return array of service names (s)

#### func (*Client) [ListServicesPath](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L59)

```go
func (c *Client) ListServicesPath() (servicesPath []string, err error)
```

ListServicesPath Return array of objects paths (o) of services in permanent configuration.

#### func (*Client) [ListZones](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L659) 

```go
func (c *Client) ListZones() (zonesPath []string, err error)
```

ListZones List object paths of zones known to permanent environment.

#### func (*Client) [Reload](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L15) 

```go
func (c *Client) Reload() (err error)
```

Reload firewall rules and keep state information. Current permanent configuration will become new runtime configuration, i.e. all runtime only changes done until reload are lost with reload if they have not been also in permanent configuration.

#### func (*Client) [RemoveForwardPort](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L503)

```go
func (c *Client) RemoveForwardPort(zone, port, protocol, toPort, toAddress string) error
```

RemoveForwardPort remove (port, protocol, toport, toaddr) from list of forward ports of zone.

#### func (*Client) [RemovePort](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L521) 

```go
func (c *Client) RemovePort(zone, port, protocol string) error
```

RemovePort If zone is empty, use default zone.

#### func (*Client) [RemoveProtocol](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L539) 

```go
func (c *Client) RemoveProtocol(zone, protocol string) error
```

RemoveProtocol remove protocol from zone.

#### func (*Client) [RemoveRichRule](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L557)

```go
func (c *Client) RemoveRichRule(zone, rule string) error
```

RemoveRichRule remove rule from list of rich-language rules in zone.

#### func (*Client) [RemoveService](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L575) 

```go
func (c *Client) RemoveService(zone, service string) error
```

RemoveService remove service from list of services used in zone.

#### func (*Client) [RemoveSource](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L593)
```go
func (c *Client) RemoveSource(zone, source string) error
```

RemoveSource remove source from list of source addresses bound to zone.

#### func (*Client) [RemoveSourcePort](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L611) 

```go
func (c *Client) RemoveSourcePort(zone, port, protocol string) error
```

RemoveSourcePort remove (port, protocol) from list of source ports of zone.

#### func (*Client) [RuntimeToPermanent](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L26)

```go
func (c *Client) RuntimeToPermanent() (err error)
```

RuntimeToPermanent Make runtime settings permanent. Replaces permanent settings with runtime settings for zones, services, icmptypes, direct (deprecated) and policies (lockdown whitelist).

#### func (*Client) [SetDefaultZone](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L688) 

```go
func (c *Client) SetDefaultZone(zone string) (err error)
```

SetDefaultZone Set default zone for connections and interfaces where no zone has been selected to zone. Setting the default zone changes the zone for the connections or interfaces, that are using the default zone. This is a runtime and permanent change.

#### func (*Client) [SetForwardPorts](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L697) 

```go
func (c *Client) SetForwardPorts(zone string, fps ForwardPorts) error
```

SetForwardPorts Permanently set forward ports of zone

#### func (*Client) [SetPorts](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L718) 

```go
func (c *Client) SetPorts(zone string, ports Ports) error
```

SetPorts Permanently set ports of zone

#### func (*Client) [SetProtocols](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L739) 

```go
func (c *Client) SetProtocols(zone string, protocols []string) error
```

SetProtocols Permanently set list of protocols used in zone to protocols.

#### func (*Client) [SetRichRules](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L756) 

```go
func (c *Client) SetRichRules(zone string, rules []string) error
```

SetRichRules Permanently set list of rich-language rules to rules.

#### func (*Client) [SetServices](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L773)

```go
func (c *Client) SetServices(zone string, services []string) error
```

SetServices Permanently set list of services used in zone to services.

#### func (*Client) [SetSourcePorts](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L790) 

```go
func (c *Client) SetSourcePorts(zone string, ports Ports) error
```

SetSourcePorts Permanently set source-ports of zone to list

#### func (*Client) [SetSources](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/method.go#L811) 

```go
func (c *Client) SetSources(zone string, sources []string) error
```

SetSources Permanently set list of source addresses bound to zone to sources.

#### type [ForwardPort](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/defines.go#L26) 

```go
type ForwardPort struct {
	Port      string
	Protocol  string
	ToPort    string
	ToAddress string
}
```

#### type [ForwardPorts](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/defines.go#L33) 

```go
type ForwardPorts []ForwardPort
```

#### type [Options](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/options.go#L8)

```go
type Options struct {
	//dbusRuntimePath          dbus.ObjectPath
	//dbusRuntimeInterface     string
	//dbusRuntimeZoneInterface string
	//dbusPermanentPath        dbus.ObjectPath
	//dbusPermanentInterface   string
	Zone      string
	Permanent bool
}
```

Options keeps the settings to set up firewalld connection.

#### type [Port](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/defines.go#L21) 

```go
type Port struct {
	Port     string
	Protocol string
}
```

#### type [Ports](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/defines.go#L32) 

```go
type Ports []Port
```

#### type [ServiceSetting](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/defines.go#L60)
```go
type ServiceSetting struct {
	Version      string
	Name         string
	Description  string
	Ports        Ports
	ModuleNames  []string
	Destinations map[string]string
	Protocols    []string
	SourcePorts  Ports
}
```

#### type [ServiceSettingMap](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/defines.go#L59) 

```go
type ServiceSettingMap map[string]interface{}
```

#### func (ServiceSettingMap) [ToStruct](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/defines.go#L72) 

```go
func (ssm ServiceSettingMap) ToStruct() (ss ServiceSetting, err error)
```

#### type [ZoneSetting](https://gitee.com/weidongkl/go-firewalld/blob/v1.0.0/defines.go#L35) 

```go
type ZoneSetting struct {
	Version            string
	Name               string
	Description        string
	Unused             bool
	Target             string
	Services           []string
	Ports              Ports
	IcmpBlocks         []string
	Masquerade         bool
	ForwardPorts       ForwardPorts
	Interfaces         []string
	SourceAddresses    []string
	RichRules          []string
	Protocols          []string
	SourcePorts        Ports
	IcmpBlockInversion bool
}
```

