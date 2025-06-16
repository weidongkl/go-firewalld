/*
 * Copyright © 2024 weidongkl <weidongkx@gmail.com>
 */

package firewalld

import (
	"errors"
)

var (
	// Base errors
	ErrNotSupportPermanent = errors.New("this method not supported permanent call")
	ErrUnimplemented       = errors.New("this method is not yet implemented")
	ErrNotSupportRuntime   = errors.New("this method not supported Runtime call")

	// D-Bus related errors
	ErrDBusConnection = errors.New("failed to establish D-Bus connection")
	ErrDBusCall       = errors.New("D-Bus call failed")

	// Configuration errors
	ErrInvalidZone       = errors.New("invalid zone configuration")
	ErrInvalidService    = errors.New("invalid service configuration")
	ErrInvalidPort       = errors.New("invalid port configuration")
	ErrInvalidProtocol   = errors.New("invalid protocol configuration")
	ErrInvalidTimeout    = errors.New("invalid timeout value")
	ErrInvalidRetryCount = errors.New("invalid retry count value")
)
