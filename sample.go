package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	pkgnames := []string{"fmt", "time", "strings", "os", "strconv"}
	for _, pkgname := range pkgnames {
		cmd := exec.Command("go", "vet", "-vettool="+pwd()+"/isUsePointer.exe", pkgname)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		
	}

}

func pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
