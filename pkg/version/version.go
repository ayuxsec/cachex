package version

import "runtime/debug"

var Version = "dev"

func String() string {
	// GoReleaser build path
	if Version != "dev" {
		return Version
	}

	// go install @x.x.x path
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			return info.Main.Version
		}
	}

	// local build
	return "dev"
}
