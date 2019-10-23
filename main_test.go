package main

import (
  "testing"
  "os"
  "bufio"
  "strings"
)

var path string = "list.txt"

func TestOpenFile(t *testing.T) {
  if _, err := os.Open(path); err != nil {
      t.Error("Canno't open ", path)
  }
}


func TestReadFile(t *testing.T) {
  file := OpenFile(path)
  defer file.Close()

  scanner := bufio.NewScanner(file)
  AnalyzeFile(scanner)
}

func TestAnalyzeFile(t *testing.T) {

  lines := []string{}
  duplicateLines := []string{}

  file := OpenFile(path)
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      line := scanner.Text()
      findDuplicateLine := false

      if len(lines) != 0 {
          for i := range lines {
              if strings.Compare(lines[i], line) == 0 {
                  duplicateLines = append(duplicateLines, line)
                  findDuplicateLine = true
                  break
              }
          }
      }

      if !findDuplicateLine {
          lines = append(lines, line)
      }
  }
}

func TestWriteOutput(t *testing.T) {
  file, err := os.Create("TestWriteOutput.txt")
  if err != nil {
    t.Error("Something wrong with file creation")
  }
  dataWriter := bufio.NewWriter(file)
  _, _ = dataWriter.WriteString("test" + "\r\n")

}
