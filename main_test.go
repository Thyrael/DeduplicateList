package main

import (
  "testing"
  "os"
  "bufio"
)

var path string = "list.txt"

func TestOpenFile(t *testing.T) {
  if _, err := os.Open(path); err != nil {
      t.Error("Canno't open ", path)
  }
}


func TestReadFile(t *testing.T) {
  if file, err := OpenFile(path)
  defer file.Close()

  scanner := bufio.NewScanner(file)
  AnalyzeFile(scanner)
}
