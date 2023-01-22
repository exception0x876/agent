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
	path, err := exec.LookPath(config.smartctl)
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd, err := exec.Command(path, "--json", "-a", "/dev/disk0").Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode := exitError.ExitCode()
			//just print the exit code for now
			fmt.Println(exitCode)
		}
	}
	var parsed SmartInfo
	err = json.Unmarshal(cmd, &parsed)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d here\n", parsed.temperature.current)
}
