package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main () {
  fmt.Println("Hello, world.")

  path := os.Args[1] // Get file path
  fmt.Println("File to De-Duplicate : " +path)

  lines, duplicateLines := ReadFile(path)

  WriteOutputFileWithoutDuplicatesLines(lines)
  WriteOutputFileWithDuplicateLines(duplicateLines)

  fmt.Println("End")
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

func WriteOutput(file *os.File, s string) {
    dataWriter := bufio.NewWriter(file)
    _, _ = dataWriter.WriteString(s + "\r\n")
    defer dataWriter.Flush()
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


func ReadFile(path string) (lines []string, duplicateLines[]string) {
    file, err := os.Open(path)
    if err != nil {
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if len(lines) == 0 {
            lines = append(lines, line)
        } else {
            findDuplicateLine := false
            for i := range lines {
                if strings.Compare(lines[i], line) == 0 {
                    duplicateLines = append(duplicateLines, line)
                    findDuplicateLine = true
                    break
                }
            }
            if !findDuplicateLine {
                lines = append(lines, line)
            }
        }
    }
    return lines, duplicateLines
}
