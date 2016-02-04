package util

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func GetUserHome(userName string) string {
	usr, err := user.Lookup(userName)
	if err == nil {
		return usr.HomeDir
	}

	return fmt.Sprintf("/home/%s", userName)
}

func GetMyHome() string {
	home := os.Getenv("HOME")
	if home != "" {
		return home
	}

	usr, err := user.Current()
	if err == nil {
		return usr.HomeDir
	}

	userName := os.Getenv("USER")
	if userName == "" {
		log.Println(err)
		return ""
	}

	return GetUserHome(userName)
}

func PathResolv(base string, s string) string {
	if len(s) > 0 && s[0] == '~' {
		a := strings.Split(s, "/")
		if a[0] == "~" {
			a[0] = GetMyHome()
		} else {
			a[0] = GetUserHome(a[0][1:])
		}
		return filepath.Join(a...)
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
