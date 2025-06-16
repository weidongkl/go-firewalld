/*
Copyright © 2024 weidongkl <weidongkx@gmail.com>
*/

package firewalld

import "fmt"

const (
	// path
	dbusRuntimePath              = "/org/fedoraproject/FirewallD1"
	dbusPermanentPath            = "/org/fedoraproject/FirewallD1/config"
	dbusPermanentZoneBasePath    = "/org/fedoraproject/FirewallD1/config/zone"
	dbusPermanentServiceBasePath = "/org/fedoraproject/FirewallD1/config/service"
	// interface
	dbusRuntimeInterface          = "org.fedoraproject.FirewallD1"
	dbusRuntimeZoneInterface      = "org.fedoraproject.FirewallD1.zone"
	dbusPermanentInterface        = "org.fedoraproject.FirewallD1.config"
	dbusPermanentZoneInterface    = "org.fedoraproject.FirewallD1.config.zone"
	dbusPermanentServiceInterface = "org.fedoraproject.FirewallD1.config.service"
)

type Port struct {
	Port     string
	Protocol string
}

type ForwardPort struct {
	Port      string
	Protocol  string
	ToPort    string
	ToAddress string
}
type Ports []Port
type ForwardPorts []ForwardPort

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

type ActivateZone struct {
	Interfaces []string
	Sources    []string
}

type ServiceSettingMap map[string]interface{}
type ServiceSetting struct {
	Version      string
	Name         string
	Description  string
	Ports        Ports
	ModuleNames  []string
	Destinations map[string]string
	Protocols    []string
	SourcePorts  Ports
	//Includes     []string
}

func (ssm ServiceSettingMap) ToStruct() (ss ServiceSetting, err error) {
	if ssm["version"] != nil {
		ss.Version = ssm["version"].(string)
	}
	if ssm["name"] != nil {
		ss.Name = ssm["name"].(string)
	}
	if ssm["description"] != nil {
		ss.Description = ssm["description"].(string)
	}
	if ssm["ports"] != nil {
		ss.Ports, err = convertToPorts(ssm["ports"].([][]string))
		if err != nil {
			return ss, err
		}
	}
	if ssm["module names"] != nil {
		ss.ModuleNames = ssm["module names"].([]string)
	}
	if ssm["destinations"] != nil {
		ss.Destinations = ssm["destinations"].(map[string]string)

	}
	if ssm["protocols"] != nil {
		ss.Protocols = ssm["protocols"].([]string)
	}
	if ssm["sourceports"] != nil {
		ss.SourcePorts, err = convertToPorts(ssm["sourceports"].([][]string))
		if err != nil {
			return ss, err
		}

	}
	if ssm["includes"] != nil {
		//ss.Includes = ssm["includes"].([]string)
	}
	return
}

func convertToPorts(strSlice [][]string) (Ports, error) {
	ports := make(Ports, 0, len(strSlice))
	for _, item := range strSlice {
		if len(item) != 2 {
			return nil, fmt.Errorf("invalid port format, expected [port, protocol], got %v", item)
		}
		ports = append(ports, Port{
			Port:     item[0],
			Protocol: item[1],
		})
	}
	return ports, nil
}

func convertToForwardPorts(strSlice [][]interface{}) (ForwardPorts, error) {
	if len(strSlice) == 0 {
		return nil, nil
	}

	forwardPorts := make(ForwardPorts, 0, len(strSlice))

	for i, item := range strSlice {
		if len(item) != 4 {
			return nil, fmt.Errorf("invalid forward port format at index %d: expected 4 elements (port, protocol, to-address, to-port), got %d", i, len(item))
		}

		port, ok1 := item[0].(string)
		protocol, ok2 := item[1].(string)
		toAddr, ok3 := item[2].(string)
		toPort, ok4 := item[3].(string)

		if !ok1 || !ok2 || !ok3 || !ok4 {
			return nil, fmt.Errorf("type assertion failed at index %d: all elements must be strings, got %T, %T, %T, %T",
				i, item[0], item[1], item[2], item[3])
		}

		forwardPorts = append(forwardPorts, ForwardPort{
			Port:      port,
			Protocol:  protocol,
			ToAddress: toAddr,
			ToPort:    toPort,
		})
	}

	return forwardPorts, nil
}

func convertToZoneSetting(kv map[string]interface{}) (ZoneSetting, error) {
	var (
		zoneSetting ZoneSetting
		err         error
	)
	if kv["version"] != nil {
		zoneSetting.Version = kv["version"].(string)
	}
	if kv["name"] != nil {
		zoneSetting.Name = kv["name"].(string)
	}
	if kv["description"] != nil {
		zoneSetting.Description = kv["description"].(string)
	}
	if kv["unused"] != nil {
		zoneSetting.Unused = kv["unused"].(bool)
	}
	if kv["target"] != nil {
		zoneSetting.Target = kv["target"].(string)
	}
	if kv["services"] != nil {
		zoneSetting.Services = kv["services"].([]string)
	}
	if kv["ports"] != nil {
		zoneSetting.Ports, err = convertToPorts(kv["ports"].([][]string))
		if err != nil {
			return zoneSetting, err
		}
	}
	if kv["icmp_blocks"] != nil {
		zoneSetting.IcmpBlocks = kv["icmp_blocks"].([]string)
	}
	if kv["masquerade"] != nil {
		zoneSetting.Masquerade = kv["masquerade"].(bool)
	}
	if kv["forward_ports"] != nil {
		zoneSetting.ForwardPorts, err = convertToForwardPorts(kv["forward_ports"].([][]interface{}))
		if err != nil {
			return zoneSetting, err
		}
	}
	if kv["interfaces"] != nil {
		zoneSetting.Interfaces = kv["interfaces"].([]string)
	}
	if kv["sources"] != nil {
		zoneSetting.SourceAddresses = kv["sources"].([]string)
	}
	if kv["rules_str"] != nil {
		zoneSetting.RichRules = kv["rules_str"].([]string)
	}
	if kv["protocols"] != nil {
		zoneSetting.Protocols = kv["protocols"].([]string)
	}
	if kv["source_ports"] != nil {
		zoneSetting.SourcePorts, err = convertToPorts(kv["source_ports"].([][]string))
	}
	if kv["icmp_block_inversion"] != nil {
		zoneSetting.IcmpBlockInversion = kv["icmp_block_inversion"].(bool)
	}
	return zoneSetting, nil
}
