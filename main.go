package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var version string
	flag.StringVar(&version, "go", "1.20", "Go module version")
	flag.Parse()

	path, err := os.MkdirTemp("", "sandbox_*")
	if err != nil {
		fmt.Printf("can't create temporary directory (%s)", err)
		return
	}

	f1, err := os.Create(filepath.Join(path, "main.go"))
	if err != nil {
		fmt.Printf("can't create main.go file (%s)", err)
		return
	}

	defer f1.Close()
	f1.WriteString(`package main

func main() {
	
}
`)

	f2, err := os.Create(filepath.Join(path, "go.mod"))
	if err != nil {
		fmt.Printf("can't create go.mod file (%s)", err)
		return
	}

	defer f2.Close()
	f2.WriteString(`module sandbox

go `)
	f2.WriteString(version)
	f2.WriteString("\n")

	if _, err := exec.Command("code", path, "--goto", f1.Name()+":4:2").Output(); err != nil {
		fmt.Printf("can't get command output (%s)", err)
		return
	}
}
