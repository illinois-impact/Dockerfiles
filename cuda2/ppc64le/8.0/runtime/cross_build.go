package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func crossBuildStart() {
	err := os.Remove("/bin/sh")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Link("/usr/bin/cross_build", "/bin/sh")
	if err != nil {
		log.Fatal(err)
	}
}

func crossBuildEnd() {
	err := os.Remove("/bin/sh")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Link("/bin/sh.real", "/bin/sh")
	if err != nil {
		log.Fatal(err)
	}
}

func runShell() error {
	// qemu := fmt.Sprintf("/usr/bin/qemu-%s-static", os.Getenv("ARCH"))
	qemu := "/usr/bin/qemu-ppc64le-static"
	// println("running = ", append([]string{"-0", "/bin/sh", "/bin/sh"}, os.Args[1:]...))
	cmd := exec.Command(qemu, append([]string{"-0", "/bin/sh", "/bin/sh"}, os.Args[1:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {

	log.SetOutput(os.Stderr)

	switch os.Args[0] {
	case "cross-build-start":
		crossBuildStart()
	case "cross-build-end":
		crossBuildEnd()
	case "/bin/sh":
		code := 0
		crossBuildEnd()

		if err := runShell(); err != nil {
			code = 1
			if exiterr, ok := err.(*exec.ExitError); ok {
				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					code = status.ExitStatus()
				}
			}
		}

		crossBuildStart()

		os.Exit(code)
	}
}
