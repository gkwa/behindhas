package main

import (
   "testing"
)

func TestRunProblem(t *testing.T) {
   err := run("echo", "Running command with:", "-flag1", "-flag2", SRC...)
   if err == nil {
   	t.Error("Expected an error, but got none")
   }
}

func TestRunSolution(t *testing.T) {
   args := append([]string{"Running command with:", "-flag1", "-flag2"}, SRC...)
   err := run("echo", args...)
   if err != nil {
   	t.Errorf("Unexpected error: %v", err)
   }
}

