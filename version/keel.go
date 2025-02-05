package version

import (
	"runtime"

	"github.com/alwinius/bow/types"
)

// Generic tool info
const (
	ProductName string = "bow"
	APIVersion         = "1"
)

// Revision that was compiled. This will be filled in by the compiler.
var Revision string

// BuildDate is when the binary was compiled.  This will be filled in by the
// compiler.
var BuildDate string

// Version number that is being run at the moment.  Version should use semver.
var Version string

// Experimental is intended to be used to enable alpha features.
var Experimental string

// GetbowVersion returns version info.
func GetbowVersion() types.VersionInfo {
	v := types.VersionInfo{
		Name:       ProductName,
		Revision:   Revision,
		BuildDate:  BuildDate,
		Version:    Version,
		APIVersion: APIVersion,
		GoVersion:  runtime.Version(),
		OS:         runtime.GOOS,
		Arch:       runtime.GOARCH,
	}

	return v
}
