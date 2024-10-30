/*
 * Copyright © 2024 weidongkl <weidongkx@gmail.com>
 */

package firewalld

import (
	"fmt"
	"github.com/godbus/dbus"
	"strconv"
	"strings"
)

const dbusName = "org.fedoraproject.FirewallD1"

type Client struct {
	opt           *Options
	conn          *dbus.Conn
	dbusInterface string
	dbusName      string
	dbusPath      dbus.ObjectPath
	obj           dbus.BusObject
}

func NewClient(opt *Options) (*Client, error) {
	c := &Client{
		opt: opt,
	}
	c.opt.init()
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	c.conn = conn
	if c.opt.Permanent {
		c.dbusPath = dbusPermanentPath
		c.dbusInterface = dbusPermanentInterface
	} else {
		c.dbusPath = dbusRuntimePath
		c.dbusInterface = dbusRuntimeInterface
	}
	c.dbusName = dbusName
	c.obj = conn.Object(c.dbusName, c.dbusPath)
	return c, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) CallMethod(method string, args ...interface{}) (*dbus.Call, error) {
	return c.obj.Call(c.dbusInterface+"."+method, 0, args...), nil
}

func (c *Client) CallRuntimeZoneMethod(method string, args ...interface{}) (*dbus.Call, error) {
	return c.obj.Call(dbusRuntimeZoneInterface+"."+method, 0, args...), nil
}

func (c *Client) CallPermanentZoneMethod(zoneId int, method string, args ...interface{}) (*dbus.Call, error) {
	objPath := dbus.ObjectPath(fmt.Sprintf("%s/%d", dbusPermanentZoneBasePath, zoneId))
	obj := c.conn.Object(c.dbusName, objPath)
	return obj.Call(dbusPermanentZoneInterface+"."+method, 0, args...), nil
}

func (c *Client) CallPermanentZoneMethod2(zone string, method string, args ...interface{}) (*dbus.Call, error) {
	zoneId, err := c.getZoneID(zone)
	if err != nil {
		return nil, err
	}
	objPath := dbus.ObjectPath(fmt.Sprintf("%s/%d", dbusPermanentZoneBasePath, zoneId))
	obj := c.conn.Object(c.dbusName, objPath)
	return obj.Call(dbusPermanentZoneInterface+"."+method, 0, args...), nil
}
func (c *Client) CallPermanentServiceMethod(svcId int, method string, args ...interface{}) (*dbus.Call, error) {
	objPath := dbus.ObjectPath(fmt.Sprintf("%s/%d", dbusPermanentServiceBasePath, svcId))
	obj := c.conn.Object(c.dbusName, objPath)
	return obj.Call(dbusPermanentServiceInterface+"."+method, 0, args...), nil
}

func (c *Client) CallPermanentServiceMethod2(svc string, method string, args ...interface{}) (*dbus.Call, error) {
	svcId, err := c.getServiceID(svc)
	if err != nil {
		return nil, err
	}
	objPath := dbus.ObjectPath(fmt.Sprintf("%s/%d", dbusPermanentServiceBasePath, svcId))
	obj := c.conn.Object(c.dbusName, objPath)
	return obj.Call(dbusPermanentServiceInterface+"."+method, 0, args...), nil
}

// get id from object path
func getIdByPath(path string) (id int, err error) {
	pathSlice := strings.Split(path, "/")
	id, err = strconv.Atoi(pathSlice[len(pathSlice)-1])
	return
}
