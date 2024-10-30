/*
Copyright © 2024 weidongkl <weidongkx@gmail.com>
*/

package firewalld

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
		ss.Ports, err = convertToPorts(ssm["ports"].([][]interface{}))
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
		ss.SourcePorts, err = convertToPorts(ssm["sourceports"].([][]interface{}))
		if err != nil {
			return ss, err
		}

	}
	if ssm["includes"] != nil {
		//ss.Includes = ssm["includes"].([]string)
	}
	return
}

func convertToPorts(ports [][]interface{}) (structPorts Ports, err error) {
	for _, pp := range ports {
		_port := Port{
			Port:     pp[0].(string),
			Protocol: pp[1].(string),
		}
		structPorts = append(structPorts, _port)
	}
	return structPorts, nil
}
