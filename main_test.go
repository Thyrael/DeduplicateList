package main

import (
  "testing"
  "os"
)

func TestOpenFile(t *testing.T) {
  path := "list.txt"
  _, err := os.Open(path)
  if err != nil {
      t.Error("Canno't open ", path)
  }
}
