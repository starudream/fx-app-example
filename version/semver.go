package version

import (
	"github.com/Masterminds/semver/v3"
)

type Version = semver.Version

var (
	New = semver.New

	NewVersion       = semver.NewVersion
	StrictNewVersion = semver.StrictNewVersion

	Parse     = semver.NewVersion
	MustParse = semver.MustParse
)
