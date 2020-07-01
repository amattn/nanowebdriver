package util

import (
	"fmt"
	"time"
)

const (
	internalIdentifier           = "github.com/amattn/wdc"
	internalBuildTimestamp int64 = 1593638058
	internalBuildNumber    int64 = 3
	internalVersionString        = "0.0.1"
)

func BuildDate() time.Time {
	return time.Unix(internalBuildTimestamp, 0)
}

func BuildNumber() int64 {
	return internalBuildNumber
}

func Version() string {
	return internalVersionString
}

func VersionInfo() string {
	return fmt.Sprintf("%s (%v, build %v, build date:%v)", internalIdentifier, Version(), BuildNumber(), BuildDate())
}
