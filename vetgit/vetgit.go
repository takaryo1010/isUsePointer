package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// コマンドラインから取得する
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: govetanddelete <package>")
		os.Exit(1)
	}
	pkg := os.Args[1]

	// パッケージをダウンロードする
	cmd := exec.Command("go", "get", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// go vet を実行する
	cmd = exec.Command("go", "vet", "-vettool="+pwd()+"/../isUsePointer.exe", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// パッケージを削除する
	err = os.RemoveAll(filepath.Join(os.Getenv("GOPATH"), "pkg", "mod", pkg))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Package %s successfully downloaded, vetted, and deleted.\n", pkg)
}
func pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
