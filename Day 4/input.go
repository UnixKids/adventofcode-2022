package main

import (
	"os"
	"strings"
)

func ReadInput() ([]string, error) {
	openedFile, err := os.ReadFile("/home/jwilliams/Sync/homelab/golang/adventofcode-2022/Day 4/inputs.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(openedFile), "\n")
	return lines, err
}
