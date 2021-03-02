//go:build !go1.16
// +build !go1.16

//go:generate statik -src=. -include=*.

package witticism

import (
	"github.com/rakyll/statik/fs"

	_ "github.com/imishinist/ayb/pkg/witticism/statik"
)

var (
	witticismJson []byte
)

func init() {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	witticismJson, err = fs.ReadFile(statikFS, "/witticism.json")
	if err != nil {
		panic(err)
	}
}
