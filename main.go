package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

var Version string

func GetLogin() (string, error) {
	const command = "who | grep console | awk '{print $1}'"
	login, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(login), "\n"), nil
}

func DetectMagicKeyboard() (bool, error) {
	out, err := exec.Command("bash", "-c", "\"system_profiler SPUSBDataType 2> /dev/null | grep 'Magic Keyboard'\"").Output()
	if err != nil {
		return false, err
	}
	if len(out) == 0 {
		return false, err
	}
	return true, nil
}

func GetHostname() string {
	hostname, _ := os.Hostname()
	strs := strings.Split(hostname, ".")
	return strs[0]
}

func main() {
	log.Print("Version: ", Version)

	log.Println("KeyboardChecker init. delayed start 1 minute")
	time.Sleep(time.Minute)

	conn, err := net.Dial("udp", os.Args[1])
	if err != nil {
		log.Println(err)
	}
	usb := true
	for true {
		flag, err := DetectMagicKeyboard()
		if err != nil {
			log.Println(err)
			continue
		}

		login, err := GetLogin()
		if err != nil {
			log.Println(err)
			continue
		}

		if flag {
			if usb {
				continue
			} else {
				usb = true
			}
		} else {
			if usb {
				conn.Write([]byte(fmt.Sprintf("<30> %s [%s] Magic Keyboard: disabled", GetHostname(), login)))
				usb = false
			}
		}

		time.Sleep(5 * time.Second)
	}
}
