package util

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func PathResolv(base string, s string) string {
	var err error
	if s[0] == '~' {
		a := strings.Split(s, "/")
		var usr *user.User
		if a[0] == "~" {
			usr, err = user.Current()
		} else {
			usr, err = user.Lookup(a[0][1:])
		}
		if err != nil {
			log.Printf("resolv error: %v\n", err)
		} else {
			a[0] = usr.HomeDir
			return filepath.Join(a...)
		}
	}

	if filepath.IsAbs(s) {
		return s
	}

	return filepath.Join(base, s)
}

func PathResolvWithMkdirAll(base string, s string) string {
	resolved := PathResolv(base, s)
	pathname, _ := filepath.Split(resolved)
	os.MkdirAll(pathname, 0777)
	return resolved
}
