package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

func main() {
	err := mountfs()
	if err != nil {
		panic(err)
	}

	fmt.Println("type `poweroff` to quit")
	fmt.Print("$ ")

	scanner := bufio.NewScanner(
		os.Stdin,
	)

	for scanner.Scan() {
		text := scanner.Text()

		switch text {
		case "poweroff":
			err = unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF)
			if err != nil {
				panic(err)
			}

			panic("unreachable")

		case "ls":
			text = "ls ."
		}

		switch {
		case strings.HasPrefix(text, "ls "):
			args := strings.TrimPrefix(text, "ls ")

			info, err := os.Stat(args)
			if err != nil {
				fmt.Println(err)
				break
			}

			if !info.IsDir() {
				fmt.Println(
					info.Mode(),
					info.Name(),
				)
				break
			}

			list, err := os.ReadDir(args)
			if err != nil {
				fmt.Println(err)
				break
			}

			for _, entry := range list {
				fmt.Println(
					entry.Type(),
					entry.Name(),
				)
			}

		case len(text) > 0:
			fmt.Println(text)
		}

		fmt.Print("$ ")
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}
}

func mountfs() error {
	err := unix.Mount("proc", "/proc", "proc", 0, "")
	if err != nil {
		return err
	}

	err = unix.Mount("sysfs", "/sys", "sysfs", 0, "")
	if err != nil {
		return err
	}

	return nil
}
