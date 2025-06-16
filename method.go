/*
Copyright © 2024 weidongkl <weidongkx@gmail.com>
*/

package firewalld

import (
	"fmt"

	"github.com/godbus/dbus"
)

// Helper function to handle common DBus call patterns
func (c *Client) handleCall(call *dbus.Call, err error) error {
	if err != nil {
		return fmt.Errorf("dbus call failed: %w", err)
	}
	return call.Err
}

// Helper function to handle DBus calls that return a single value
func (c *Client) handleDBusCall(call *dbus.Call, err error, result interface{}) error {
	if err != nil {
		return fmt.Errorf("dbus call failed: %w", err)
	}
	if err := call.Store(result); err != nil {
		return fmt.Errorf("failed to store result: %w", err)
	}
	return nil
}

// Reload firewall rules and keep state information.
// Current permanent configuration will become new runtime configuration,
// i.e. all runtime only changes done until reload are lost with reload if
// they have not been also in permanent configuration.
func (c *Client) Reload() error {
	call, err := c.CallMethod("reload")
	return c.handleCall(call, err)
}

// RuntimeToPermanent Make runtime settings permanent.
// Replaces permanent settings with runtime settings for zones, services, icmptypes,
// direct (deprecated) and policies (lockdown whitelist).
func (c *Client) RuntimeToPermanent() error {
	call, err := c.CallMethod("runtimeToPermanent")
	return c.handleCall(call, err)
}

// CheckPermanentConfig Run checks on the permanent configuration.
// This is most useful if changes were made manually to configuration files.
func (c *Client) CheckPermanentConfig() error {
	call, err := c.CallMethod("checkPermanentConfig")
	return c.handleCall(call, err)
}

// ListServices Return array of service names
func (c *Client) ListServices() ([]string, error) {
	if c.opt.Permanent {
		return c.GetServiceNames()
	}
	var services []string
	call, err := c.CallMethod("listServices")
	if err := c.handleDBusCall(call, err, &services); err != nil {
		return nil, err
	}
	return services, nil
}

// ListServicesPath  Return array of objects paths (o) of services in permanent configuration.
func (c *Client) ListServicesPath() (servicesPath []string, err error) {
	if !c.opt.Permanent {
		return nil, ErrNotSupportRuntime
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
func (c *Client) AddZone(zoneSet ZoneSetting) error {
	if !c.opt.Permanent {
		return ErrNotSupportRuntime
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
	return c.handleCall(call, err)
}

// GetServiceByName Return object path (permanent configuration) of service
// with given name.
func (c *Client) GetServiceByName(service string) (string, error) {
	if !c.opt.Permanent {
		return "", ErrNotSupportRuntime
	}
	var path string
	call, err := c.CallMethod("getServiceByName", service)
	if err := c.handleDBusCall(call, err, &path); err != nil {
		return "", err
	}
	return path, nil
}

// GetServiceNames  Return list of service names (permanent configuration).
func (c *Client) GetServiceNames() ([]string, error) {
	if !c.opt.Permanent {
		return nil, ErrNotSupportRuntime
	}
	var names []string
	call, err := c.CallMethod("getServiceNames")
	if err := c.handleDBusCall(call, err, &names); err != nil {
		return nil, err
	}
	return names, nil
}

// GetZoneByName Return object path (permanent configuration) of zone with given name.
func (c *Client) GetZoneByName(zone string) (string, error) {
	if !c.opt.Permanent {
		return "", ErrNotSupportRuntime
	}
	var path string
	call, err := c.CallMethod("getZoneByName", zone)
	if err := c.handleDBusCall(call, err, &path); err != nil {
		return "", err
	}
	return path, nil
}

// GetZoneNames  Return list of zone names (permanent configuration).
func (c *Client) GetZoneNames() (names []string, err error) {
	if !c.opt.Permanent {
		return names, ErrNotSupportRuntime
	}
	call, err := c.CallMethod("getZoneNames")
	if err != nil {
		return names, err
	}
	err = call.Store(&names)
	return names, err
}

// GetZoneOfSource Return name of zone the source is bound to or empty string.
func (c *Client) GetZoneOfSource(source string) (string, error) {
	if !c.opt.Permanent {
		return "", ErrNotSupportRuntime
	}
	var zoneName string
	call, err := c.CallMethod("getZoneOfSource", source)
	if err := c.handleDBusCall(call, err, &zoneName); err != nil {
		return "", err
	}
	return zoneName, nil
}

// GetServiceSettings Return permanent settings of a service.
func (c *Client) GetServiceSettings(svc string) (ServiceSetting, error) {
	var svcSet ServiceSetting
	var call *dbus.Call
	var err error

	if c.opt.Permanent {
		call, err = c.CallPermanentServiceMethod2(svc, "getSettings")
	} else {
		call, err = c.CallMethod("getServiceSettings", svc)
	}
	if err := c.handleDBusCall(call, err, &svcSet); err != nil {
		return ServiceSetting{}, err
	}
	return svcSet, nil
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
	return c.handleCall(call, err)
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
	return c.handleCall(call, err)
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
	return c.handleCall(call, err)
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
func (c *Client) GetActiveZones() (map[string]ActivateZone, error) {
	var (
		azMap = make(map[string]map[string][]string)
		azs   = make(map[string]ActivateZone)
	)
	call, err := c.CallRuntimeZoneMethod("getActiveZones")
	if err := c.handleDBusCall(call, err, &azMap); err != nil {
		return nil, err
	}
	for zoneName, zone := range azMap {
		azs[zoneName] = ActivateZone{
			Interfaces: zone["interfaces"],
			Sources:    zone["sources"],
		}
	}
	return azs, nil
}

// GetForwardPorts Get list of (port, protocol, toport, toaddr) defined in zone.
func (c *Client) GetForwardPorts(zone string) (ForwardPorts, error) {
	var fps ForwardPorts
	var call *dbus.Call
	var err error

	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getForwardPorts")
	} else {
		call, err = c.CallRuntimeZoneMethod("getForwardPorts", zone)
	}
	if err := c.handleDBusCall(call, err, &fps); err != nil {
		return nil, err
	}
	return fps, nil
}

// GetInterfaces Return array of interfaces (s) previously bound with zone.
func (c *Client) GetInterfaces(zone string) ([]string, error) {
	var interfaces []string
	var call *dbus.Call
	var err error

	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getInterfaces")
	} else {
		call, err = c.CallRuntimeZoneMethod("getInterfaces", zone)
	}
	if err := c.handleDBusCall(call, err, &interfaces); err != nil {
		return nil, err
	}
	return interfaces, nil
}

// GetPorts Return array of ports (2-tuple of port and protocol) previously enabled in zone
func (c *Client) GetPorts(zone string) (Ports, error) {
	var (
		rawPorts [][]string
		ports    Ports
		call     *dbus.Call
		err      error
	)

	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getPorts")
	} else {
		call, err = c.CallRuntimeZoneMethod("getPorts", zone)
	}
	if err := c.handleDBusCall(call, err, &rawPorts); err != nil {
		return nil, err
	}
	ports, err = convertToPorts(rawPorts)
	return ports, err
}

// GetProtocols Return array of protocols (s) previously enabled in zone.
func (c *Client) GetProtocols(zone string) ([]string, error) {
	var protocols []string
	var call *dbus.Call
	var err error

	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getProtocols")
	} else {
		call, err = c.CallRuntimeZoneMethod("getProtocols", zone)
	}
	if err := c.handleDBusCall(call, err, &protocols); err != nil {
		return nil, err
	}
	return protocols, nil
}

// GetRichRules Get list of rich-language rules in zone.
func (c *Client) GetRichRules(zone string) ([]string, error) {
	var rules []string
	var call *dbus.Call
	var err error

	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getRichRules")
	} else {
		call, err = c.CallRuntimeZoneMethod("getRichRules", zone)
	}
	if err := c.handleDBusCall(call, err, &rules); err != nil {
		return nil, err
	}
	return rules, nil
}

// GetServices Get list of service names used in zone.
func (c *Client) GetServices(zone string) ([]string, error) {
	var services []string
	var call *dbus.Call
	var err error

	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getServices")
	} else {
		call, err = c.CallRuntimeZoneMethod("getServices", zone)
	}
	if err := c.handleDBusCall(call, err, &services); err != nil {
		return nil, err
	}
	return services, nil
}

// GetSourcePorts Get list of (port, protocol) defined in zone.
func (c *Client) GetSourcePorts(zone string) (Ports, error) {
	var (
		rawPorts [][]string
		ports    Ports
		call     *dbus.Call
		err      error
	)

	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getSourcePorts")
	} else {
		call, err = c.CallRuntimeZoneMethod("getSourcePorts", zone)
	}
	if err := c.handleDBusCall(call, err, &rawPorts); err != nil {
		return nil, err
	}
	ports, err = convertToPorts(rawPorts)
	return ports, err
}

// GetSources Get list of source addresses bound to zone.
func (c *Client) GetSources(zone string) ([]string, error) {
	var sources []string
	var call *dbus.Call
	var err error

	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getSources")
	} else {
		call, err = c.CallRuntimeZoneMethod("getSources", zone)
	}
	if err := c.handleDBusCall(call, err, &sources); err != nil {
		return nil, err
	}
	return sources, nil
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
	return c.handleCall(call, err)
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
	return c.handleCall(call, err)
}

// RemoveProtocol Remove protocol from zone.
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
	return c.handleCall(call, err)
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
	return c.handleCall(call, err)
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
	return c.handleCall(call, err)
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
	return c.handleCall(call, err)
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
	return c.handleCall(call, err)
}

// GetDefaultZone Return default zone.
func (c *Client) GetDefaultZone() (string, error) {
	if c.opt.Permanent {
		return "", ErrNotSupportPermanent
	}
	var defaultZone string
	call, err := c.CallMethod("getDefaultZone")
	if err := c.handleDBusCall(call, err, &defaultZone); err != nil {
		return "", err
	}
	return defaultZone, nil
}

// GetZones Return array of names (s) of predefined zones known to
// current runtime environment.
func (c *Client) GetZones() ([]string, error) {
	if c.opt.Permanent {
		return nil, ErrNotSupportPermanent
	}
	var zones []string
	call, err := c.CallMethod("getZones")
	if err := c.handleDBusCall(call, err, &zones); err != nil {
		return nil, err
	}
	return zones, nil
}

// ListZones List object paths of zones known to permanent environment.
func (c *Client) ListZones() ([]string, error) {
	if !c.opt.Permanent {
		return nil, ErrNotSupportRuntime
	}
	var zonesPath []string
	call, err := c.CallMethod("listZones")
	if err := c.handleDBusCall(call, err, &zonesPath); err != nil {
		return nil, err
	}
	return zonesPath, nil
}

// GetZoneSettings Return zone settings.
func (c *Client) GetZoneSettings(zone string) (ZoneSetting, error) {
	if c.opt.Permanent {
		return ZoneSetting{}, ErrNotSupportPermanent
	}
	var (
		zs   ZoneSetting
		call *dbus.Call
		err  error
	)

	call, err = c.CallMethod("getZoneSettings", zone)
	if err := c.handleDBusCall(call, err, &zs); err != nil {
		return ZoneSetting{}, err
	}
	return zs, nil
}

func (c *Client) GetZoneSettings2(zone string) (ZoneSetting, error) {
	var (
		zs   ZoneSetting
		call *dbus.Call
		err  error
		sm   map[string]interface{}
	)
	if c.opt.Permanent {
		call, err = c.CallPermanentZoneMethod2(zone, "getSettings2")
	} else {
		call, err = c.CallRuntimeZoneMethod("getZoneSettings2", zone)
	}
	if err := c.handleDBusCall(call, err, &sm); err != nil {
		return ZoneSetting{}, err
	}
	zs, err = convertToZoneSetting(sm)
	if zs.Name == "" {
		zs.Name = zone
	}
	return zs, err
}

func (c *Client) GetSettings2(zone string) (ZoneSetting, error) {
	if !c.opt.Permanent {
		return ZoneSetting{}, ErrNotSupportRuntime
	}
	return c.GetZoneSettings2(zone)
}

// SetDefaultZone Set default zone for connections and interfaces where no zone has been selected to zone.
// Setting the default zone changes the zone for the connections or interfaces,
// that are using the default zone. This is a runtime and permanent change.
func (c *Client) SetDefaultZone(zone string) error {
	call, err := c.CallMethod("setDefaultZone", zone)
	return c.handleCall(call, err)
}

// SetForwardPorts Permanently set forward ports of zone
func (c *Client) SetForwardPorts(zone string, fps ForwardPorts) error {
	if !c.opt.Permanent {
		return ErrNotSupportRuntime
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
	return c.handleCall(call, err)
}

// SetPorts Permanently set ports of zone
func (c *Client) SetPorts(zone string, ports Ports) error {
	if !c.opt.Permanent {
		return ErrNotSupportRuntime
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
	return c.handleCall(call, err)
}

// SetProtocols Permanently set list of protocols used in zone to protocols.
func (c *Client) SetProtocols(zone string, protocols []string) error {
	if !c.opt.Permanent {
		return ErrNotSupportRuntime
	}
	var (
		call *dbus.Call
		err  error
	)
	call, err = c.CallPermanentZoneMethod2(zone, "setProtocols", protocols)
	return c.handleCall(call, err)
}

// SetRichRules Permanently set list of rich-language rules to rules.
func (c *Client) SetRichRules(zone string, rules []string) error {
	if !c.opt.Permanent {
		return ErrNotSupportRuntime
	}
	var (
		call *dbus.Call
		err  error
	)
	call, err = c.CallPermanentZoneMethod2(zone, "setRichRules", rules)
	return c.handleCall(call, err)
}

// SetServices Set services in zone.
func (c *Client) SetServices(zone string, services []string) error {
	if !c.opt.Permanent {
		return ErrNotSupportRuntime
	}
	var (
		call *dbus.Call
		err  error
	)
	call, err = c.CallPermanentZoneMethod2(zone, "setServices", services)
	return c.handleCall(call, err)
}

// SetSourcePorts Set source ports in zone.
func (c *Client) SetSourcePorts(zone string, ports Ports) error {
	if !c.opt.Permanent {
		return ErrNotSupportRuntime
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
	return c.handleCall(call, err)
}

// SetSources Permanently set list of source addresses bound to zone to sources.
func (c *Client) SetSources(zone string, sources []string) error {
	if !c.opt.Permanent {
		return ErrNotSupportRuntime
	}
	var (
		call *dbus.Call
		err  error
	)
	call, err = c.CallPermanentZoneMethod2(zone, "setSources", sources)

	return c.handleCall(call, err)
}

// getZoneID Get zone ID by name.
func (c *Client) getZoneID(zone string) (int, error) {
	path, err := c.GetZoneByName(zone)
	if err != nil {
		return 0, err
	}
	return getIdByPath(path)
}

// getServiceID Get service ID by name.
func (c *Client) getServiceID(svc string) (int, error) {
	path, err := c.GetServiceByName(svc)
	if err != nil {
		return 0, err
	}
	return getIdByPath(path)
}
