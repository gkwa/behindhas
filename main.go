package main

import (
   "fmt"
   "os"
   "os/exec"
)

var SRC = []string{"file1.go", "file2.go", "file3.go"}

func main() {
   fmt.Println("Demonstrating the solution:")
   args := append([]string{"Running command with:", "-flag1", "-flag2"}, SRC...)
   err := run("echo", args...)
   if err != nil {
   	fmt.Println("Error:", err)
   }
}

func run(name string, args ...string) error {
   cmd := exec.Command(name, args...)
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   return cmd.Run()
}
