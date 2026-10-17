package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	build := exec.Command("go", "build", "-o", "web/game.wasm", ".")
	build.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	build.Stdout, build.Stderr = os.Stdout, os.Stderr
	if err := build.Run(); err != nil {
		log.Fatal(err)
	}

	goroot, err := exec.Command("go", "env", "GOROOT").Output()
	if err != nil {
		log.Fatal(err)
	}
	js, err := os.ReadFile(filepath.Join(strings.TrimSpace(string(goroot)), "lib", "wasm", "wasm_exec.js"))
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("web/wasm_exec.js", js, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Println("built web/game.wasm")
}
