package main

import (
	"os"
	"strings"
)

func ReadInput() ([]string, error) {
	openedFile, err := os.ReadFile("/home/jwilliams/Sync/homelab/golang/2022-advent-of-code/Day 2/inputs.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(openedFile), "\n")
	return lines, err
}
