package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func runTool(cmd string, args []string) {
	c := exec.Command(cmd, args...)

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error running %s: %s", cmd, err)
		os.Exit(1)
	}
}

func ensureInstalled(bin string, pkg string) {
	_, err := exec.LookPath(bin)
	if err != nil {
		fmt.Printf("%s not installed, installing from %s\n", bin, pkg)
		err = exec.Command("go", "get", pkg).Run()
		if err != nil {
			fmt.Printf("Failed to install %s via go get %s\n", bin, pkg)
			panic(err)
		}
	}
}

func addUniqueString(strs []string, str string) []string {
	for _, s := range strs {
		if s == str {
			return strs
		}
	}

	return append(strs, str)
}

func dirNames(filepaths []string) []string {
	var dirs []string

	for _, fpath := range filepaths {
		dirs = addUniqueString(dirs, filepath.Dir(fpath))
	}

	return dirs
}

func pkgNames(dirnames []string) []string {
	var pkgs []string

	for _, dir := range dirnames {
		pkgs = append(pkgs, "./"+dir)
	}

	return pkgs
}
