/*
 * Copyright © 2024 weidongkl <weidongkx@gmail.com>
 */

package firewalld

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/godbus/dbus"
)

const dbusName = "org.fedoraproject.FirewallD1"

type Client struct {
	opt           *Options
	conn          *dbus.Conn
	dbusInterface string
	dbusName      string
	dbusPath      dbus.ObjectPath
	obj           dbus.BusObject
	connMutex     sync.RWMutex // Add mutex for thread safety
}

// NewClient creates a new firewalld client with optimized connection handling
func NewClient(opt *Options) (*Client, error) {
	if opt == nil {
		return nil, fmt.Errorf("options cannot be nil")
	}

	c := &Client{
		opt: opt,
	}
	c.opt.init()

	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to system bus: %w", err)
	}

	c.conn = conn
	c.dbusName = dbusName

	if c.opt.Permanent {
		c.dbusPath = dbusPermanentPath
		c.dbusInterface = dbusPermanentInterface
	} else {
		c.dbusPath = dbusRuntimePath
		c.dbusInterface = dbusRuntimeInterface
	}

	c.obj = conn.Object(c.dbusName, c.dbusPath)
	return c, nil
}

// Close safely closes the DBus connection
func (c *Client) Close() error {
	c.connMutex.Lock()
	defer c.connMutex.Unlock()

	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}
	return nil
}

// ensureConnection checks if the connection is valid and reconnects if necessary
func (c *Client) ensureConnection() error {
	c.connMutex.RLock()
	if c.conn != nil {
		c.connMutex.RUnlock()
		return nil
	}
	c.connMutex.RUnlock()

	c.connMutex.Lock()
	defer c.connMutex.Unlock()

	// Double check after acquiring write lock
	if c.conn != nil {
		return nil
	}

	conn, err := dbus.SystemBus()
	if err != nil {
		return fmt.Errorf("failed to reconnect to system bus: %w", err)
	}

	c.conn = conn
	c.obj = conn.Object(c.dbusName, c.dbusPath)
	return nil
}

// CallMethod is an optimized version that handles connection state
func (c *Client) CallMethod(method string, args ...interface{}) (*dbus.Call, error) {
	if err := c.ensureConnection(); err != nil {
		return nil, err
	}
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
