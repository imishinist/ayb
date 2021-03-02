//go:build go1.16
// +build go1.16

package witticism

import (
	_ "embed"
)

var (
	//go:embed witticism.json
	witticismJson []byte
)
