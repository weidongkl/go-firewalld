/*
 * Copyright © 2024 weidongkl <weidongkx@gmail.com>
 */

package firewalld

import "errors"

var (
	NotSupportPermanentErr = errors.New("this method not supported permanent call")
	UnimplementedErr       = errors.New("this method is not yet implemented")
	NotSupportRuntimeErr   = errors.New("this method not supported Runtime call")
)
