/*
Copyright © 2024 weidongkl <weidongkx@gmail.com>
*/

package firewalld

import (
	"github.com/godbus/dbus"
)

// Reload firewall rules and keep state information.
// Current permanent configuration will become new runtime configuration,
// i.e. all runtime only changes done until reload are lost with reload if
// they have not been also in permanent configuration.
func (c *Client) Reload() (err error) {
	call, err := c.CallMethod("reload")
	if err != nil {
		return err
	}
	return call.Err
}

// RuntimeToPermanent Make runtime settings permanent.
// Replaces permanent settings with runtime settings for zones, services, icmptypes,
// direct (deprecated) and policies (lockdown whitelist).
func (c *Client) RuntimeToPermanent() (err error) {
	call, err := c.CallMethod("runtimeToPermanent")
	if err != nil {
		return err
	}
	return call.Err
}

// CheckPermanentConfig Run checks on the permanent configuration.
// This is most useful if changes were made manually to configuration files.
func (c *Client) CheckPermanentConfig() (err error) {
	call, err := c.CallMethod("checkPermanentConfig")
	if err != nil {
		return err
	}
	return call.Err
}

// ListServices  Return array of service names (s)
func (c *Client) ListServices() (services []string, err error) {
	if c.opt.Permanent {
		return c.GetServiceNames()
	} else {
		call, err := c.CallMethod("listServices")
		if err != nil {
			return services, err
		}
		err = call.Store(&services)
		return services, err
	}
}

// ListServicesPath  Return array of objects paths (o) of services in permanent configuration.
func (c *Client) ListServicesPath() (servicesPath []string, err error) {
	if !c.opt.Permanent {
		return nil, NotSupportRuntimeErr
	} else {
		call, err := c.CallMethod("listServices")
		if err != nil {
			return servicesPath, err
		}
		err = call.Store(&servicesPath)
		return servicesPath, err
	}
}

// AddZone Add zone with given settings into permanent configuration.
func (c *Client) AddZone(zoneSet ZoneSetting) (err error) {
	if !c.opt.Permanent {
		return NotSupportRuntimeErr
	}
	zsSlice := []interface{}{
		zoneSet.Version,
		zoneSet.Name,
		zoneSet.Description,
		zoneSet.Unused,
		zoneSet.Target,
		zoneSet.Services,
		zoneSet.Ports,
		zoneSet.IcmpBlocks,
		zoneSet.Masquerade,
		zoneSet.ForwardPorts,
		zoneSet.Interfaces,
		zoneSet.SourceAddresses,
		zoneSet.RichRules,
		zoneSet.Protocols,
		zoneSet.SourcePorts,
		zoneSet.IcmpBlockInversion,
	}
	call, err := c.CallMethod("addZone", zoneSet.Name, zsSlice)
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// GetServiceByName Return object path (permanent configuration) of service
// with given name.
func (c *Client) GetServiceByName(service string) (path string, err error) {
	if !c.opt.Permanent {
		return path, NotSupportRuntimeErr
	}
	call, err := c.CallMethod("getServiceByName", service)
	if err != nil {
		return path, err
	}
	err = call.Store(&path)
	return path, err
}

// GetServiceNames  Return list of service names (permanent configuration).
func (c *Client) GetServiceNames() (names []string, err error) {
	if !c.opt.Permanent {
		return names, NotSupportRuntimeErr
	}
	call, err := c.CallMethod("getServiceNames")
	if err != nil {
		return names, err
	}
	err = call.Store(&names)
	return names, err
}

// GetZoneByName Return object path (permanent configuration) of zone with given name.
func (c *Client) GetZoneByName(zone string) (path string, err error) {
	if !c.opt.Permanent {
		return path, NotSupportRuntimeErr
	}
	call, err := c.CallMethod("getZoneByName", zone)
	if err != nil {
		return path, err
	}
	err = call.Store(&path)
	return path, err
}

// GetZoneNames  Return list of zone names (permanent configuration).
func (c *Client) GetZoneNames() (names []string, err error) {
	if !c.opt.Permanent {
		return names, NotSupportRuntimeErr
	}
	call, err := c.CallMethod("getZoneNames")
	if err != nil {
		return names, err
	}
	err = call.Store(&names)
	return names, err
}

// GetZoneOfSource Return name of zone the source is bound to or empty string.
func (c *Client) GetZoneOfSource(source string) (zoneName string, err error) {
	if !c.opt.Permanent {
		return zoneName, NotSupportRuntimeErr
	}
	call, err := c.CallMethod("getZoneOfSource", source)
	if err != nil {
		return zoneName, err
	}
	err = call.Store(&zoneName)
	return zoneName, err
}

// GetServiceSettings Return permanent settings of a service.
func (c *Client) GetServiceSettings(svc string) (svcSet ServiceSetting, err error) {
	var (
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentServiceMethod2(svc, "getSettings")
	} else {
		call, err = c.CallMethod("getServiceSettings", svc)
	}
	if err != nil {
		return svcSet, err
	}
	err = call.Store(&svcSet)
	return svcSet, err
}

// AddForwardPort Add the IPv4 forward port into zone. If zone is empty, use default zone. The port can either be a
// single port number portid or a port range portid-portid. The protocol can either be tcp or udp. The destination
// address is a simple IP address. If timeout(The timeout configuration does not take effect for permanent
// configuration) is non-zero, the operation will be active only for the amount of seconds.
func (c *Client) AddForwardPort(zone, port, protocol, toPort, toAddress string, timeout int) error {
	var (
		call *dbus.Call
		err  error
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "addForwardPort", port, protocol, toPort, toAddress)
	} else {
		call, err = c.CallRuntimeZoneMethod("addForwardPort", zone, port, protocol, toPort, toAddress, timeout)
	}
	if err != nil {
		return err
	}
	return call.Err
}

// AddInterface Bind interface with zone.
func (c *Client) AddInterface(zone, interFace string) error {
	var (
		call *dbus.Call
		err  error
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "addInterface", interFace)
	} else {
		call, err = c.CallRuntimeZoneMethod("addInterface", zone, interFace)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// AddPort when the timeout((The timeout configuration does not take effect for permanent
// configuration) is set to 0, the timeout is ignored.
func (c *Client) AddPort(zone, port, protocol string, timeout int) error {
	var (
		call *dbus.Call
		err  error
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "addPort", port, protocol)
	} else {
		call, err = c.CallRuntimeZoneMethod("addPort", zone, port, protocol, timeout)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// AddProtocol add protocol into zone. The protocol can be any protocol supported by the system. Please have a look at /etc/protocols for supported protocols.
func (c *Client) AddProtocol(zone, protocol string, timeout int) error {
	var (
		call *dbus.Call
		err  error
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "addProtocol", protocol)
	} else {
		call, err = c.CallRuntimeZoneMethod("addProtocol", zone, protocol, timeout)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// AddRichRule add rule to list of rich-language rules in zone.
func (c *Client) AddRichRule(zone, rule string, timeout int) error {
	var (
		call *dbus.Call
		err  error
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "addRichRule", rule)
	} else {
		call, err = c.CallRuntimeZoneMethod("addRichRule", zone, rule, timeout)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// AddService Add service into zone.  If timeout is non-zero,
// the operation will be active only for the amount of seconds.
func (c *Client) AddService(zone, service string, timeout int) error {
	var (
		call *dbus.Call
		err  error
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "addService", service)
	} else {
		call, err = c.CallRuntimeZoneMethod("addService", zone, service, timeout)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// AddSource add source to list of source addresses bound  to zone.
func (c *Client) AddSource(zone, source string, timeout int) error {
	var (
		call *dbus.Call
		err  error
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "addSource", source)
	} else {
		call, err = c.CallRuntimeZoneMethod("addSource", zone, source, timeout)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// AddSourcePort add (port, protocol) to list of source ports of zone.
func (c *Client) AddSourcePort(zone, port, protocol string, timeout int) error {
	var (
		call *dbus.Call
		err  error
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "addSourcePort", port, protocol)
	} else {
		call, err = c.CallRuntimeZoneMethod("addSourcePort", zone, port, protocol, timeout)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// GetActiveZones Return dictionary of currently active zones altogether with interfaces and sources used in these
// zones. Active zones are zones, that have a binding to an interface or source.
func (c *Client) GetActiveZones() (azs map[string]ActivateZone, err error) {
	call, err := c.CallRuntimeZoneMethod("getActiveZones")
	if err != nil {
		return azs, err
	}
	azMap := make(map[string]map[string][]string)
	err = call.Store(&azMap)
	if err != nil {
		return azs, err
	}
	azs = make(map[string]ActivateZone)
	for zoneName, zone := range azMap {
		azs[zoneName] = ActivateZone{
			Interfaces: zone["interfaces"],
			Sources:    zone["sources"],
		}
	}
	return azs, err
}

// GetForwardPorts Get list of (port, protocol, toport, toaddr) defined in zone.
func (c *Client) GetForwardPorts(zone string) (fps ForwardPorts, err error) {
	var (
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getForwardPorts")
	} else {
		call, err = c.CallRuntimeZoneMethod("getForwardPorts", zone)
	}
	if err != nil {
		return fps, err
	}
	err = call.Store(&fps)
	return fps, err
}

// GetInterfaces Return array of interfaces (s) previously bound with zone.
func (c *Client) GetInterfaces(zone string) (Interfaces []string, err error) {
	var (
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getInterfaces")
	} else {
		call, err = c.CallRuntimeZoneMethod("getInterfaces", zone)
	}
	if err != nil {
		return Interfaces, err
	}
	err = call.Store(&Interfaces)
	return Interfaces, err
}

// GetPorts Return array of ports (2-tuple of port and protocol) previously enabled in zone
func (c *Client) GetPorts(zone string) (ports Ports, err error) {
	var (
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getPorts")
	} else {
		call, err = c.CallRuntimeZoneMethod("getPorts", zone)
	}
	if err != nil {
		return ports, err
	}
	var _ports [][]interface{}
	err = call.Store(&_ports)
	if err != nil {
		return nil, err
	}
	ports, err = convertToPorts(_ports)
	return ports, err
}

// GetProtocols Return array of protocols (s) previously enabled in zone.
func (c *Client) GetProtocols(zone string) (protocols []string, err error) {
	var (
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getProtocols")
	} else {
		call, err = c.CallRuntimeZoneMethod("getProtocols", zone)
	}
	if err != nil {
		return protocols, err
	}
	err = call.Store(&protocols)
	return protocols, err
}

// GetRichRules Get list of rich-language rules in zone.
func (c *Client) GetRichRules(zone string) (richRules []string, err error) {
	var (
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getRichRules")
	} else {
		call, err = c.CallRuntimeZoneMethod("getRichRules", zone)
	}
	if err != nil {
		return richRules, err
	}
	err = call.Store(&richRules)
	return richRules, err
}

// GetServices Get list of service names used in zone.
func (c *Client) GetServices(zone string) (services []string, err error) {
	var (
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getServices")
	} else {
		call, err = c.CallRuntimeZoneMethod("getServices", zone)
	}
	if err != nil {
		return services, err
	}
	err = call.Store(&services)
	return services, err
}

// GetSourcePorts Get list of (port, protocol) defined in zone.
func (c *Client) GetSourcePorts(zone string) (ports Ports, err error) {
	var (
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getSourcePorts")
	} else {
		call, err = c.CallRuntimeZoneMethod("getSourcePorts", zone)
	}
	if err != nil {
		return ports, err
	}
	var _ports [][]interface{}
	err = call.Store(&_ports)
	if err != nil {
		return nil, err
	}
	ports, err = convertToPorts(_ports)
	return ports, err
}

// GetSources Get list of source addresses bound to zone.
func (c *Client) GetSources(zone string) (sources []string, err error) {
	var (
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getSources")
	} else {
		call, err = c.CallRuntimeZoneMethod("getSources", zone)
	}
	if err != nil {
		return sources, err
	}
	err = call.Store(&sources)
	return sources, err
}

// RemoveForwardPort remove (port, protocol, toport, toaddr) from  list of forward ports of zone.
func (c *Client) RemoveForwardPort(zone, port, protocol, toPort, toAddress string) error {
	var (
		err  error
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "removeForwardPort", port, protocol, toPort, toAddress)
	} else {
		call, err = c.CallRuntimeZoneMethod("removeForwardPort", zone, port, protocol, toPort, toAddress)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// RemovePort If zone is empty, use default zone.
func (c *Client) RemovePort(zone, port, protocol string) error {
	var (
		err  error
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "removePort", port, protocol)
	} else {
		call, err = c.CallRuntimeZoneMethod("removePort", zone, port, protocol)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// RemoveProtocol remove protocol from zone.
func (c *Client) RemoveProtocol(zone, protocol string) error {
	var (
		err  error
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "removeProtocol", protocol)
	} else {
		call, err = c.CallRuntimeZoneMethod("removeProtocol", zone, protocol)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// RemoveRichRule remove rule from list of rich-language rules  in zone.
func (c *Client) RemoveRichRule(zone, rule string) error {
	var (
		err  error
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "removeRichRule", rule)
	} else {
		call, err = c.CallRuntimeZoneMethod("removeRichRule", zone, rule)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// RemoveService remove service from list of services used in zone.
func (c *Client) RemoveService(zone, service string) error {
	var (
		err  error
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "removeService", service)
	} else {
		call, err = c.CallRuntimeZoneMethod("removeService", zone, service)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// RemoveSource remove source from list of source addresses  bound to zone.
func (c *Client) RemoveSource(zone, source string) error {
	var (
		err  error
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "removeSource", source)
	} else {
		call, err = c.CallRuntimeZoneMethod("removeSource", zone, source)
	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// RemoveSourcePort remove (port, protocol) from list of source ports of zone.
func (c *Client) RemoveSourcePort(zone, port, protocol string) error {
	var (
		err  error
		call *dbus.Call
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "removeSourcePort", port, protocol)
	} else {
		call, err = c.CallRuntimeZoneMethod("removeSourcePort", zone, port, protocol)

	}
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// GetDefaultZone Return default zone.
func (c *Client) GetDefaultZone() (defaultZone string, err error) {
	if c.opt.Permanent {
		return "", NotSupportPermanentErr
	}
	call, err := c.CallMethod("getDefaultZone")
	if err != nil {
		return "", err
	}
	err = call.Store(&defaultZone)
	return defaultZone, err
}

// GetZones Return array of names (s) of predefined zones known to
// current runtime environment.
func (c *Client) GetZones() (zones []string, err error) {
	var call *dbus.Call
	if c.opt.Permanent {
		return nil, NotSupportPermanentErr
	} else {
		call, err = c.CallRuntimeZoneMethod("getZones")
	}
	if err != nil {
		return zones, err
	}
	err = call.Store(&zones)
	return
}

// ListZones List object paths of zones known to permanent environment.
func (c *Client) ListZones() (zonesPath []string, err error) {
	if !c.opt.Permanent {
		return nil, NotSupportRuntimeErr
	}
	call, err := c.CallMethod("listZones")
	if err != nil {
		return zonesPath, err
	}
	err = call.Store(&zonesPath)
	return
}

func (c *Client) GetZoneSettings(zone string) (zs ZoneSetting, err error) {
	call, err := c.CallMethod("getZoneSettings", zone)
	if err != nil {
		return zs, err
	}
	err = call.Store(&zs)
	return zs, err
}

// not implemented yet
//func (c *Client) GetZoneSettings2(zone string) (zs ZoneSetting, err error) {
//	return zs, UnimplementedErr
//}

// SetDefaultZone Set default zone for connections and interfaces where no zone has been selected to zone.
// Setting the default zone changes the zone for the connections or interfaces,
// that are using the default zone. This is a runtime and permanent change.
func (c *Client) SetDefaultZone(zone string) (err error) {
	call, err := c.CallMethod("setDefaultZone", zone)
	if err != nil {
		return err
	}
	return call.Err
}

// SetForwardPorts Permanently set forward ports of zone
func (c *Client) SetForwardPorts(zone string, fps ForwardPorts) error {
	if !c.opt.Permanent {
		return NotSupportRuntimeErr
	}
	var (
		call *dbus.Call
		err  error
	)
	var dpSlice [][]string
	for _, fp := range fps {
		dpSlice = append(dpSlice, []string{fp.Port, fp.Protocol, fp.ToAddress, fp.ToPort})
	}
	call, err = c.CallPermanentZoneMethod2(zone, "setForwardPorts", dpSlice)
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// SetPorts Permanently set ports of zone
func (c *Client) SetPorts(zone string, ports Ports) error {
	if !c.opt.Permanent {
		return NotSupportRuntimeErr
	}
	var (
		call *dbus.Call
		err  error
	)
	var psSlice [][]string
	for _, port := range ports {
		psSlice = append(psSlice, []string{port.Port, port.Protocol})
	}
	call, err = c.CallPermanentZoneMethod2(zone, "setPorts", psSlice)
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// SetProtocols Permanently set list of protocols used in zone to protocols.
func (c *Client) SetProtocols(zone string, protocols []string) error {
	if !c.opt.Permanent {
		return NotSupportRuntimeErr
	}
	var (
		call *dbus.Call
		err  error
	)
	call, err = c.CallPermanentZoneMethod2(zone, "setProtocols", protocols)
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// SetRichRules Permanently set list of rich-language rules to rules.
func (c *Client) SetRichRules(zone string, rules []string) error {
	if !c.opt.Permanent {
		return NotSupportRuntimeErr
	}
	var (
		call *dbus.Call
		err  error
	)
	call, err = c.CallPermanentZoneMethod2(zone, "setRichRules", rules)
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// SetServices Permanently set list of services used in zone to services.
func (c *Client) SetServices(zone string, services []string) error {
	if !c.opt.Permanent {
		return NotSupportRuntimeErr
	}
	var (
		call *dbus.Call
		err  error
	)
	call, err = c.CallPermanentZoneMethod2(zone, "setServices", services)
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// SetSourcePorts Permanently set source-ports of zone to list
func (c *Client) SetSourcePorts(zone string, ports Ports) error {
	if !c.opt.Permanent {
		return NotSupportRuntimeErr
	}
	var (
		call *dbus.Call
		err  error
	)
	var psSlice [][]string
	for _, port := range ports {
		psSlice = append(psSlice, []string{port.Port, port.Protocol})
	}
	call, err = c.CallPermanentZoneMethod2(zone, "setSourcePorts", psSlice)
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// SetSources Permanently set list of source addresses bound to zone to sources.
func (c *Client) SetSources(zone string, sources []string) error {
	if !c.opt.Permanent {
		return NotSupportRuntimeErr
	}
	var (
		call *dbus.Call
		err  error
	)
	call, err = c.CallPermanentZoneMethod2(zone, "setSources", sources)
	if err != nil {
		return err
	}
	err = call.Err
	return err
}

// get zone object id by name
func (c *Client) getZoneID(zone string) (zoneId int, err error) {
	zonePath, err := c.GetZoneByName(zone)
	if err != nil {
		return 0, err
	}
	zoneId, err = getIdByPath(zonePath)
	return zoneId, err
}

func (c *Client) getServiceID(svc string) (svcId int, err error) {
	svcPath, err := c.GetServiceByName(svc)
	if err != nil {
		return 0, err
	}
	svcId, err = getIdByPath(svcPath)
	return svcId, err
}
