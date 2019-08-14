package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main () {
  path := os.Args[1] // Get file path
  fmt.Printf("File to De-Duplicate : %v\n", path)
  lines, duplicateLines := ReadFile(path)
  WriteOutputFileWithoutDuplicatesLines(lines)
  WriteOutputFileWithDuplicateLines(duplicateLines)
}

func ReadFile(path string) (lines []string, duplicateLines[]string) {

    file := OpenFile(path)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    return AnalyzeFile(scanner)
}

func OpenFile(path string) (file *os.File) {
  file, err := os.Open(path)
  if err != nil {
      return
  }
  return file
}

func AnalyzeFile(scanner *bufio.Scanner) (lines []string, duplicateLines[]string) {
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
  return lines, duplicateLines
}


func WriteOutputFileWithoutDuplicatesLines(lines []string ){
    f, err := os.Create("NoDuplicate.txt")
    if err != nil {
        return
    }
    defer f.Close()
    for i := range lines {
        WriteOutput(f, lines[i])
    }
}

func WriteOutputFileWithDuplicateLines(duplicateLines []string) {
    f, err := os.Create("Duplicate.txt")
    if err != nil {
        return
    }
    defer f.Close()
    WriteOutput(f, "********************************")
    WriteOutput(f, "**  List of duplicate lines   **")
    WriteOutput(f, "********************************")
    for i := range duplicateLines {
        WriteOutput(f, duplicateLines[i])
    }
}

func WriteOutput(file *os.File, s string) {
    dataWriter := bufio.NewWriter(file)
    _, _ = dataWriter.WriteString(s + "\r\n")
    defer dataWriter.Flush()
}
