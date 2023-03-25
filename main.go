package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
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

	os.Chdir(path)
	execCmd("go", "mod", "init", "sandbox")

	if editor := os.Getenv("EDITOR"); editor != "" {
		execCmd(editor, f1.Name())
		return
	}

	execCmd("code", path, "--goto", f1.Name()+":4:2")
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
