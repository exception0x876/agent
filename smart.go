package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type SmartInfoTemperature struct {
	current int
}

type SmartInfo struct {
	temperature SmartInfoTemperature
}

func monitorSmart(config *Config) {
	defer wg.Done()
	cmd, err := exec.Command("smartctl", "--json", "-a", "/dev/disk0").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", cmd)
	var parsed SmartInfo
	err = json.Unmarshal(cmd, &parsed)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d here\n", parsed.temperature.current)
}
