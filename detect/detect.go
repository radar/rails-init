package detect

import (
	"io/ioutil"
	"runtime"
	"strings"
)

func IsMac() bool {
	return runtime.GOOS == "darwin"
}

func IsUbuntu() bool {
	osRelease, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		panic(err)
	}

	if strings.Contains(string(osRelease), "NAME=\"Ubuntu\"") {
		return true
	}

	return false
}
