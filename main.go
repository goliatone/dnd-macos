package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const (
	enableFlag  = "enable"
	disableFlag = "disable"

	killall             = "killall NotificationCenter"
	setDoNotDisturb     = "defaults -currentHost write ~/Library/Preferences/ByHost/com.apple.notificationcenterui doNotDisturb -boolean true"
	setDoNotDisturbDate = "defaults -currentHost write ~/Library/Preferences/ByHost/com.apple.notificationcenterui doNotDisturbDate -date \"`date -u +\"%Y-%m-%d %H:%M:%S +000\"`\""
	unsetDoNotDisturb   = "defaults -currentHost write ~/Library/Preferences/ByHost/com.apple.notificationcenterui doNotDisturb -boolean false"
)

var usage func()

func init() {
	usage = func() {
		fmt.Println("Utility to toggle macOS do not disturb using notification center")
		fmt.Fprintf(os.Stderr, "Usage: %s [subcommand] {enable|disable}\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
}
func main() {

	flag.Parse()

	if len(flag.Args()) == 0 {
		usage()
	}

	switch os.Args[1] {
	case enableFlag:
		enable()
		break
	case disableFlag:
		disable()
		break
	default:
		usage()
	}
}

func enable() {
	cmd := exec.Command("sh", "-c", setDoNotDisturb)
	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not set doNotDisturb to true: %v", err)
		os.Exit(1)
	}

	cmd = exec.Command("sh", "-c", setDoNotDisturbDate)
	_, err = cmd.CombinedOutput()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not write time to ~/Library/Preferences/ByHost/com.apple.notificationcenterui: %v", err)
		os.Exit(1)
	}

	cmd = exec.Command("sh", "-c", killall)
	_, err = cmd.CombinedOutput()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not reset NotificationCenter: %v", err)
		os.Exit(1)
	}
}

func disable() {
	cmd := exec.Command("sh", "-c", unsetDoNotDisturb)
	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not set doNotDisturb to false: %v", err)
		os.Exit(1)
	}

	cmd = exec.Command("sh", "-c", killall)
	_, err = cmd.CombinedOutput()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not reset NotificationCenter: %v", err)
		os.Exit(1)
	}
}
