// +build darwin

package fish

import (
	"os"
)

// HomePrefix is the base path to Fish's configuration directory.
const HomePrefix = "/usr/local"

// BinPath is the path to where executables should be installed by fish.
const BinPath = HomePrefix + "/bin"

// UserHomePath is the path to $HOME
var UserHomePath = os.Getenv("HOME")

// Cache returns the path to the cache.
func (h UserHome) Cache() string {
	return h.Path("Library", "Caches", "Fish")
}
