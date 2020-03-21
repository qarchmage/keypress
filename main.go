package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-vgo/robotgo"
)

const usage = "[key] [ms]"

func main() {
	var (
		hotkey string
		ms     int
		err    error
	)
	arg := os.Args
	switch len(arg) {
	case 1:
		fmt.Println(usage)
		return
	case 2:
		hotkey = arg[1]
		ms = 100
	case 3:
		hotkey = arg[1]
		ms, err = strconv.Atoi(arg[2])
		if err != nil {
			fmt.Println("second arg must be a number")
		}
	}

	for {
		robotgo.KeyTap(hotkey)
		robotgo.MilliSleep(ms)
	}
}
