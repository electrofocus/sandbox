package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	dir, err := os.MkdirTemp(dir, "sandbox_*")
	if err != nil {
		fmt.Printf("can't create temporary directory (%s)", err)
		return
	}

	f, err := os.Create(filepath.Join(dir, "main.go"))
	if err != nil {
		fmt.Printf("can't create main.go file (%s)", err)
		return
	}

	defer f.Close()
	f.WriteString(`package main

func main() {
	
}
`)

	os.Chdir(dir)
	execCmd("go", "mod", "init", "sandbox")

	if editor := os.Getenv("EDITOR"); editor != "" {
		execCmd(editor, f.Name())
		return
	}

	execCmd("code", dir, "--goto", f.Name()+":4:2")
}

func execCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("can't exec command %q (%s)\n", name, err)
		return
	}
}
